/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/mumbledj.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"fmt"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumble_ffmpeg"
	"github.com/layeh/gumble/gumbleutil"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// MumbleDJ is a struct that keeps track of all aspects of the bot's state.
type MumbleDJ struct {
	GumbleConfig gumble.Config
	KeepAlive    chan bool
	Commander    *Commander
	State        *state.BotState
}

// OnConnect event. First moves MumbleDJ into default channel if one exists. The
// configuration is loaded and the audio stream is set up.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {
	dj.State.Client.Self.Move(dj.State.Client.Channels.Find(viper.GetStringSlice("general.defaultchannel")...))

	dj.State.AudioStream = gumble_ffmpeg.New(dj.State.Client)
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
			state, message, err := dj.Commander.FindAndExecuteCommand(dj.State, e.Sender, plainMessage[1:])
			if state != nil {
				dj.State = state
			}
			if err != nil {
				dj.SendPrivateMessage(e.Sender, fmt.Sprintf("An error occurred while executing your command: %s", err.Error()))
			} else {
				dj.State.Client.Self.Channel.Send(message, false)
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

// Start attempts to connect the bot to the server and performs necessary startup
// operations.
func (dj *MumbleDJ) Start() error {
	return nil
}
