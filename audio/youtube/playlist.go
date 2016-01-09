/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/youtube/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package youtube

import "github.com/layeh/gumble/gumble"

// Playlist is a struct that contains metadata related to a YouTube playlist.
type Playlist struct {
	ID        string
	Author    string
	Title     string
	Submitter string
}

// NewPlaylist returns a YouTube playlist struct with the provided ID and
// title.
func NewPlaylist(id, title string, submitter *gumble.User) *Playlist {
	return &Playlist{
		ID:        id,
		Title:     title,
		Submitter: submitter.Name,
	}
}

// FetchMetadata makes an API call to the YouTube API to fill in the metadata
// about this particular playlist.
func (p *Playlist) FetchMetadata() error {
	return nil
}

// GetID returns the ID of the YouTube playlist.
func (p *Playlist) GetID() string {
	return p.ID
}

// GetAuthor returns the author of the YouTube playlist.
func (p *Playlist) GetAuthor() string {
	return p.Author
}

// GetTitle returns the title of the YouTube playlist.
func (p *Playlist) GetTitle() string {
	return p.Title
}

// GetSubmitter returns the name of the user who submitted the playlist.
func (p *Playlist) GetSubmitter() string {
	return p.Submitter
}

// GetService returns the service name of the YouTube playlist.
func (p *Playlist) GetService() string {
	return "YouTube"
}
