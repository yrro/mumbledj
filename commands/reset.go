/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/reset.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// ResetCommand is a command that resets the audio queue.
type ResetCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewResetCommand returns a new ResetCommand object.
func NewResetCommand(aliases []string, isAdmin bool) *ResetCommand {
	return &ResetCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *ResetCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ResetCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ResetCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
