/*
 * MumbleDJ
 * By Matthieu Grieger
 * state/skiptracker.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import (
	"errors"

	"github.com/layeh/gumble/gumble"
)

// SkipTracker is a struct that keeps track of the list of users who have
// skipped the current track or playlist.
type SkipTracker struct {
	TrackSkips    []*gumble.User
	PlaylistSkips []*gumble.User
}

// NewSkipTracker returns an empty SkipTracker.
func NewSkipTracker() *SkipTracker {
	return &SkipTracker{
		TrackSkips:    make([]*gumble.User, 0),
		PlaylistSkips: make([]*gumble.User, 0),
	}
}

// AddTrackSkip adds a skip to the SkipTracker for the current track.
func (s *SkipTracker) AddTrackSkip(skipper *gumble.User) error {
	for _, user := range s.TrackSkips {
		if user.Name == skipper.Name {
			return errors.New("The user has already voted to skip the track.")
		}
	}
	s.TrackSkips = append(s.TrackSkips, skipper)
	return nil
}

// AddPlaylistSkip adds a skip to the SkipTracker for the current playlist.
func (s *SkipTracker) AddPlaylistSkip(skipper *gumble.User) error {
	for _, user := range s.PlaylistSkips {
		if user.Name == skipper.Name {
			return errors.New("The user has already voted to skip the playlist.")
		}
	}
	s.PlaylistSkips = append(s.PlaylistSkips, skipper)
	return nil
}

// RemoveTrackSkip removes a skip from the SkipTracker for the current track.
func (s *SkipTracker) RemoveTrackSkip(skipper *gumble.User) error {
	for i, user := range s.TrackSkips {
		if user.Name == skipper.Name {
			s.TrackSkips = append(s.TrackSkips[:i], s.TrackSkips[i+1:]...)
			return nil
		}
	}
	return errors.New("The user did not previously vote to skip the track.")
}

// RemovePlaylistSkip removes a skip from the SkipTracker for the current playlist.
func (s *SkipTracker) RemovePlaylistSkip(skipper *gumble.User) error {
	for i, user := range s.PlaylistSkips {
		if user.Name == skipper.Name {
			s.PlaylistSkips = append(s.PlaylistSkips[:i], s.PlaylistSkips[i+1:]...)
			return nil
		}
	}
	return errors.New("The user did not previously vote to skip the playlist.")
}

// ResetTrackSkips resets the skip slice for the current track.
func (s *SkipTracker) ResetTrackSkips() {
	s.TrackSkips = s.TrackSkips[:0]
}

// ResetPlaylistSkips resets the skip slice for the current playlist.
func (s *SkipTracker) ResetPlaylistSkips() {
	s.PlaylistSkips = s.PlaylistSkips[:0]
}
