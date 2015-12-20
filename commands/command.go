/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/command.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
)

// Command is an interface that all commands must implement.
type Command interface {
	Aliases() []string
	Description() string
	IsAdmin() bool
	Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error)
}
