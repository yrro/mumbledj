/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/soundcloud/playlist.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package soundcloud

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/layeh/gumble/gumble"
	"github.com/spf13/viper"
)

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
	var response *http.Response
	var value *jason.Object
	var err error
	if response, err = http.Get(fmt.Sprintf("http://api.soundcloud.com/playlists/%s?client_id=%s",
		p.ID, viper.GetString("api.soundcloudkey"))); err != nil {
		return err
	}
	if value, err = jason.NewObjectFromReader(response.Body); err != nil {
		return err
	}

	p.Title, _ = value.GetString("title")
	p.Author, _ = value.GetString("user", "username")

	// Fetch track metadata.
	maxTracks := viper.GetInt("general.maxtracksperplaylist")
	playlistInfo := Playlist{
		ID:        p.ID,
		Submitter: p.Submitter,
	}

	tracks, _ := value.GetObjectArray("tracks")

	for _, track := range tracks {
		if maxTracks != 0 && len(p.Tracks) >= maxTracks {
			break
		}

		author, _ := track.GetString("user", "username")
		title, _ := track.GetString("title")
		ID, _ := track.GetString("id")
		durationMilliseconds, _ := track.GetInt64("duration")
		duration, _ := time.ParseDuration(fmt.Sprintf("%dms", durationMilliseconds))
		thumbnail, err := track.GetString("artwork_url")
		if err != nil {
			// Track has no artwork, using profile avatar instead.
			thumbnail, _ = track.GetString("user", "avatar_url")
		}

		newTrack := &Track{
			Submitter:        p.Submitter,
			Author:           author,
			Title:            title,
			ID:               ID,
			Duration:         duration,
			Thumbnail:        thumbnail,
			PossiblePlaylist: playlistInfo,
		}
		p.Tracks = append(p.Tracks, newTrack)
	}

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
