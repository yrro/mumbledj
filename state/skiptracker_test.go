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
	"github.com/stretchr/testify/suite"
)

type SkipTrackerTestSuite struct {
	suite.Suite
	Tracker *SkipTracker
	User1   *gumble.User
	User2   *gumble.User
}

func (suite *SkipTrackerTestSuite) SetupTest() {
	suite.Tracker = NewSkipTracker()
	suite.User1 = new(gumble.User)
	suite.User1.Name = "User1"
	suite.User2 = new(gumble.User)
	suite.User2.Name = "User2"
}

func (suite *SkipTrackerTestSuite) TestNewSkipTracker() {
	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be empty upon initialization.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be empty upon initialization.")
}

func (suite *SkipTrackerTestSuite) TestAddTrackSkip() {
	suite.Tracker.AddTrackSkip(suite.User1)

	suite.Equal(1, len(suite.Tracker.TrackSkips), "There should now be one user who has skipped the current track.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be unaffected.")

	suite.Tracker.AddTrackSkip(suite.User2)

	suite.Equal(2, len(suite.Tracker.TrackSkips), "There should now be two users who have skipped the current track.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be unaffected.")

	suite.Tracker.AddTrackSkip(suite.User1)
	suite.Tracker.AddTrackSkip(suite.User2)

	suite.Equal(2, len(suite.Tracker.TrackSkips), "This is a duplicate skip, so the track skip slice should be unaffected.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be unaffected.")
}

func (suite *SkipTrackerTestSuite) TestAddPlaylistSkip() {
	suite.Tracker.AddPlaylistSkip(suite.User1)

	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(1, len(suite.Tracker.PlaylistSkips), "There should now be one user who has skipped the current playlist.")

	suite.Tracker.AddPlaylistSkip(suite.User2)

	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(2, len(suite.Tracker.PlaylistSkips), "There should now be two users who have skipped the current playlist.")

	suite.Tracker.AddPlaylistSkip(suite.User1)
	suite.Tracker.AddPlaylistSkip(suite.User2)

	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(2, len(suite.Tracker.PlaylistSkips), "This is a duplicate skip, so the playlist skip slice should be unaffected.")
}

func (suite *SkipTrackerTestSuite) TestRemoveTrackSkip() {
	suite.Tracker.AddTrackSkip(suite.User1)

	suite.Tracker.RemoveTrackSkip(suite.User2)

	suite.Equal(1, len(suite.Tracker.TrackSkips), "User2 has not skipped the track so the track skip slice should be unaffected.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be unaffected.")

	suite.Tracker.RemoveTrackSkip(suite.User1)

	suite.Zero(len(suite.Tracker.TrackSkips), "User1 skipped the track, so their skip should be removed.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be unaffected.")
}

func (suite *SkipTrackerTestSuite) TestRemovePlaylistSkip() {
	suite.Tracker.AddPlaylistSkip(suite.User1)

	suite.Tracker.RemovePlaylistSkip(suite.User2)

	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(1, len(suite.Tracker.PlaylistSkips), "User2 has not skipped the playlist so the playlist skip slice should be unaffected.")

	suite.Tracker.RemovePlaylistSkip(suite.User1)

	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice should be unaffected.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "User1 skipped the playlist, so their skip should be removed.")
}

func (suite *SkipTrackerTestSuite) TestResetTrackSkips() {
	suite.Tracker.AddTrackSkip(suite.User1)
	suite.Tracker.AddTrackSkip(suite.User2)
	suite.Tracker.AddPlaylistSkip(suite.User1)
	suite.Tracker.AddPlaylistSkip(suite.User2)

	suite.Tracker.ResetTrackSkips()

	suite.Zero(len(suite.Tracker.TrackSkips), "The track skip slice has been reset, so the length should be zero.")
	suite.Equal(2, len(suite.Tracker.PlaylistSkips), "The playlist skip slice should be unaffected.")
}

func (suite *SkipTrackerTestSuite) TestResetPlaylistSkips() {
	suite.Tracker.AddTrackSkip(suite.User1)
	suite.Tracker.AddTrackSkip(suite.User2)
	suite.Tracker.AddPlaylistSkip(suite.User1)
	suite.Tracker.AddPlaylistSkip(suite.User2)

	suite.Tracker.ResetPlaylistSkips()

	suite.Equal(2, len(suite.Tracker.TrackSkips), "The track skip slice should be unaffected.")
	suite.Zero(len(suite.Tracker.PlaylistSkips), "The playlist skip slice has been reset, so the length should be zero.")
}

func TestSkipTrackerTestSuite(t *testing.T) {
	suite.Run(t, new(SkipTrackerTestSuite))
}
