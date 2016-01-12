/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/track.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package audio

import "time"

// Track is an interface that includes methods that perform necessary actions
// on an audio track.
type Track interface {
	FetchMetadata() error
	Download() error
	Delete() error
	GetSubmitter() string
	GetAuthor() string
	GetTitle() string
	GetID() string
	GetFilename() string
	GetDuration() time.Duration
	GetThumbnail() string
	GetService() string
	GetPlaylist() Playlist
}
