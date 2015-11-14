/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/skipplaylist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
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

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SkipPlaylistCommand) IsAdmin() bool {
	return viper.GetBool("permissions.skipplaylist")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SkipPlaylistCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
