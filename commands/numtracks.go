/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numtracks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// NumTracksCommand is a command that outputs the current number of tracks in
// the audio queue.
type NumTracksCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewNumTracksCommand returns a new NumTracksCommand object.
func NewNumTracksCommand(aliases []string, isAdmin bool) *NumTracksCommand {
	return &NumTracksCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *NumTracksCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NumTracksCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NumTracksCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
