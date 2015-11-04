/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/help.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// HelpCommand is a command that outputs a help message that shows the available
// commands and their aliases.
type HelpCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewHelpCommand returns a new HelpCommand object.
func NewHelpCommand(aliases []string, isAdmin bool) *HelpCommand {
	return &HelpCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *HelpCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *HelpCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *HelpCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
