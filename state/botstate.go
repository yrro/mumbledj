/*
 * MumbleDJ
 * By Matthieu Grieger
 * state/botstate.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import (
	"log"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleffmpeg"
)

// BotState is a struct that allows for important pieces of the bot's state
// to be passed back and forth between commands.
type BotState struct {
	Client       *gumble.Client
	GumbleConfig *gumble.Config
	AudioStream  *gumbleffmpeg.Stream
	Queue        *AudioQueue
	Cache        *AudioCache
	Skips        *SkipTracker
	Log          log.Logger
}
