/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/soundcloud/track.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package soundcloud

import "time"

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
	var parsedOffset time.Duration

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
