/*
 * MumbleDJ
 * By Matthieu Grieger
 * constants.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

// SkipType represents either a current track skip or a current playlist skip.
type SkipType int

const (
	// CurrentTrackSkipType maps to the skiplist for the current track.
	CurrentTrackSkipType SkipType = iota
	// CurrentPlaylistSkipType maps to the skiplist for the current playlist.
	CurrentPlaylistSkipType SkipType = iota
)
