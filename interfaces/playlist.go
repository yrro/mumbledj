/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package interfaces

// Playlist is an interface that represents all valid playlists of multiple
// audio tracks.
type Playlist interface {
	ID() string
	Title() string
}
