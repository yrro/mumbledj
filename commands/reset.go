/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/reset.go
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

// ResetCommand is a command that resets the audio queue.
type ResetCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ResetCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.reset")
}

// Description returns a description of the command.
func (c *ResetCommand) Description() string {
	return viper.GetString("descriptions.reset")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ResetCommand) IsAdmin() bool {
	return viper.GetBool("permissions.reset")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ResetCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("The queue is already empty.")
	}

	state.Queue.Queue = state.Queue.Queue[:0]
	if state.AudioStream != nil {
		state.AudioStream.Stop()
	}

	if err := state.Cache.DeleteAll(); err != nil {
		return nil, "", true, err
	}

	return state, fmt.Sprintf("<b>%s</b> has reset the audio queue.", user.Name), false, nil
}
