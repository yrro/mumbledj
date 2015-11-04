/*
 * MumbleDJ
 * By Matthieu Grieger
 * objects/botstate.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package objects

import "github.com/matthieugrieger/mumbledj/interfaces"

// BotState is a struct that allows for important pieces of the bot's state
// to be passed back and forth between commands.
type BotState struct {
	Queue  *interfaces.Queue
	Cache  *interfaces.Cache
	Config *Config
}
