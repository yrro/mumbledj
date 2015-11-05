/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/nexttrack.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// NextTrackCommand is a command that outputs the next track in the audio queue (if exists).
type NextTrackCommand struct{}

// Aliases is a method that returns the current aliases for the add command.
func (c *NextTrackCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.nexttrack")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NextTrackCommand) IsAdmin() bool {
	return viper.GetBool("permissions.nexttrack")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NextTrackCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
