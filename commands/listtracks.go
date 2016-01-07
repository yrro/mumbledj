/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/listtracks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ListTracksCommand is a command that lists the tracks that are currently in the queue.
type ListTracksCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ListTracksCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.listtracks")
}

// Description returns a description of the command.
func (c *ListTracksCommand) Description() string {
	return viper.GetString("descriptions.listtracks")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ListTracksCommand) IsAdmin() bool {
	return viper.GetBool("permissions.listtracks")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ListTracksCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("There are no tracks currently in the queue.")
	}

	numTracksToList := len(state.Queue.Queue)
	if len(args) != 0 {
		if parsedNum, err := strconv.Atoi(args[0]); err == nil {
			numTracksToList = parsedNum
		}
	}

	var buffer bytes.Buffer
	state.Queue.Traverse(func(i int, track audio.Track) {
		if i < numTracksToList {
			buffer.WriteString(fmt.Sprintf("%d: \"%s\", added by <b>%s</b>.</br>", i+1, track.Title(), track.Submitter()))
		}
	})

	return nil, buffer.String(), true, nil
}
