/*
 * MumbleDJ
 * By Matthieu Grieger
 * state/config_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	State *BotState
	suite.Suite
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.State = new(BotState)
	suite.State.BotConfig = new(Config)
	suite.State.BotConfig.SetDefaultConfiguration()
}

func (suite *ConfigTestSuite) TestDefaultConfiguration() {
	suite.Equal("!", viper.GetString("general.commandprefix"))
	suite.Equal(0.5, viper.GetFloat64("general.skipratio"))
	suite.Equal(0.5, viper.GetFloat64("general.playlistskipratio"))
	suite.Equal("Hello! I am a bot. Type !help for a list of commands.", viper.GetString("general.defaultcomment"))
	suite.Equal([]string{""}, viper.GetStringSlice("general.defaultchannel"))
	suite.Zero(viper.GetInt("general.maxtrackduration"))
	suite.Equal(50, viper.GetInt("general.maxtracksperplaylist"))
	suite.False(viper.GetBool("general.automaticshuffleon"))

	suite.Equal("127.0.0.1", viper.GetString("connection.address"))
	suite.Equal(64738, viper.GetInt("connection.port"))
	suite.Equal("", viper.GetString("connection.password"))
	suite.Equal("MumbleDJ", viper.GetString("connection.username"))
	suite.False(viper.GetBool("connection.insecure"))
	suite.Equal("", viper.GetString("connection.cert"))
	suite.Equal("", viper.GetString("connection.key"))
	suite.Equal([]string{""}, viper.GetStringSlice("connection.accesstokens"))
	suite.True(viper.GetBool("connection.retryenabled"))
	suite.Equal(10, viper.GetInt("connection.retryattempts"))
	suite.Equal(5, viper.GetInt("connection.retryinterval"))

	suite.Equal("", viper.GetString("api.youtubekey"))
	suite.Equal("", viper.GetString("api.soundcloudkey"))

	suite.Equal(0.4, viper.GetFloat64("volume.default"))
	suite.Equal(0.01, viper.GetFloat64("volume.lowest"))
	suite.Equal(0.8, viper.GetFloat64("volume.highest"))

	suite.False(viper.GetBool("cache.enabled"))
	suite.Equal(512, viper.GetInt("cache.maximumsize"))
	suite.Equal(24, viper.GetInt("cache.expiretime"))
	suite.Equal(5, viper.GetInt("cache.checkinterval"))
	suite.Equal("~/.mumbledj/cache", viper.GetString("cache.directory"))

	suite.Equal([]string{"add", "a"}, viper.GetStringSlice("aliases.add"))
	suite.Equal([]string{"skip", "s"}, viper.GetStringSlice("aliases.skip"))
	suite.Equal([]string{"skipplaylist", "sp"}, viper.GetStringSlice("aliases.skipplaylist"))
	suite.Equal([]string{"forceskip", "fs"}, viper.GetStringSlice("aliases.forceskip"))
	suite.Equal([]string{"forceskipplaylist", "fsp"}, viper.GetStringSlice("aliases.forceskipplaylist"))
	suite.Equal([]string{"help", "h"}, viper.GetStringSlice("aliases.help"))
	suite.Equal([]string{"volume", "v"}, viper.GetStringSlice("aliases.volume"))
	suite.Equal([]string{"move", "m"}, viper.GetStringSlice("aliases.move"))
	suite.Equal([]string{"reload", "r"}, viper.GetStringSlice("aliases.reload"))
	suite.Equal([]string{"reset", "re"}, viper.GetStringSlice("aliases.reset"))
	suite.Equal([]string{"numtracks", "nt"}, viper.GetStringSlice("aliases.numtracks"))
	suite.Equal([]string{"nexttrack", "next"}, viper.GetStringSlice("aliases.nexttrack"))
	suite.Equal([]string{"currenttrack", "current"}, viper.GetStringSlice("aliases.currenttrack"))
	suite.Equal([]string{"setcomment", "sc"}, viper.GetStringSlice("aliases.setcomment"))
	suite.Equal([]string{"numcached", "nc"}, viper.GetStringSlice("aliases.numcached"))
	suite.Equal([]string{"cachesize", "cs"}, viper.GetStringSlice("aliases.cachesize"))
	suite.Equal([]string{"kill", "k"}, viper.GetStringSlice("aliases.kill"))
	suite.Equal([]string{"shuffle", "shuf", "sh"}, viper.GetStringSlice("aliases.shuffle"))
	suite.Equal([]string{"shuffleon", "shufon", "shon"}, viper.GetStringSlice("aliases.shuffleon"))
	suite.Equal([]string{"shuffleoff", "shufoff", "shoff"}, viper.GetStringSlice("aliases.shuffleoff"))
	suite.Equal([]string{"listtracks", "listsongs", "list", "l"}, viper.GetStringSlice("aliases.listtracks"))

	suite.True(viper.GetBool("permissions.adminsenabled"))
	suite.Equal([]string{"Matt"}, viper.GetStringSlice("permissions.admins"))
	suite.False(viper.GetBool("permissions.add"))
	suite.False(viper.GetBool("permissions.skip"))
	suite.False(viper.GetBool("permissions.help"))
	suite.False(viper.GetBool("permissions.volume"))
	suite.True(viper.GetBool("permissions.move"))
	suite.True(viper.GetBool("permissions.reload"))
	suite.True(viper.GetBool("permissions.reset"))
	suite.False(viper.GetBool("permissions.numtracks"))
	suite.False(viper.GetBool("permissions.nexttrack"))
	suite.False(viper.GetBool("permissions.currenttrack"))
	suite.True(viper.GetBool("permissions.setcomment"))
	suite.True(viper.GetBool("permissions.numcached"))
	suite.True(viper.GetBool("permissions.cachesize"))
	suite.True(viper.GetBool("permissions.kill"))
	suite.True(viper.GetBool("permissions.shuffle"))
	suite.True(viper.GetBool("permissions.shuffletoggle"))
	suite.False(viper.GetBool("permissions.listtracks"))

	suite.Equal("Adds a track or playlist from YouTube or SoundCloud to the audio queue.", viper.GetString("descriptions.add"))
	suite.Equal("Places a vote to skip the current track.", viper.GetString("descriptions.skip"))
	suite.Equal("Places a vote to skip the current playlist.", viper.GetString("descriptions.skipplaylist"))
	suite.Equal("Immediately skips the current track.", viper.GetString("descriptions.forceskip"))
	suite.Equal("Immediately skips the current playlist.", viper.GetString("descriptions.forceskipplaylist"))
	suite.Equal("Outputs this list of commands.", viper.GetString("descriptions.help"))
	suite.Equal("Changes the volume if an argument is provided, outputs the current volume otherwise.", viper.GetString("descriptions.volume"))
	suite.Equal("Moves the bot into the Mumble channel provided via argument.", viper.GetString("descriptions.move"))
	suite.Equal("Reloads the configuration file.", viper.GetString("descriptions.reload"))
	suite.Equal("Resets the audio queue by removing all queue items.", viper.GetString("descriptions.reset"))
	suite.Equal("Outputs the number of tracks currently in the audio queue.", viper.GetString("descriptions.numtracks"))
	suite.Equal("Outputs the title and submitter of the next track in the queue if one exists.", viper.GetString("descriptions.nexttrack"))
	suite.Equal("Outputs the title and submitter of the current track if one exists.", viper.GetString("descriptions.currenttrack"))
	suite.Equal("Sets the comment displayed next to MumbleDJ's username in Mumble.", viper.GetString("descriptions.setcomment"))
	suite.Equal("Outputs the number of tracks cached on disk if caching is enabled.", viper.GetString("descriptions.numcached"))
	suite.Equal("Outputs the file size of the cache in MiB if caching is enabled.", viper.GetString("descriptions.cachesize"))
	suite.Equal("Stops the bot and cleans its cache directory.", viper.GetString("descriptions.kill"))
	suite.Equal("Randomizes the tracks currently in the audio queue.", viper.GetString("descriptions.shuffle"))
	suite.Equal("Toggles permanent track shuffling on.", viper.GetString("descriptions.shuffleon"))
	suite.Equal("Toggles permanent track shuffling off.", viper.GetString("descriptions.shuffleoff"))
	suite.Equal("Outputs a list of the tracks currently in the queue.", viper.GetString("descriptions.listtracks"))
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
