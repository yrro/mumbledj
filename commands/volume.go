/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/volume.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// VolumeCommand is a command that changes the volume of the audio output.
type VolumeCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *VolumeCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.volume")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *VolumeCommand) IsAdmin() bool {
	return viper.GetBool("permissions.volume")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *VolumeCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, error) {
	if len(args) == 0 {
		// Send the user the current volume level.
		return nil, fmt.Sprintf("The current volume is <b>%.2f</b>.", state.AudioStream.Volume), nil
	}
	if len(args) > 1 {
		return nil, "", errors.New("The volume command only accepts 0-1 arguments.")
	}

	newVolume, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, "", errors.New("An error occurred while parsing the requested volume.")
	}

	if newVolume < viper.GetFloat64("volume.lowest") || newVolume > viper.GetFloat64("volume.highest") {
		return nil, "", fmt.Errorf("Volumes must be between the values <b>%.2f</b> and <b>%.2f</b>",
			viper.GetFloat64("volume.lowest"), viper.GetFloat64("volume.highest"))
	}

	state.AudioStream.Volume = float32(newVolume)
	return state, fmt.Sprintf("<b>%s</b> has changed the volume to <b>%.2f</b>.", user.Name, newVolume), nil
}
