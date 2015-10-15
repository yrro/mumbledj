/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/command.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
)

// Command is an interface that all commands must implement.
type Command interface {
	Execute(user *gumble.User, args ...string) (string, error)
}
