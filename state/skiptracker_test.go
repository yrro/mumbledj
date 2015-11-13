/*
 * MumbleDJ
 * By Matthieu Grieger
 * state/skiptracker_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import (
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SkipTrackerTestSuite struct {
	suite.Suite
	Tracker *SkipTracker
}

type MockedGumbleUser struct {
	gumble.User
	mock.Mock
}

func (suite *SkipTrackerTestSuite) SetupTest() {
	suite.Tracker = NewSkipTracker()
}

func (suite *SkipTrackerTestSuite) TestNewSkipTracker() {
	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be empty upon initialization.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be empty upon initialization.")
}

func (suite *SkipTrackerTestSuite) TestAddTrackSkip() {

}

func (suite *SkipTrackerTestSuite) TestAddPlaylistSkip() {

}

func (suite *SkipTrackerTestSuite) TestRemoveTrackSkip() {

}

func (suite *SkipTrackerTestSuite) TestRemovePlaylistSkip() {

}

func (suite *SkipTrackerTestSuite) TestResetTrackSkips() {

}

func (suite *SkipTrackerTestSuite) TestResetPlaylistSkips() {

}

func TestSkipTrackerTestSuite(t *testing.T) {
	suite.Run(t, new(SkipTrackerTestSuite))
}
