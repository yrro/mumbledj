/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffleon.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ShuffleOnCommand is a command that turns on automatic audio queue shuffling.
type ShuffleOnCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ShuffleOnCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.shuffleon")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ShuffleOnCommand) IsAdmin() bool {
	return viper.GetBool("permissions.shuffleon")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ShuffleOnCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
