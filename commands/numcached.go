/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numcached.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// NumCachedCommand is a command that outputs the number of audio tracks that
// are currently cached on disk.
type NumCachedCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewNumCachedCommand returns a new NumCachedCommand object.
func NewNumCachedCommand(aliases []string, isAdmin bool) *NumCachedCommand {
	return &NumCachedCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *NumCachedCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NumCachedCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NumCachedCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
