/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/mumbledj.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"strings"
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
			command, err := dj.Commander.FindCommand(plainMessage[1:])
			if err == nil {
				command.Execute(dj.State, e.Sender, strings.Split(plainMessage, " ")[1:]...)
			}
		}
	}
}

// OnUserChange event. Checks UserChange type and adjusts skip counts to reflect the
// current status of the users on the server.
func (dj *MumbleDJ) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeDisconnected) || e.Type.Has(gumble.UserChangeChannel) {
		dj.RemoveSkip(e.User, 0)
		dj.RemoveSkip(e.User, 1)
	}
}

// AddSkip adds the username of a user that has skipped the current track/playlist to the
// desired skiplist.
func (dj *MumbleDJ) AddSkip(user *gumble.User, skipType int) {
	for _, username := range dj.State.Skips[skipType] {
		if username == user.Name {
			dj.State.Log.Printf("%s has already skipped this item.\n", user.Name)
		}
	}
	dj.State.Skips[skipType] = append(dj.State.Skips[skipType], user.Name)
	dj.State.Log.Printf("%s's skip has been successfully added for this item.\n", user.Name)
}

// RemoveSkip removes the username of a user who has previously skipped the current track/playlist
// from the desired skiplist.
func (dj *MumbleDJ) RemoveSkip(user *gumble.User, skipType int) {
	for i, username := range dj.State.Skips[skipType] {
		if username == user.Name {
			dj.State.Skips[skipType] = append(dj.State.Skips[skipType][:i], dj.State.Skips[skipType][i+1:]...)
			dj.State.Log.Printf("%s's skip has been successfully removed from this item.\n", user.Name)
			return
		}
	}
	dj.State.Log.Printf("%s never skipped this item.\n", user.Name)
}

// ResetSkips resets the skiplist for either the current track or the current playlist.
func (dj *MumbleDJ) ResetSkips(skipType int) {
	dj.State.Skips[skipType] = dj.State.Skips[skipType][:0]
}

// HasPermission checks if a particular user has the necessary permissions to execute a command.
// Permissions are specified in the user configuration if it exists.
func (dj *MumbleDJ) HasPermission(user *gumble.User, isAdminCommand bool) bool {
	if viper.GetBool("permissions.adminsenabled") && isAdminCommand {
		for _, username := range viper.GetStringSlice("permissions.admins") {
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
	if targetUser := dj.State.Client.Self.Channel.Users.Find(user.Name); targetUser != nil {
		targetUser.Send(message)
	}
}

// Start attempts to connect the bot to the server and performs necessary startup
// operations.
func (dj *MumbleDJ) Start() error {
	return nil
}
