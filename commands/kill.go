/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/kill.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// KillCommand is a command that safely kills the bot.
type KillCommand struct{}

// Aliases is a method that returns the current aliases for the add command.
func (c *KillCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.kill")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *KillCommand) IsAdmin() bool {
	return viper.GetBool("permissions.kill")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *KillCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
