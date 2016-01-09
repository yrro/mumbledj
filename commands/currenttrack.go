/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/currenttrack.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// CurrentTrackCommand is a command that outputs the track being currently played (if exists).
type CurrentTrackCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *CurrentTrackCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.currenttrack")
}

// Description returns a description of the command.
func (c *CurrentTrackCommand) Description() string {
	return viper.GetString("descriptions.currenttrack")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *CurrentTrackCommand) IsAdmin() bool {
	return viper.GetBool("permissions.currenttrack")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *CurrentTrackCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("There are no tracks in the queue.")
	}

	currentTrack := state.Queue.Queue[0]

	return nil, fmt.Sprintf("The current track is <b>%s</b>, added by <b>%s</b>.", currentTrack.GetTitle(), currentTrack.GetSubmitter()), true, nil
}
