/*
 * MumbleDJ
 * By Matthieu Grieger
 * config_test.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestDefaultValues(t *testing.T) {
	SetDefaults()

	assert.Equal(t, viper.GetString("Server"), "127.0.0.1")
	assert.Equal(t, viper.GetString("Port"), "64738")
	assert.Equal(t, viper.GetString("Username"), "MumbleDJ")
	assert.Equal(t, viper.GetString("Password"), "")
	assert.Equal(t, viper.GetString("Channel"), "root")
	assert.Equal(t, viper.GetString("Cert"), "")
	assert.Equal(t, viper.GetString("Key"), "")
	assert.False(t, viper.GetBool("Insecure"))
	assert.Equal(t, viper.GetString("AccessTokens"), "")
	assert.True(t, viper.GetBool("RetryConnectionEnabled"))
	assert.Equal(t, viper.GetInt("RetryInterval"), 30)
	assert.Equal(t, viper.GetInt("RetryAttempts"), 10)

	assert.Equal(t, viper.GetString("CommandPrefix"), "!")
	assert.Equal(t, viper.GetFloat64("SkipRatio"), 0.5)
	assert.Equal(t, viper.GetFloat64("PlaylistSkipRatio"), 0.5)
	assert.Equal(t, viper.GetString("DefaultComment"), "Hello! I am a bot. Type !help for a list of commands.")
	assert.Equal(t, viper.GetInt("MaxTrackDuration"), 0)

	assert.False(t, viper.GetBool("CacheEnabled"))
	assert.Equal(t, viper.GetInt("CacheMaximumSize"), 512)
	assert.Equal(t, viper.GetInt("CacheExpiry"), 24)

	assert.Equal(t, viper.GetFloat64("DefaultVolume"), 0.2)
	assert.Equal(t, viper.GetFloat64("LowestVolume"), 0.01)
	assert.Equal(t, viper.GetFloat64("HighestVolume"), 0.8)

	assert.Equal(t, viper.GetStringSlice("AddAliases"), []string{"add", "a"})
	assert.Equal(t, viper.GetStringSlice("SkipAliases"), []string{"skip", "s"})
	assert.Equal(t, viper.GetStringSlice("SkipPlaylistAliases"), []string{"skipplaylist", "sp"})
	assert.Equal(t, viper.GetStringSlice("AdminSkipAliases"), []string{"forceskip", "fs"})
	assert.Equal(t, viper.GetStringSlice("AdminSkipPlaylistAliases"), []string{"forceskipplaylist", "fsp"})
	assert.Equal(t, viper.GetStringSlice("HelpAliases"), []string{"help", "h"})
	assert.Equal(t, viper.GetStringSlice("VolumeAliases"), []string{"volume", "v"})
	assert.Equal(t, viper.GetStringSlice("MoveAliases"), []string{"move", "m"})
	assert.Equal(t, viper.GetStringSlice("ReloadAliases"), []string{"reload", "r"})
	assert.Equal(t, viper.GetStringSlice("QueueResetAliases"), []string{"resetqueue", "reset", "rq"})
	assert.Equal(t, viper.GetStringSlice("NumTracksAliases"), []string{"numtracks", "numsongs", "nt", "ns"})
	assert.Equal(t, viper.GetStringSlice("NextTrackAliases"), []string{"nexttrack", "nextsong", "next"})
	assert.Equal(t, viper.GetStringSlice("CurrentTrackAliases"), []string{"currenttrack", "currentsong", "current"})
	assert.Equal(t, viper.GetStringSlice("SetCommentAliases"), []string{"setcomment", "comment"})
	assert.Equal(t, viper.GetStringSlice("NumCachedAliases"), []string{"numcached", "nc"})
	assert.Equal(t, viper.GetStringSlice("CacheSizeAliases"), []string{"cachesize", "cached"})
	assert.Equal(t, viper.GetStringSlice("KillAliases"), []string{"kill", "k"})

	assert.True(t, viper.GetBool("AdminsEnabled"))
	assert.Equal(t, viper.GetStringSlice("Admins"), []string{"Matt"})
	assert.False(t, viper.GetBool("AdminAdd"))
	assert.False(t, viper.GetBool("AdminAddPlaylist"))
	assert.False(t, viper.GetBool("AdminSkip"))
	assert.False(t, viper.GetBool("AdminHelp"))
	assert.False(t, viper.GetBool("AdminVolume"))
	assert.True(t, viper.GetBool("AdminMove"))
	assert.True(t, viper.GetBool("AdminReload"))
	assert.True(t, viper.GetBool("AdminReset"))
	assert.False(t, viper.GetBool("AdminNumTracks"))
	assert.False(t, viper.GetBool("AdminNextTrack"))
	assert.False(t, viper.GetBool("AdminCurrentTrack"))
	assert.True(t, viper.GetBool("AdminSetComment"))
	assert.True(t, viper.GetBool("AdminNumCached"))
	assert.True(t, viper.GetBool("AdminCacheSize"))
	assert.True(t, viper.GetBool("AdminKill"))
}
