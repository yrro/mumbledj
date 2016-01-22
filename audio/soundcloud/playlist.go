/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/soundcloud/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package soundcloud

import "github.com/layeh/gumble/gumble"

// Playlist is a struct that contains metadata related to a SoundCloud playlist.
type Playlist struct {
	ID        string
	Author    string
	Title     string
	Submitter string
	Tracks    []*Track
}

// NewPlaylist returns a SoundCloud playlist struct with the provided ID.
func NewPlaylist(id string, submitter *gumble.User) *Playlist {
	return &Playlist{
		ID:        id,
		Submitter: submitter.Name,
	}
}

// FetchMetadata makes an API call to the SoundCloud API to fill in the metadata
// about this particular playlist.
func (p *Playlist) FetchMetadata() error {
	return nil
}

// GetID returns the ID of the SoundCloud playlist.
func (p *Playlist) GetID() string {
	return p.ID
}

// GetAuthor returns the author of the SoundCloud playlist.
func (p *Playlist) GetAuthor() string {
	return p.Author
}

// GetTitle returns the title of the SoundCloud playlist.
func (p *Playlist) GetTitle() string {
	return p.Title
}

// GetSubmitter returns the name of the user who submitted the playlist.
func (p *Playlist) GetSubmitter() string {
	return p.Submitter
}

// GetService returns the service name of the SoundCloud playlist.
func (p *Playlist) GetService() string {
	return "SoundCloud"
}
