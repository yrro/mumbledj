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
)

// CacheSizeCommand is a command that outputs the current size of the cache.
type CacheSizeCommand struct {
	CurrentAliases []string
	IsAdminCommand bool
}

// NewCacheSizeCommand returns a new CacheSizeCommand object.
func NewCacheSizeCommand(aliases []string, isAdmin bool) *CacheSizeCommand {
	return &CacheSizeCommand{
		CurrentAliases: aliases,
		IsAdminCommand: isAdmin,
	}
}

// Aliases is a method that returns the current aliases for the cachesize command.
func (c *CacheSizeCommand) Aliases() []string {
	return c.CurrentAliases
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *CacheSizeCommand) IsAdmin() bool {
	return c.IsAdminCommand
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *CacheSizeCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}
