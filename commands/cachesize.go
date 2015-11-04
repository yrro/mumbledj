/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/cachesize.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
	"github.com/spf13/viper"
)

// CacheSizeCommand is a command that outputs the current size of the cache.
type CacheSizeCommand struct{}

// Aliases is a method that returns the current aliases for the cachesize command.
func (c *CacheSizeCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.cachesize")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *CacheSizeCommand) IsAdmin() bool {
	return viper.GetBool("permissions.cachesize")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *CacheSizeCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
