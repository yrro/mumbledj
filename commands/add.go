/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/add.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// AddCommand is a command that adds an audio track associated with a supported
// URL to the audio queue.
type AddCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewAddCommand returns a new AddCommand object.
func NewAddCommand(aliases []string, isAdmin bool) *AddCommand {
	return &AddCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *AddCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *AddCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *AddCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
