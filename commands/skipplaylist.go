/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/skipplaylist.go
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

// SkipPlaylistCommand is a command that places a vote to skip the current audio track.
type SkipPlaylistCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *SkipPlaylistCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.skipplaylist")
}

// Description returns a description of the command.
func (c *SkipPlaylistCommand) Description() string {
	return viper.GetString("descriptions.skipplaylist")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SkipPlaylistCommand) IsAdmin() bool {
	return viper.GetBool("permissions.skipplaylist")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SkipPlaylistCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("The queue is currently empty. There is no playlist to skip.")
	}

	if state.Queue.Queue[0].GetPlaylist() == nil {
		return nil, "", true, errors.New("The current track is not part of a playlist.")
	}

	if err := state.Skips.AddPlaylistSkip(user); err != nil {
		return nil, "", true, errors.New("You have already voted to skip this playlist.")
	}

	return state, fmt.Sprintf("<b>%s</b> has voted to skip the current playlist.", user.Name), false, nil
}
