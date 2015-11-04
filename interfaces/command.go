/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/command.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package interfaces

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
)

// Command is an interface that all commands must implement.
type Command interface {
	Aliases() []string
	IsAdmin() bool
	Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error)
}
