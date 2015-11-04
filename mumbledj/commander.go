/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj/commander.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package mumbledj

import (
	"github.com/matthieugrieger/mumbledj/commands"
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
	commands := []*interfaces.Command{
		&commands.AddCommand,
		&commands.CacheSizeCommand,
		&commands.CurrentTrackCommand,
		&commands.HelpCommand,
		&commands.KillCommand,
		&commands.MoveCommand,
		&commands.NextTrackCommand,
		&commands.NumCachedCommand,
		&commands.NumTracksCommand,
		&commands.ReloadCommand,
		&commands.ResetCommand,
		&commands.SetCommentCommand,
		&commands.ShuffleCommand,
		&commands.SkipCommand,
		&commands.VolumeCommand,
	}

	return &Commander{
		Commands: commands,
	}
}
