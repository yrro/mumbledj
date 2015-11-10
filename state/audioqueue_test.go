/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj/audioqueue_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/audio"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AudioQueueTestSuite struct {
	suite.Suite
	Queue       *AudioQueue
	FirstTrack  MockedAudioTrack
	SecondTrack MockedAudioTrack
	ThirdTrack  MockedAudioTrack
}

type MockedAudioTrack struct {
	mock.Mock
	audio.Track
	Identifier string
}

func (suite *AudioQueueTestSuite) SetupSuite() {
	viper.Set("general.automaticshuffleon", false)
	suite.FirstTrack = MockedAudioTrack{Identifier: "first"}
	suite.SecondTrack = MockedAudioTrack{Identifier: "second"}
	suite.ThirdTrack = MockedAudioTrack{Identifier: "third"}
}

func (suite *AudioQueueTestSuite) SetupTest() {
	suite.Queue = NewAudioQueue()
}

func (suite *AudioQueueTestSuite) TestNewAudioQueue() {
	suite.Zero(len(suite.Queue.Queue), "The queue should be empty.")
}

func (suite *AudioQueueTestSuite) TestAddNewTrack() {
	suite.Zero(len(suite.Queue.Queue), "The queue should be empty.")
	suite.Queue.AddTrack(MockedAudioTrack{})
	suite.Equal(1, len(suite.Queue.Queue), "There should now be one track in the queue.")
}

func (suite *AudioQueueTestSuite) TestCurrentTrack() {
	result, err := suite.Queue.CurrentTrack()
	suite.Nil(result, "Result should be nil as there are no tracks in the queue.")
	suite.NotNil(err, "An error should be returned as there are no tracks in the queue.")

	suite.Queue.AddTrack(suite.FirstTrack)

	result, err = suite.Queue.CurrentTrack()
	suite.NotNil(result, "Result should not be nil as there is now a track in the queue.")
	suite.Nil(err, "An error shouldn't have occurred.")
}

func (suite *AudioQueueTestSuite) TestPeekNextTrack() {
	result, err := suite.Queue.PeekNextTrack()
	suite.Nil(result, "Result should be nil as there are no tracks in the queue.")
	suite.NotNil(err, "An error should be returned as there are no tracks in the queue.")

	suite.Queue.AddTrack(suite.FirstTrack)

	result, err = suite.Queue.PeekNextTrack()
	suite.Nil(result, "Result should be nil as there is only one track in the queue.")
	suite.NotNil(err, "An error should be returned as there is only one track in the queue.")

	suite.Queue.AddTrack(suite.SecondTrack)

	result, err = suite.Queue.PeekNextTrack()
	suite.NotNil(result, "Result should not be nil as there are now two tracks in the queue.")
	suite.Equal(suite.SecondTrack, result, "The result should be equal to the second mocked track.")
	suite.Nil(err, "An error shouldn't have occurred.")
}

func (suite *AudioQueueTestSuite) TestTraverse() {
	suite.Queue.AddTrack(suite.FirstTrack)
	suite.Queue.Traverse(func(i int, t audio.Track) {})
}

func (suite *AudioQueueTestSuite) TestShuffleTracks() {
	suite.Queue.AddTrack(suite.FirstTrack)

	suite.Queue.ShuffleTracks()
	suite.Equal(suite.FirstTrack, suite.Queue.Queue[0], "Shuffle shouldn't do anything when only one track is in the queue.")

	suite.Queue.AddTrack(suite.SecondTrack)

	suite.Queue.ShuffleTracks()
	suite.Equal(suite.FirstTrack, suite.Queue.Queue[0], "Shuffle shouldn't do anything when only two tracks are in the queue.")
	suite.Equal(suite.SecondTrack, suite.Queue.Queue[1], "Shuffle shouldn't do anything when only two tracks are in the queue.")

	suite.Queue.AddTrack(suite.ThirdTrack)
	suite.Queue.AddTrack(MockedAudioTrack{Identifier: "fourth"})
	suite.Queue.AddTrack(MockedAudioTrack{Identifier: "fifth"})

	queueBefore := make([]audio.Track, len(suite.Queue.Queue))
	copy(queueBefore, suite.Queue.Queue)
	suite.Queue.ShuffleTracks()
	suite.NotEqual(queueBefore, suite.Queue.Queue, "The shuffled queue should not be the same as the original queue.")
}

func (suite *AudioQueueTestSuite) TestRandomNextTrack() {
	suite.Queue.AddTrack(suite.FirstTrack)

	suite.Queue.RandomNextTrack(true)
	suite.Equal(suite.FirstTrack, suite.Queue.Queue[0], "RandomNextTrack shouldn't do anything when there is only one track in the queue.")

	suite.Queue.AddTrack(suite.SecondTrack)

	suite.Queue.RandomNextTrack(false)
	suite.Equal(suite.FirstTrack, suite.Queue.Queue[0], "RandomNextTrack shouldn't do anything when there are only two tracks in the queue.")
}

func TestAudioQueueTestSuite(t *testing.T) {
	suite.Run(t, new(AudioQueueTestSuite))
}
