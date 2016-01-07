/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/skip.go
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

// SkipCommand is a command that places a vote to skip the current audio track.
type SkipCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *SkipCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.skip")
}

// Description returns a description of the command.
func (c *SkipCommand) Description() string {
	return viper.GetString("descriptions.skip")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SkipCommand) IsAdmin() bool {
	return viper.GetBool("permissions.skip")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SkipCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("The queue is currently empty. There is no track to skip.")
	}

	if err := state.Skips.AddTrackSkip(user); err != nil {
		return nil, "", true, errors.New("You have already voted to skip this track.")
	}

	return state, fmt.Sprintf("<b>%s</b> has voted to skip the current track.", user.Name), false, nil
}
