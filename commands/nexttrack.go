/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/nexttrack.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// NextTrackCommand is a command that outputs the next track in the audio queue (if exists).
type NextTrackCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewNextTrackCommand returns a new NextTrackCommand object.
func NewNextTrackCommand(aliases []string, isAdmin bool) *NextTrackCommand {
	return &NextTrackCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *NextTrackCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NextTrackCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NextTrackCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
