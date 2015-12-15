/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/mumbledj_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import (
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleffmpeg"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MumbleDJTestSuite struct {
	Bot MumbleDJ
	suite.Suite
}

type MockedGumbleClient struct {
	gumble.Client
	mock.Mock
}

type MockedGumbleConfig struct {
	gumble.Config
	mock.Mock
}

type MockedAudioStream struct {
	gumbleffmpeg.Stream
	mock.Mock
}

type MockedQueue struct {
	state.AudioQueue
	mock.Mock
}

type MockedCache struct {
	state.AudioCache
	mock.Mock
}

type MockedCommander struct {
	Commander
	mock.Mock
}

func (suite *MumbleDJTestSuite) SetupTest() {
	viper.Set("general.defaultchannel", []string{""})
	viper.Set("volume.default", 0.4)
	viper.Set("general.defaultcomment", "I'm a bot.")
	viper.Set("cache.enabled", false)
	viper.Set("connection.retryenabled", false)
	viper.Set("connection.retryinterval", 1)
	viper.Set("connection.retryattempts", 1)
	viper.Set("general.commandprefix", "!")
	viper.Set("permissions.adminsenabled", false)
	viper.Set("permissions.admins", []string{"Matt"})
}

func (suite *MumbleDJTestSuite) TestOnConnect() {

}

func (suite *MumbleDJTestSuite) TestOnConnectCacheEnabled() {

}

func (suite *MumbleDJTestSuite) TestOnDisconnect() {

}

func (suite *MumbleDJTestSuite) TestOnDisconnectRetryEnabled() {

}

func (suite *MumbleDJTestSuite) TestOnTextMessage() {

}

func (suite *MumbleDJTestSuite) TestOnUserChange() {

}

func (suite *MumbleDJTestSuite) TestAddSkip() {

}

func (suite *MumbleDJTestSuite) TestRemoveSkip() {

}

func (suite *MumbleDJTestSuite) TestResetSkips() {

}

func (suite *MumbleDJTestSuite) TestHasPermission() {

}

func (suite *MumbleDJTestSuite) TestSendPrivateMessage() {

}

func (suite *MumbleDJTestSuite) TestStart() {

}

func TestMumbleDJTestSuite(t *testing.T) {
	suite.Run(t, new(MumbleDJTestSuite))
}
