/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/forceskipplaylist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
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

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ForceSkipPlaylistCommand) IsAdmin() bool {
	return viper.GetBool("permissions.forceskipplaylist")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ForceSkipPlaylistCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	return nil, "", false, nil
}
