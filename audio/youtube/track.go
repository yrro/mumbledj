/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/youtube/track.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package youtube

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/spf13/viper"
)

// Track is a struct that represents the metadata of a YouTube track.
type Track struct {
	Submitter        string
	Uploader         string
	Title            string
	ID               string
	Filename         string
	Duration         time.Duration
	Offset           time.Duration
	Thumbnail        string
	PossiblePlaylist Playlist
}

// NewTrack initializes and returns a YouTube Track struct that contains the
// name of the submitter, ID of the Track, and the playlist it is associated with
// if any.
func NewTrack(submitter, id, offset string, playlist Playlist) (*Track, error) {
	var parsedOffset time.Duration
	var err error
	if parsedOffset, err = parseDurationString(offset, `\?T\=(?<days>\d+D)?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`); err != nil {
		return nil, err
	}
	return &Track{
		Submitter:        submitter,
		ID:               id,
		Offset:           parsedOffset,
		PossiblePlaylist: playlist,
	}, nil
}

// FetchMetadata makes an API call to the YouTube API to fill in the metadata
// about this particular track.
func (t *Track) FetchMetadata() error {
	var response *http.Response
	var value *jason.Object
	var err error
	if response, err = http.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails&id=%s&key=%s",
		t.ID, viper.GetString("api.youtubekey"))); err != nil {
		return err
	}
	if value, err = jason.NewObjectFromReader(response.Body); err != nil {
		return err
	}

	t.Title, _ = value.GetString("items", "0", "snippet", "title")
	t.Thumbnail, _ = value.GetString("items", "0", "snippet", "thumbnails", "high", "url")
	durationString, _ := value.GetString("items", "0", "contentDetails", "duration")
	t.Duration, _ = parseDurationString(durationString, `P(?P<days>\d+D)?T(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	t.Uploader, _ = value.GetString("items", "0", "snippet", "channelTitle")

	return nil
}

// Download downloads a copy of the YouTube track via youtube-dl to the
// configured download location.
func (t *Track) Download() error {
	return nil
}

// Delete removes the local copy of the YouTube track from the configured
// download location.
func (t *Track) Delete() error {
	return nil
}

// GetSubmitter returns the name of the submitter of the YouTube track.
func (t *Track) GetSubmitter() string {
	return t.Submitter
}

// GetUploader returns the name of the uploader of the YouTube track.
func (t *Track) GetUploader() string {
	return t.Uploader
}

// GetTitle returns the title of the YouTube track.
func (t *Track) GetTitle() string {
	return t.Title
}

// GetID returns the ID of the YouTube track.
func (t *Track) GetID() string {
	return t.ID
}

// GetFilename returns the filename of the YouTube track when stored locally.
func (t *Track) GetFilename() string {
	return t.Filename
}

// GetDuration returns the duration of the YouTube track.
func (t *Track) GetDuration() time.Duration {
	return t.Duration
}

// GetThumbnail returns the URL to the thumbnail for the YouTube track.
func (t *Track) GetThumbnail() string {
	return t.Thumbnail
}

// GetService returns the name of the service.
func (t *Track) GetService() string {
	return "YouTube"
}

// GetPlaylist returns the Playlist struct this track is associated with if it
// is part of a YouTube playlist. Playlist is nil if the track is not associated
// with a playlist.
func (t *Track) GetPlaylist() Playlist {
	return t.PossiblePlaylist
}

func parseDurationString(duration, regex string) (time.Duration, error) {
	var days, hours, minutes, seconds, totalSeconds int64
	var err error
	if duration != "" {
		timestampExp := regexp.MustCompile(regex)
		timestampMatch := timestampExp.FindStringSubmatch(strings.ToUpper(duration))
		timestampResult := make(map[string]string)
		for i, name := range timestampExp.SubexpNames() {
			if i < len(timestampMatch) {
				timestampResult[name] = timestampMatch[i]
			}
		}

		if timestampResult["days"] != "" {
			if days, err = strconv.ParseInt(strings.TrimSuffix(timestampResult["days"], "D"), 10, 32); err != nil {
				return 0, errors.New("Malformed duration string.")
			}
		}
		if timestampResult["hours"] != "" {
			if hours, err = strconv.ParseInt(strings.TrimSuffix(timestampResult["hours"], "H"), 10, 32); err != nil {
				return 0, errors.New("Malformed duration string.")
			}
		}
		if timestampResult["minutes"] != "" {
			if minutes, err = strconv.ParseInt(strings.TrimSuffix(timestampResult["minutes"], "M"), 10, 32); err != nil {
				return 0, errors.New("Malformed duration string.")
			}
		}
		if timestampResult["seconds"] != "" {
			if seconds, err = strconv.ParseInt(strings.TrimSuffix(timestampResult["seconds"], "S"), 10, 32); err != nil {
				return 0, errors.New("Malformed duration string.")
			}
		}

		totalSeconds = int64((days * 86400) + (hours * 3600) + (minutes * 60) + seconds)
	} else {
		totalSeconds = 0
	}

	output, _ := time.ParseDuration(strconv.Itoa(int(totalSeconds)) + "s")
	return output, nil
}
