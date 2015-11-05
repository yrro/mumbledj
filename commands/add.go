/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/add.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// AddCommand is a command that adds an audio track associated with a supported
// URL to the audio queue.
type AddCommand struct{}

// Aliases is a method that returns the current aliases for the add command.
func (c *AddCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.add")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *AddCommand) IsAdmin() bool {
	return viper.GetBool("permissions.add")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *AddCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
