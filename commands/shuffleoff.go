/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffleoff.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ShuffleOffCommand is a command that turns off automatic audio queue shuffling.
type ShuffleOffCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ShuffleOffCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.shuffleoff")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ShuffleOffCommand) IsAdmin() bool {
	return viper.GetBool("permissions.shuffleoff")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ShuffleOffCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if viper.GetBool("general.automaticshuffleon") {
		viper.Set("general.automaticshuffleon", false)
		return nil, "Automatic shuffling has been toggled off.", false, nil
	}

	return nil, "", true, errors.New("Automatic shuffling is already toggled off.")
}
