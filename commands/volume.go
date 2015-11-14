/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/volume.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// VolumeCommand is a command that changes the volume of the audio output.
type VolumeCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *VolumeCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.volume")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *VolumeCommand) IsAdmin() bool {
	return viper.GetBool("permissions.volume")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *VolumeCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	return nil, "", nil
}
