/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/volume.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// VolumeCommand is a command that changes the volume of the audio output.
type VolumeCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewVolumeCommand returns a new VolumeCommand object.
func NewVolumeCommand(aliases []string, isAdmin bool) *VolumeCommand {
	return &VolumeCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the add command.
func (c *VolumeCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *VolumeCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *VolumeCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
