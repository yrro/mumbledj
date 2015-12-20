/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/move.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// MoveCommand is a command that moves the bot from one channel to another.
type MoveCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *MoveCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.move")
}

// Description returns a description of the command.
func (c *MoveCommand) Description() string {
	return viper.GetString("descriptions.move")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *MoveCommand) IsAdmin() bool {
	return viper.GetBool("permissions.move")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *MoveCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(args) == 0 {
		return nil, "", true, errors.New("A destination channel must be supplied to move the bot.")
	}

	if channels := strings.Split(args[0], "/"); state.Client.Channels.Find(channels...) != nil {
		state.Client.Self.Move(state.Client.Channels.Find(channels...))
	} else {
		return nil, "", true, errors.New("The provided channel does not exist.")
	}

	return state, fmt.Sprintf("You have successfully moved the bot to <b>%s</b>.", args[0]), true, nil
}
