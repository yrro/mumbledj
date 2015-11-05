/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/constants.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

// SkipType represents either a current track skip or a current playlist skip.
type SkipType int

const (
	// CurrentTrackSkipType maps to the skiplist for the current track.
	CurrentTrackSkipType SkipType = iota
	// CurrentPlaylistSkipType maps to the skiplist for the current playlist.
	CurrentPlaylistSkipType SkipType = iota
)

// BytesInMebibyte is the number of bytes in mebibytes used for cache filesize
// conversions.
const BytesInMebibyte int64 = 1048576
