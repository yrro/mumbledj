/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/add.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// AddCommand is a command that adds an audio track associated with a supported
// URL to the audio queue.
type AddCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *AddCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.add")
}

// Description returns a description of the command.
func (c *AddCommand) Description() string {
	return viper.GetString("descriptions.add")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *AddCommand) IsAdmin() bool {
	return viper.GetBool("permissions.add")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *AddCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	var allTracks []audio.Track

	if len(args) == 0 {
		return nil, "", true, errors.New("A URL must be supplied with the add command.")
	}

	for _, arg := range args {
		tracks, err := state.Handler.GetTracks(arg)
		if err == nil {
			allTracks = append(allTracks, tracks...)
		}
	}

	if len(allTracks) == 0 {
		return nil, "", true, errors.New("No valid audio tracks were found with the provided URL(s).")
	}

	addString := fmt.Sprintf("<b>%s</b> added <b>%d</b> tracks to the queue:</br>", user.Name, len(allTracks))

	for _, track := range allTracks {
		state.Queue.AddTrack(track)
		addString += fmt.Sprintf("\"%s\" from %s</br>", track.Title(), track.Service())
	}

	return state, addString, false, nil
}
