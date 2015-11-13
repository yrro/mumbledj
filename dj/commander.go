/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/commander.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"errors"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/commands"
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// Commander is a struct that holds all available commands and provides
// methods that allow interactions with said commands.
type Commander struct {
	Commands []interfaces.Command
}

// NewCommander returns a new commander with an initialized command list.
func NewCommander() *Commander {
	commands := []interfaces.Command{
		new(commands.AddCommand),
		new(commands.CacheSizeCommand),
		new(commands.CurrentTrackCommand),
		new(commands.HelpCommand),
		new(commands.KillCommand),
		new(commands.MoveCommand),
		new(commands.NextTrackCommand),
		new(commands.NumCachedCommand),
		new(commands.NumTracksCommand),
		new(commands.ReloadCommand),
		new(commands.ResetCommand),
		new(commands.SetCommentCommand),
		new(commands.ShuffleCommand),
		new(commands.SkipCommand),
		new(commands.VolumeCommand),
	}

	return &Commander{
		Commands: commands,
	}
}

// FindAndExecuteCommand attempts to find a reference to a command in an incoming message.
// If a command is found the command object is returned.
func (c *Commander) FindAndExecuteCommand(currentState *state.BotState, user *gumble.User, message string) (*state.BotState, string, error) {
	possibleCommand := strings.ToLower(message[0:strings.Index(message, " ")])
	for _, command := range c.Commands {
		for _, alias := range command.Aliases() {
			if possibleCommand == alias {
				return c.executeCommand(currentState, user, message, command)
			}
		}
	}
	return nil, "", errors.New("No matching command found.")
}

func (c *Commander) executeCommand(currentState *state.BotState, user *gumble.User, message string, command interfaces.Command) (*state.BotState, string, error) {
	var canExecute bool
	if viper.GetBool("permissions.adminsenabled") && command.IsAdmin() {
		for _, username := range viper.GetStringSlice("permissions.admins") {
			if user.Name == username {
				canExecute = true
			}
		}
		canExecute = false
	} else {
		canExecute = true
	}

	if canExecute {
		return command.Execute(currentState, user, strings.Split(message, " ")[1:]...)
	}
	return nil, "", errors.New("You do not have permission to execute this command.")
}
