/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/currenttrack.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
	"github.com/spf13/viper"
)

// CurrentTrackCommand is a command that outputs the track being currently played (if exists).
type CurrentTrackCommand struct{}

// Aliases is a method that returns the current aliases for the add command.
func (c *CurrentTrackCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.currenttrack")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *CurrentTrackCommand) IsAdmin() bool {
	return viper.GetBool("permissions.currenttrack")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *CurrentTrackCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
