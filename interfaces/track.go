/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/track.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package interfaces

import "time"

// Track is an interface that includes methods that perform necessary actions
// on an audio track.
type Track interface {
	Download() error
	Play()
	Delete() error
	Submitter() string
	Title() string
	ID() string
	Filename() string
	Duration() time.Duration
	Thumbnail() string
	Playlist() Playlist
}
