/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/skip.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// SkipCommand is a command that places a vote to skip the current audio track.
type SkipCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *SkipCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.skip")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SkipCommand) IsAdmin() bool {
	return viper.GetBool("permissions.skip")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SkipCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	return nil, "", false, nil
}
