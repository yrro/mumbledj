/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/help.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// HelpCommand is a command that outputs a help message that shows the available
// commands and their aliases.
type HelpCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *HelpCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.help")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *HelpCommand) IsAdmin() bool {
	return viper.GetBool("permissions.help")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *HelpCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
