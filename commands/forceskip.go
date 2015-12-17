/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/forceskip.go
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

// ForceSkipCommand is a command that immediately skips the current audio track.
type ForceSkipCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ForceSkipCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.forceskip")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ForceSkipCommand) IsAdmin() bool {
	return viper.GetBool("permissions.forceskip")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ForceSkipCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("The queue is currently empty. There are no tracks to skip.")
	}

	state.Queue.Skip()

	return state, fmt.Sprintf("The current track has been forcibly skipped by <b>%s</b>.", user.Name), false, nil
}
