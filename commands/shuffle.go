/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffle.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// ShuffleCommand is a command that shuffles the audio queue.
type ShuffleCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewShuffleCommand returns a new ShuffleCommand object.
func NewShuffleCommand(aliases []string, isAdmin bool) *ShuffleCommand {
	return &ShuffleCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *ShuffleCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ShuffleCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ShuffleCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
