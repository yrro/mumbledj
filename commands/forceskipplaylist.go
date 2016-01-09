/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/forceskipplaylist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ForceSkipPlaylistCommand is a command that immediately skips the current playlist.
type ForceSkipPlaylistCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ForceSkipPlaylistCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.forceskipplaylist")
}

// Description returns a description of the command.
func (c *ForceSkipPlaylistCommand) Description() string {
	return viper.GetString("descriptions.forceskipplaylist")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ForceSkipPlaylistCommand) IsAdmin() bool {
	return viper.GetBool("permissions.forceskipplaylist")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ForceSkipPlaylistCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("The queue is currently empty. There are no tracks to skip.")
	}

	if state.Queue.Queue[0].GetPlaylist() == nil {
		return nil, "", true, errors.New("The current track is not part of a playlist.")
	}

	state.Queue.SkipPlaylist()

	return state, fmt.Sprintf("The current playlist has been forcibly skipped by <b>%s</b>.", user.Name), false, nil
}
