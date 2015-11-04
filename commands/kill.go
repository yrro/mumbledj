/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/kill.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// KillCommand is a command that safely kills the bot.
type KillCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewKillCommand returns a new KillCommand object.
func NewKillCommand(aliases []string, isAdmin bool) *KillCommand {
	return &KillCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *KillCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *KillCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *KillCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
