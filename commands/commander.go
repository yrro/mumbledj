/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/commander.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// Commander is a struct that holds all available commands and provides
// methods that allow interactions with said commands.
type Commander struct {
	Commands []Command
}

// NewCommander returns a new commander with an initialized command list.
func NewCommander() *Commander {
	commands := []Command{
		new(AddCommand),
		new(CacheSizeCommand),
		new(CurrentTrackCommand),
		new(ForceSkipCommand),
		new(ForceSkipPlaylistCommand),
		new(HelpCommand),
		new(KillCommand),
		new(ListTracksCommand),
		new(MoveCommand),
		new(NextTrackCommand),
		new(NumCachedCommand),
		new(NumTracksCommand),
		new(ReloadCommand),
		new(ResetCommand),
		new(SetCommentCommand),
		new(ShuffleCommand),
		new(ShuffleOffCommand),
		new(ShuffleOnCommand),
		new(SkipCommand),
		new(SkipPlaylistCommand),
		new(VolumeCommand),
	}

	return &Commander{
		Commands: commands,
	}
}

// FindAndExecuteCommand attempts to find a reference to a command in an incoming message.
// If a command is found the command object is returned.
func (c *Commander) FindAndExecuteCommand(currentState *state.BotState, user *gumble.User, message string) (*state.BotState, string, bool, error) {
	command, err := c.FindCommand(message)
	if err != nil {
		return nil, "", true, errors.New("No command was found in this message.")
	}

	return c.ExecuteCommand(currentState, user, message, command)
}

// FindCommand returns the command that corresponds with the incoming message.
func (c *Commander) FindCommand(message string) (Command, error) {
	possibleCommand := strings.ToLower(message[0:strings.Index(message, " ")])
	for _, command := range c.Commands {
		for _, alias := range command.Aliases() {
			if possibleCommand == alias {
				return command, nil
			}
		}
	}
	return nil, errors.New("No command was found in this message.")
}

// ExecuteCommand executes the passed command with the corresponding state, user, and message. The message is split by whitespace to make up the arguments
// of a command.
func (c *Commander) ExecuteCommand(currentState *state.BotState, user *gumble.User, message string, command Command) (*state.BotState, string, bool, error) {
	canExecute := false
	if viper.GetBool("permissions.adminsenabled") && command.IsAdmin() {
		for _, username := range viper.GetStringSlice("permissions.admins") {
			if user.Name == username {
				canExecute = true
			}
		}
	} else {
		canExecute = true
	}

	if canExecute {
		return command.Execute(currentState, user, strings.Split(message, " ")[1:]...)
	}
	return nil, "", true, errors.New("You do not have permission to execute this command.")
}
