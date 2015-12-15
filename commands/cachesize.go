/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/cachesize.go
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

// CacheSizeCommand is a command that outputs the current size of the cache.
type CacheSizeCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *CacheSizeCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.cachesize")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *CacheSizeCommand) IsAdmin() bool {
	return viper.GetBool("permissions.cachesize")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *CacheSizeCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if !viper.GetBool("cache.enabled") {
		return nil, "", true, errors.New("Caching is currently disabled.")
	}

	state.Cache.UpdateStats()
	return nil, fmt.Sprintf("The current size of the cache is %.2f MiB.", state.Cache.TotalFileSize/1048576), true, nil
}
