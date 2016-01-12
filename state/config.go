/*
 * MumbleDJ
 * By Matthieu Grieger
 * state/config.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package state

import "github.com/spf13/viper"

// Config is a struct that gathers all logic related to configuration via
// environment variables, commandline arguments, and configuration files.
type Config struct {
	ConfigFileLocation string
}

// SetDefaultConfiguration initalizes a Viper configuration with the default
// values.
func (c *Config) SetDefaultConfiguration() {
	// General configuration
	viper.SetDefault("general.commandprefix", "!")
	viper.SetDefault("general.skipratio", 0.5)
	viper.SetDefault("general.playlistskipratio", 0.5)
	viper.SetDefault("general.defaultcomment", "Hello! I am a bot. Type !help for a list of commands.")
	viper.SetDefault("general.defaultchannel", []string{""})
	viper.SetDefault("general.maxtrackduration", 0)
	viper.SetDefault("general.maxtracksperplaylist", 50)
	viper.SetDefault("general.automaticshuffleon", false)
	viper.SetDefault("general.playercommand", "ffmpeg")

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
	viper.SetDefault("aliases.forceskip", []string{"forceskip", "fs"})
	viper.SetDefault("aliases.forceskipplaylist", []string{"forceskipplaylist", "fsp"})
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
	viper.SetDefault("aliases.listtracks", []string{"listtracks", "listsongs", "list", "l"})

	// Permissions configuration
	viper.SetDefault("permissions.adminsenabled", true)
	viper.SetDefault("permissions.admins", []string{"Matt"})
	viper.SetDefault("permissions.add", false)
	viper.SetDefault("permissions.skip", false)
	viper.SetDefault("permissions.skipplaylist", false)
	viper.SetDefault("permissions.forceskip", true)
	viper.SetDefault("permissions.forceskipplaylist", true)
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
	viper.SetDefault("permissions.listtracks", false)

	// Command description configuration
	viper.SetDefault("descriptions.add", "Adds a track or playlist from YouTube or SoundCloud to the audio queue.")
	viper.SetDefault("descriptions.skip", "Places a vote to skip the current track.")
	viper.SetDefault("descriptions.skipplaylist", "Places a vote to skip the current playlist.")
	viper.SetDefault("descriptions.forceskip", "Immediately skips the current track.")
	viper.SetDefault("descriptions.forceskipplaylist", "Immediately skips the current playlist.")
	viper.SetDefault("descriptions.help", "Outputs this list of commands.")
	viper.SetDefault("descriptions.volume", "Changes the volume if an argument is provided, outputs the current volume otherwise.")
	viper.SetDefault("descriptions.move", "Moves the bot into the Mumble channel provided via argument.")
	viper.SetDefault("descriptions.reload", "Reloads the configuration file.")
	viper.SetDefault("descriptions.reset", "Resets the audio queue by removing all queue items.")
	viper.SetDefault("descriptions.numtracks", "Outputs the number of tracks currently in the audio queue.")
	viper.SetDefault("descriptions.nexttrack", "Outputs the title and submitter of the next track in the queue if one exists.")
	viper.SetDefault("descriptions.currenttrack", "Outputs the title and submitter of the current track if one exists.")
	viper.SetDefault("descriptions.setcomment", "Sets the comment displayed next to MumbleDJ's username in Mumble.")
	viper.SetDefault("descriptions.numcached", "Outputs the number of tracks cached on disk if caching is enabled.")
	viper.SetDefault("descriptions.cachesize", "Outputs the file size of the cache in MiB if caching is enabled.")
	viper.SetDefault("descriptions.kill", "Stops the bot and cleans its cache directory.")
	viper.SetDefault("descriptions.shuffle", "Randomizes the tracks currently in the audio queue.")
	viper.SetDefault("descriptions.shuffleon", "Toggles permanent track shuffling on.")
	viper.SetDefault("descriptions.shuffleoff", "Toggles permanent track shuffling off.")
	viper.SetDefault("descriptions.listtracks", "Outputs a list of the tracks currently in the queue.")
}

// LoadFromConfigFile loads configuration values from the filepath specified via
// the filepath argument.
func (c *Config) LoadFromConfigFile(filepath string) error {
	return nil
}

// LoadFromEnvironmentVariables loads configuration values from environment
// variables.
func (c *Config) LoadFromEnvironmentVariables() error {
	return nil
}

// LoadFromCommandline loads configuration values from the commandline.
func (c *Config) LoadFromCommandline() error {
	return nil
}
