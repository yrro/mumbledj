/*
 * MumbleDJ
 * By Matthieu Grieger
 * testutils/mocks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package testutils

import (
	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/stretchr/testify/mock"
)

// MockedAudioTrack is a mocked audio track for testing purposes.
type MockedAudioTrack struct {
	audio.Track
	mock.Mock
}
