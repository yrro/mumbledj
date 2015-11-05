/*
 * MumbleDJ
 * By Matthieu Grieger
 * dj/config.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package dj

import "github.com/spf13/viper"

// SetDefaultConfiguration initalizes a Viper configuration with the default
// values.
func SetDefaultConfiguration() {
	// General configuration
	viper.SetDefault("general.commandprefix", "!")
	viper.SetDefault("general.skipratio", 0.5)
	viper.SetDefault("general.playlistskipratio", 0.5)
	viper.SetDefault("general.defaultcomment", "Hello! I am a bot. Type !help for a list of commands.")
	viper.SetDefault("general.defaultchannel", []string{""})
	viper.SetDefault("general.maxtrackduration", 0)
	viper.SetDefault("general.maxtracksperplaylist", 50)
	viper.SetDefault("general.automaticshuffleon", false)

	// Connection configuration
	viper.SetDefault("connection.address", "127.0.0.1")
	viper.SetDefault("connection.port", 64738)
	viper.SetDefault("connection.password", "")
	viper.SetDefault("connection.username", "MumbleDJ")
	viper.SetDefault("connection.insecure", false)
	viper.SetDefault("connection.cert", "")
	viper.SetDefault("connection.key", "")
	viper.SetDefault("connection.accesstokens", []string{""})
	viper.SetDefault("connection.retryenabled", true)
	viper.SetDefault("connection.retryattempts", 10)
	viper.SetDefault("connection.retryinterval", 5)

	// API key configuration
	viper.SetDefault("api.youtubekey", "")
	viper.SetDefault("api.soundcloudkey", "")

	// Volume configuration
	viper.SetDefault("volume.default", 0.4)
	viper.SetDefault("volume.lowest", 0.01)
	viper.SetDefault("volume.highest", 0.8)

	// Cache configuration
	viper.SetDefault("cache.enabled", false)
	viper.SetDefault("cache.maximumsize", 512)
	viper.SetDefault("cache.expiretime", 24)
	viper.SetDefault("cache.checkinterval", 5)
	viper.SetDefault("cache.directory", "~/.mumbledj/cache")

	// Command alias configuration
	viper.SetDefault("aliases.add", []string{"add", "a"})
	viper.SetDefault("aliases.skip", []string{"skip", "s"})
	viper.SetDefault("aliases.skipplaylist", []string{"skipplaylist", "sp"})
	viper.SetDefault("aliases.adminskip", []string{"forceskip", "fs"})
	viper.SetDefault("aliases.adminskipplaylist", []string{"forceskipplaylist", "fsp"})
	viper.SetDefault("aliases.help", []string{"help", "h"})
	viper.SetDefault("aliases.volume", []string{"volume", "v"})
	viper.SetDefault("aliases.move", []string{"move", "m"})
	viper.SetDefault("aliases.reload", []string{"reload", "r"})
	viper.SetDefault("aliases.reset", []string{"reset", "re"})
	viper.SetDefault("aliases.numtracks", []string{"numtracks", "nt"})
	viper.SetDefault("aliases.nexttrack", []string{"nexttrack", "next"})
	viper.SetDefault("aliases.currenttrack", []string{"currenttrack", "current"})
	viper.SetDefault("aliases.setcomment", []string{"setcomment", "sc"})
	viper.SetDefault("aliases.numcached", []string{"numcached", "nc"})
	viper.SetDefault("aliases.cachesize", []string{"cachesize", "cs"})
	viper.SetDefault("aliases.kill", []string{"kill", "k"})
	viper.SetDefault("aliases.shuffle", []string{"shuffle", "shuf", "sh"})
	viper.SetDefault("aliases.shuffleon", []string{"shuffleon", "shufon", "shon"})
	viper.SetDefault("aliases.shuffleoff", []string{"shuffleoff", "shufoff", "shoff"})

	// Permissions configuration
	viper.SetDefault("permissions.adminsenabled", true)
	viper.SetDefault("permissions.admins", []string{"Matt"})
	viper.SetDefault("permissions.add", false)
	viper.SetDefault("permissions.skip", false)
	viper.SetDefault("permissions.help", false)
	viper.SetDefault("permissions.volume", false)
	viper.SetDefault("permissions.move", true)
	viper.SetDefault("permissions.reload", true)
	viper.SetDefault("permissions.reset", true)
	viper.SetDefault("permissions.numtracks", false)
	viper.SetDefault("permissions.nexttrack", false)
	viper.SetDefault("permissions.currenttrack", false)
	viper.SetDefault("permissions.setcomment", true)
	viper.SetDefault("permissions.numcached", true)
	viper.SetDefault("permissions.cachesize", true)
	viper.SetDefault("permissions.kill", true)
	viper.SetDefault("permissions.shuffle", true)
	viper.SetDefault("permissions.shuffletoggle", true)
}
