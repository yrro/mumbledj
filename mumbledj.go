/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleutil"
	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/matthieugrieger/mumbledj/commands"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// MumbleDJ is a struct that keeps track of all aspects of the bot's state.
type MumbleDJ struct {
	KeepAlive chan bool
	Commander *commands.Commander
	State     *state.BotState
}

// OnConnect event. First moves MumbleDJ into default channel if one exists. The
// configuration is loaded and the audio stream is set up.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {
	dj.State.Client.Self.Move(dj.State.Client.Channels.Find(viper.GetStringSlice("general.defaultchannel")...))

	dj.State.AudioStream = nil
	dj.State.AudioStream.Volume = float32(viper.GetFloat64("volume.default"))

	dj.State.Client.Self.SetComment(viper.GetString("general.defaultcomment"))

	if viper.GetBool("cache.enabled") {
		dj.State.Cache.UpdateStats()
		go dj.State.Cache.CleanPeriodically()
	}
}

// OnDisconnect event. Terminates MumbleDJ thread, or retries connection if
// automatic connection retries are enabled.
func (dj *MumbleDJ) OnDisconnect(e *gumble.DisconnectEvent) {
	if viper.GetBool("connection.retryenabled") && (e.Type == gumble.DisconnectError || e.Type == gumble.DisconnectKicked) {
		dj.State.Log.Printf("Disconnected from server. Retrying connection every %d seconds %d times.\n",
			viper.GetInt("connection.retryinterval"),
			viper.GetInt("connection.retryattempts"))

		success := false
		for retries := 0; retries < viper.GetInt("connection.retryattempts"); retries++ {
			dj.State.Log.Println("Retrying connection...")
			if err := dj.State.Client.Connect(); err == nil {
				dj.State.Log.Println("Successfully reconnected to the server!")
				success = true
				break
			}
			time.Sleep(time.Duration(viper.GetInt("connection.retryinterval")) * time.Second)
		}
		if !success {
			dj.KeepAlive <- true
			dj.State.Log.Fatalln("Could not reconnect to server. Exiting...")
		}
	} else {
		dj.KeepAlive <- true
		dj.State.Log.Fatalln("Disconnected from server. No connection retries will be attempted.")
	}
}

// OnTextMessage event. Checks for command prefix and passes it to the CommandExecutor if it
// exists. Ignores the incoming message otherwise.
func (dj *MumbleDJ) OnTextMessage(e *gumble.TextMessageEvent) {
	plainMessage := gumbleutil.PlainText(&e.TextMessage)
	if len(plainMessage) != 0 {
		if plainMessage[0] == viper.GetString("general.commandprefix")[0] && plainMessage != viper.GetString("general.commandprefix") {
			state, message, isPrivateMessage, err := dj.Commander.FindAndExecuteCommand(dj.State, e.Sender, plainMessage[1:])
			if state != nil {
				dj.State = state
			}
			if err != nil {
				dj.SendPrivateMessage(e.Sender, fmt.Sprintf("An error occurred while executing your command: %s", err.Error()))
			} else {
				if isPrivateMessage {
					dj.SendPrivateMessage(e.Sender, message)
				} else {
					dj.State.Client.Self.Channel.Send(message, false)
				}
			}
		}
	}
}

// OnUserChange event. Checks UserChange type and adjusts skip counts to reflect the
// current status of the users on the server.
func (dj *MumbleDJ) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeDisconnected) || e.Type.Has(gumble.UserChangeChannel) {
		dj.State.Skips.RemoveTrackSkip(e.User)
		dj.State.Skips.RemovePlaylistSkip(e.User)
	}
}

// SendPrivateMessage sends a private message to a user. This method verifies that the targeted
// user is still present in the server before attempting to send the message.
func (dj *MumbleDJ) SendPrivateMessage(user *gumble.User, message string) {
	if targetUser := dj.State.Client.Self.Channel.Users.Find(user.Name); targetUser != nil {
		targetUser.Send(message)
	}
}

// CheckAPIKeys checks the configuration for API keys. If no API keys are found
// an error is returned as MumbleDJ will not do anything without API keys.
func (dj *MumbleDJ) CheckAPIKeys() {
	anyDisabled := false

	// Check for YouTube API key.
	if viper.GetString("api.youtubekey") == "" {
		anyDisabled = true
		dj.State.Log.Println("The YouTube service has been disabled as you do not have a YouTube API key defined in your config file.")
	} else {
		dj.State.Handler.AddService("YouTube")
	}

	// Checks for SoundCloud API key.
	if viper.GetString("api.soundcloudkey") == "" {
		anyDisabled = true
		dj.State.Log.Println("The SoundCloud service has been disabled as you do not have a SoundCloud API key defined in your config file.")
	} else {
		dj.State.Handler.AddService("SoundCloud")
	}

	// Check to see if any service was disabled. If so, display a help message.
	if anyDisabled {
		dj.State.Log.Println("Please see the following link for information on how to enable missing services: https://github.com/matthieugrieger/mumbledj.")
	}

	// Exits application if no services are enabled.
	if len(dj.State.Handler.GetAvailableServices()) == 0 {
		dj.State.Log.Fatalln("No services are enabled, meaning MumbleDJ cannot do anything. Exiting...")
	}
}

func main() {
	// Initialize MumbleDJ and its data structures.
	dj := new(MumbleDJ)

	dj.Commander = commands.NewCommander()

	dj.State = new(state.BotState)
	dj.State.BotConfig = new(state.Config)
	dj.State.Queue = state.NewAudioQueue()
	dj.State.Cache = state.NewAudioCache()
	dj.State.Skips = state.NewSkipTracker()
	dj.State.Handler = new(audio.ServiceHandler)
	dj.State.Log = log.New(os.Stderr, "MumbleDJ", 0)

	// Initialize MumbleDJ config with default values and values provided by the user.
	dj.State.BotConfig.SetDefaultConfiguration()
	dj.State.BotConfig.LoadFromCommandline()
	dj.State.BotConfig.LoadFromConfigFile("")

	// Create Gumble config.
	dj.State.GumbleConfig = &gumble.Config{
		Username: viper.GetString("connection.username"),
		Password: viper.GetString("connection.password"),
		Address:  fmt.Sprintf("%s:%d", viper.GetString("connection.address"), viper.GetInt("connection.port")),
		Tokens:   viper.GetStringSlice("connection.accesstokens"),
	}

	// Create Gumble client.
	dj.State.Client = gumble.NewClient(dj.State.GumbleConfig)

	// Initialize key pair if needed.
	dj.State.GumbleConfig.TLSConfig.InsecureSkipVerify = true
	if !viper.GetBool("connection.insecure") {
		gumbleutil.CertificateLockFile(dj.State.Client, fmt.Sprintf("%s/.mumbledjcert.lock", os.Getenv("HOME")))
	}
	if viper.GetString("connection.cert") != "" {
		if viper.GetString("connection.key") == "" {
			viper.Set("connection.key", viper.GetString("connection.cert"))
		}
		if certificate, err := tls.LoadX509KeyPair(viper.GetString("connection.cert"), viper.GetString("connection.key")); err != nil {
			panic(err)
		} else {
			dj.State.GumbleConfig.TLSConfig.Certificates = append(dj.State.GumbleConfig.TLSConfig.Certificates, certificate)
		}
	}

	dj.CheckAPIKeys()

	dj.State.Client.Attach(gumbleutil.Listener{
		Connect:     dj.OnConnect,
		Disconnect:  dj.OnDisconnect,
		TextMessage: dj.OnTextMessage,
		UserChange:  dj.OnUserChange,
	})
	dj.State.Client.Attach(gumbleutil.AutoBitrate)

	if err := dj.State.Client.Connect(); err != nil {
		dj.State.Log.Fatalf("Could not connect to Mumble server at %s:%d.\n", viper.GetString("connection.address"), viper.GetInt("connection.port"))
	}

	<-dj.KeepAlive
}
