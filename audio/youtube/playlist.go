/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/youtube/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package youtube

import (
	"fmt"
	"net/http"

	"github.com/antonholmquist/jason"
	"github.com/layeh/gumble/gumble"
	"github.com/spf13/viper"
)

// Playlist is a struct that contains metadata related to a YouTube playlist.
type Playlist struct {
	ID        string
	Author    string
	Title     string
	Submitter string
}

// NewPlaylist returns a YouTube playlist struct with the provided ID.
func NewPlaylist(id string, submitter *gumble.User) *Playlist {
	return &Playlist{
		ID:        id,
		Submitter: submitter.Name,
	}
}

// FetchMetadata makes an API call to the YouTube API to fill in the metadata
// about this particular playlist.
func (p *Playlist) FetchMetadata() error {
	var response *http.Response
	var value *jason.Object
	var err error
	if response, err = http.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlists?part=snippet&id=%s&key=%s",
		p.ID, viper.GetString("api.youtubekey"))); err != nil {
		return err
	}
	if value, err = jason.NewObjectFromReader(response.Body); err != nil {
		return err
	}

	p.Title, _ = value.GetString("items", "0", "snippet", "title")
	p.Author, _ = value.GetString("items", "0", "snippet", "channelTitle")

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
