/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/help.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// HelpCommand is a command that outputs a help message that shows the available
// commands and their aliases.
type HelpCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *HelpCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.help")
}

// Description returns a description of the command.
func (c *HelpCommand) Description() string {
	return viper.GetString("descriptions.help")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *HelpCommand) IsAdmin() bool {
	return viper.GetBool("permissions.help")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *HelpCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	commandString := "<b>%s</b> -- %s</br>"
	regularCommands := ""
	adminCommands := ""
	totalString := ""
	commander := NewCommander()

	for _, command := range commander.Commands {
		currentString := fmt.Sprintf(commandString, command.Aliases(), command.Description())
		if command.IsAdmin() {
			adminCommands += currentString
		} else {
			regularCommands += currentString
		}
	}

	totalString = "<b>Commands:</b></br>" + regularCommands

	isAdmin := false
	if viper.GetBool("permissions.adminsenabled") {
		for _, username := range viper.GetStringSlice("permissions.admins") {
			if user.Name == username {
				isAdmin = true
			}
		}
	} else {
		isAdmin = true
	}

	if isAdmin {
		totalString += "</br><b>Admin Commands:</b></br>" + adminCommands
	}

	return nil, totalString, true, nil
}
