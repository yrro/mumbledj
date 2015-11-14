/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/reload.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ReloadCommand is a command that reloads the configuration values for the bot.
type ReloadCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ReloadCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.reload")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ReloadCommand) IsAdmin() bool {
	return viper.GetBool("permissions.reload")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ReloadCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
