/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/move.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// MoveCommand is a command that moves the bot from one channel to another.
type MoveCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewMoveCommand returns a new MoveCommand object.
func NewMoveCommand(aliases []string, isAdmin bool) *MoveCommand {
	return &MoveCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *MoveCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *MoveCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *MoveCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
