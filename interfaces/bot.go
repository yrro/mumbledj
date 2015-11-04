/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/bot.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package interfaces

import "github.com/layeh/gumble/gumble"

// Bot is an interface that defines all of the available methods for interacting
// with the bot.
type Bot interface {
	AddSkip(user *gumble.User, skipType int)
	RemoveSkip(user *gumble.User, skipType int)
	ResetSkips(skipType int)
}
