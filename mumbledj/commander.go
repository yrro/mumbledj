/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj/commander.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package mumbledj

import (
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/matthieugrieger/mumbledj/objects"
)

// Commander is a struct that holds all available commands and provides
// methods that allow interactions with said commands.
type Commander struct {
	Commands []*interfaces.Command
}

// NewCommander returns a new commander with an initialized command list.
func NewCommander(config *objects.AliasConfig) *Commander {
	commands := []*interfaces.Command{}

	return &Commander{
		Commands: commands,
	}
}
