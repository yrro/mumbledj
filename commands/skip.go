/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/skip.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// SkipCommand is a command that places a vote to skip the current audio track.
type SkipCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewSkipCommand returns a new SkipCommand object.
func NewSkipCommand(aliases []string, isAdmin bool) *SkipCommand {
	return &SkipCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *SkipCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SkipCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SkipCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
