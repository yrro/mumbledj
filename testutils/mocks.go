/*
 * MumbleDJ
 * By Matthieu Grieger
 * testutils/mocks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package testutils

import (
	"errors"
	"strconv"

	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/stretchr/testify/mock"
)

// MockedAudioTrack is a mocked audio track for testing purposes.
type MockedAudioTrack struct {
	audio.Track
	mock.Mock
}

// Title is a method that returns the title of the track.
func (m *MockedAudioTrack) Title() string {
	return "TestTrack"
}

// Submitter is a method that returns the submitter of the track.
func (m *MockedAudioTrack) Submitter() string {
	return "TestSubmitter"
}

// Service is a method that returns the service the track comes from.
func (m *MockedAudioTrack) Service() string {
	return "TestService"
}

// MockedAudioHandler is a mocked audio handler for testing purposes.
type MockedAudioHandler struct {
	audio.Handler
	mock.Mock
}

// GetTracks returns a number of MockedAudioTracks specified with the num argument.
// If num is 0 a test error is returned.
func (m *MockedAudioHandler) GetTracks(num string) ([]audio.Track, error) {
	var tracks []audio.Track
	numTracks, _ := strconv.ParseInt(num, 10, 32)
	if numTracks == 0 {
		return nil, errors.New("Test error.")
	}

	i := 0
	for i < int(numTracks) {
		tracks = append(tracks, new(MockedAudioTrack))
		i++
	}

	return tracks, nil
}
