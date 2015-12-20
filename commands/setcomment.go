/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/setcomment.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// SetCommentCommand is a command that changes the Mumble comment of the bot.
type SetCommentCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *SetCommentCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.setcomment")
}

// Description returns a description of the command.
func (c *SetCommentCommand) Description() string {
	return viper.GetString("descriptions.setcomment")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SetCommentCommand) IsAdmin() bool {
	return viper.GetBool("permissions.setcomment")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SetCommentCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(args) == 0 {
		state.Client.Self.SetComment("")
		return state, "The comment for the bot has been successfully removed.", true, nil
	}

	newComment := ""
	for _, arg := range args {
		newComment += arg + " "
	}
	state.Client.Self.SetComment(newComment)

	return state, fmt.Sprintf("The comment for the bot has been successfully changed to the following: %s", newComment), true, nil
}
