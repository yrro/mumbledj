/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type MumbleDJTestSuite struct {
	Bot MumbleDJ
	suite.Suite
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
