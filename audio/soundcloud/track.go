/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/soundcloud/track.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package soundcloud

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/spf13/viper"
)

// Track is a struct that represents the metadata of a SoundCloud track.
type Track struct {
	Submitter        string
	Author           string
	Title            string
	ID               string
	Filename         string
	Duration         time.Duration
	Offset           time.Duration
	Thumbnail        string
	PossiblePlaylist Playlist
}

// NewTrack initializes and returns a SoundCloud Track struct that contains the
// name of the submitter, ID of the Track, and the playlist it is associated with
// if any.
func NewTrack(submitter, id, offset string, playlist Playlist) (*Track, error) {
	timesplit := strings.Split(offset, ":")
	offsetSeconds := 0
	multiplier := 1
	for i := len(timesplit) - 1; i >= 0; i-- {
		time, _ := strconv.Atoi(timesplit[i])
		offsetSeconds += time * multiplier
		multiplier *= 60
	}

	parsedOffset, _ := time.ParseDuration(fmt.Sprintf("%ds", offsetSeconds))

	return &Track{
		Submitter:        submitter,
		ID:               id,
		Offset:           parsedOffset,
		PossiblePlaylist: playlist,
	}, nil
}

// FetchMetadata makes an API call to the SoundCloud API to fill in the metadata
// about this particular track.
func (t *Track) FetchMetadata() error {
	var response *http.Response
	var value *jason.Object
	var err error
	if response, err = http.Get(fmt.Sprintf("http://api.soundcloud.com/resolve?url=%s&client_id=%s",
		t.ID, viper.GetString("api.soundcloudkey"))); err != nil {
		return err
	}
	if value, err = jason.NewObjectFromReader(response.Body); err != nil {
		return err
	}

	t.Title, _ = value.GetString("title")
	durationMilliseconds, _ := value.GetInt64("duration")
	t.Duration, _ = time.ParseDuration(fmt.Sprintf("%dms", durationMilliseconds))
	t.Author, _ = value.GetString("user", "username")
	t.Thumbnail, err = value.GetString("artwork_url")
	if err != nil {
		// Track has no artwork, using profile avatar instead.
		t.Thumbnail, _ = value.GetString("user", "avatar_url")
	}

	return nil
}

// Download downloads a copy of the SoundCloud track via youtube-dl to the
// configured download location.
func (t *Track) Download() error {
	return nil
}

// Delete removes the local copy of the SoundCloud track from the configured
// download location.
func (t *Track) Delete() error {
	return nil
}

// GetSubmitter returns the name of the submitter of the SoundCloud track.
func (t *Track) GetSubmitter() string {
	return t.Submitter
}

// GetAuthor returns the name of the author of the SoundCloud track.
func (t *Track) GetAuthor() string {
	return t.Author
}

// GetTitle returns the title of the SoundCloud track.
func (t *Track) GetTitle() string {
	return t.Title
}

// GetID returns the ID of the SoundCloud track.
func (t *Track) GetID() string {
	return t.ID
}

// GetFilename returns the filename of the SoundCloud track when stored locally.
func (t *Track) GetFilename() string {
	return t.Filename
}

// GetDuration returns the duration of the SoundCloud track.
func (t *Track) GetDuration() time.Duration {
	return t.Duration
}

// GetThumbnail returns the URL to the thumbnail for the SoundCloud track.
func (t *Track) GetThumbnail() string {
	return t.Thumbnail
}

// GetService returns the name of the service.
func (t *Track) GetService() string {
	return "SoundCloud"
}

// GetPlaylist returns the Playlist struct this track is associated with if it
// is part of a SoundCloud playlist. Playlist is nil if the track is not associated
// with a playlist.
func (t *Track) GetPlaylist() Playlist {
	return t.PossiblePlaylist
}
