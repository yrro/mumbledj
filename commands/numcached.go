/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numcached.go
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

// NumCachedCommand is a command that outputs the number of audio tracks that
// are currently cached on disk.
type NumCachedCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *NumCachedCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.numcached")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *NumCachedCommand) IsAdmin() bool {
	return viper.GetBool("permissions.numcached")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *NumCachedCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if !viper.GetBool("cache.enabled") {
		return nil, "", true, errors.New("Caching is currently disabled.")
	}

	state.Cache.UpdateStats()
	return nil, fmt.Sprintf("There are currently %d items stored in the cache.", state.Cache.NumAudioFiles), true, nil
}
