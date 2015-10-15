/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"log"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumble_ffmpeg"
	"github.com/layeh/gumble/gumbleutil"
	"github.com/matthieugrieger/mumbledj/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// MumbleDJ is a struct that keeps track of all aspects of the bot's state.
type MumbleDJ struct {
	GumbleConfig gumble.Config
	Client       *gumble.Client
	KeepAlive    chan bool
	Queue        *AudioQueue
	AudioStream  *gumble_ffmpeg.Stream
	Skips        map[SkipType][]string
	Cache        *AudioCache
	Log          log.Logger
	Command      *commands.CommandExecutor
	Args         *cobra.Command
}

// OnConnect event. First moves MumbleDJ into default channel if one exists. The
// configuration is loaded and the audio stream is set up.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {
	dj.Client.Self.Move(dj.Client.Channels.Find(viper.GetString("DefaultChannel")))

	dj.AudioStream = gumble_ffmpeg.New(dj.Client)
	dj.AudioStream.Volume = float32(viper.GetFloat64("DefaultVolume"))

	dj.Client.Self.SetComment(viper.GetString("DefaultComment"))

	if viper.GetBool("CacheEnabled") {
		dj.Cache.Update()
		go dj.Cache.ClearExpired()
	}
}

// OnDisconnect event. Terminates MumbleDJ thread, or retries connection if
// automatic connection retries are enabled.
func (dj *MumbleDJ) OnDisconnect(e *gumble.DisconnectEvent) {
	if viper.GetBool("RetryConnectionEnabled") && (e.Type == gumble.DisconnectError || e.Type == gumble.DisconnectKicked) {
		dj.Log.Printf("Disconnected from server. Retrying connection every %d seconds %d times.\n",
			viper.GetInt("RetryInterval"),
			viper.GetInt("RetryAttempts"))

		success := false
		for retries := 0; retries <= viper.GetInt("RetryAttempts"); retries++ {
			dj.Log.Println("Retrying connection...")
			if err := dj.Client.Connect(); err == nil {
				dj.Log.Println("Successfully reconnected to the server!")
				success = true
				break
			}
			time.Sleep(time.Duration(viper.GetInt("RetryInterval")) * time.Second)
		}
		if !success {
			dj.KeepAlive <- true
			dj.Log.Fatalln("Could not reconnect to server. Exiting...")
		}
	} else {
		dj.KeepAlive <- true
		dj.Log.Fatalln("Disconnected from server. No connection retries will be attempted.")
	}
}

// OnTextMessage event. Checks for command prefix and passes it to the CommandExecutor if it
// exists. Ignores the incoming message otherwise.
func (dj *MumbleDJ) OnTextMessage(e *gumble.TextMessageEvent) {
	plainMessage := gumbleutil.PlainText(&e.TextMessage)
	if len(plainMessage) != 0 {
		if plainMessage[0] == viper.GetString("CommandPrefix")[0] && plainMessage != viper.GetString("CommandPrefix") {
			dj.Command.Execute(e.Sender, plainMessage[1:])
		}
	}
}

// OnUserChange event. Checks UserChange type and adjusts skip counts to reflect the
// current status of the users on the server.
func (dj *MumbleDJ) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeDisconnected) || e.Type.Has(gumble.UserChangeChannel) {
		dj.RemoveSkip(e.User, CurrentTrackSkipType)
		dj.RemoveSkip(e.User, CurrentPlaylistSkipType)
	}
}

// AddSkip adds the username of a user that has skipped the current track/playlist to the
// desired skiplist.
func (dj *MumbleDJ) AddSkip(user *gumble.User, skipType SkipType) {
	for _, username := range dj.Skips[skipType] {
		if username == user.Name {
			dj.Log.Printf("%s has already skipped this item.\n", user.Name)
		}
	}
	dj.Skips[skipType] = append(dj.Skips[skipType], user.Name)
	dj.Log.Printf("%s's skip has been successfully added for this item.\n", user.Name)
}

// RemoveSkip removes the username of a user who has previously skipped the current track/playlist
// from the desired skiplist.
func (dj *MumbleDJ) RemoveSkip(user *gumble.User, skipType SkipType) {
	for i, username := range dj.Skips[skipType] {
		if username == user.Name {
			dj.Skips[skipType] = append(dj.Skips[skipType][:i], dj.Skips[skipType][i+1]...)
			dj.Log.Printf("%s's skip has been successfully removed from this item.\n", user.Name)
			return
		}
	}
	dj.Log.Printf("%s never skipped this item.\n", user.Name)
}

// ResetSkips resets the skiplist for either the current track or the current playlist.
func (dj *MumbleDJ) ResetSkips(skipType SkipType) {
	dj.Skips[skipType] = dj.Skips[skipType][:0]
}

// HasPermission checks if a particular user has the necessary permissions to execute a command.
// Permissions are specified in the user configuration if it exists.
func (dj *MumbleDJ) HasPermission(user *gumble.User, isAdminCommand bool) bool {
	if viper.GetBool("AdminsEnabled") && isAdminCommand {
		for _, username := range viper.GetStringSlice("Admins") {
			if username == user.Name {
				return true
			}
		}
		return false
	}
	return true
}

// SendPrivateMessage sends a private message to a user. This method verifies that the targeted
// user is still present in the server before attempting to send the message.
func (dj *MumbleDJ) SendPrivateMessage(user *gumble.User, message string) {
	if targetUser := dj.Client.Self.Channel.Users.Find(user.Name); targetUser != nil {
		targetUser.Send(message)
	}
}

// Start attempts to connect the bot to the server and performs necessary startup
// operations.
func (dj *MumbleDJ) Start() error {

}
