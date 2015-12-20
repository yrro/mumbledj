/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numtracks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// NumTracksCommand is a command that outputs the current number of tracks in
// the audio queue.
type NumTracksCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *NumTracksCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.numtracks")
}

// Description returns a description of the command.
func (c *NumTracksCommand) Description() string {
	return viper.GetString("descriptions.numtracks")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NumTracksCommand) IsAdmin() bool {
	return viper.GetBool("permissions.numtracks")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NumTracksCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 1 {
		return nil, "There is currently <b>1</b> track in the queue.", true, nil
	}

	return nil, fmt.Sprintf("There are currently <b>%d</b> tracks in the queue.", len(state.Queue.Queue)), true, nil
}
