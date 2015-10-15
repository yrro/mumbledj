/*
 * MumbleDJ
 * By Matthieu Grieger
 * config.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package main

// GeneralConfig contains configuration variables for general options.
type GeneralConfig struct {
	CommandPrefix      string
	SkipRatio          float32
	PlaylistSkipRatio  float32
	DefaultComment     string
	DefaultChannel     []string
	MaxTrackDuration   int
	AutomaticShuffleOn bool
}

// ConnectionConfig contains configuration variables for connection options.
type ConnectionConfig struct {
	ServerAddress  string
	ServerPort     int
	ServerPassword string
	RetryEnabled   bool
	RetryAttempts  int
	RetryInterval  int
}

// CacheConfig contains configuration variables for cache options.
type CacheConfig struct {
	Enabled              bool
	MaximumSizeMebibytes int64
	ExpireTime           float64
	CheckInterval        int
}

// VolumeConfig contains configuration variables for volume options.
type VolumeConfig struct {
	DefaultVolume float32
	LowestVolume  float32
	HighestVolume float32
}

// AliasConfig contains configuration variables for command alias options.
type AliasConfig struct {
	AddAliases               []string
	SkipAliases              []string
	SkipPlaylistAliases      []string
	AdminSkipAliases         []string
	AdminSkipPlaylistAliases []string
	HelpAliases              []string
	VolumeAliases            []string
	MoveAliases              []string
	ReloadAliases            []string
	ResetAliases             []string
	NumTracksAliases         []string
	NextTrackAliases         []string
	CurrentTrackAliases      []string
	SetCommentAliases        []string
	NumCachedAliases         []string
	CacheSizeAliases         []string
	KillAliases              []string
	ShuffleAliases           []string
	ShuffleOnAliases         []string
	ShuffleOffAliases        []string
}

// PermissionsConfig contains configuration variables for permissions options.
type PermissionsConfig struct {
	AdminsEnabled      bool
	Admins             []string
	AdminAdd           bool
	AdminAddPlaylists  bool
	AdminSkip          bool
	AdminHelp          bool
	AdminVolume        bool
	AdminMove          bool
	AdminReload        bool
	AdminReset         bool
	AdminNumTracks     bool
	AdminNextTrack     bool
	AdminCurrentTrack  bool
	AdminSetComment    bool
	AdminNumCached     bool
	AdminCacheSize     bool
	AdminKill          bool
	AdminShuffle       bool
	AdminShuffleToggle bool
}

// BotConfig is a configuration struct that houses all the above configurations.
type BotConfig struct {
	General     GeneralConfig
	Connection  ConnectionConfig
	Cache       CacheConfig
	Volume      VolumeConfig
	Aliases     AliasConfig
	Permissions PermissionsConfig
}

// NewDefaultConfig returns a BotConfig filled with default values.
func NewDefaultConfig() *BotConfig {
	return &BotConfig{
		General: GeneralConfig{
			CommandPrefix:      "!",
			SkipRatio:          0.5,
			PlaylistSkipRatio:  0.5,
			DefaultComment:     "Hello! I am a bot. Type !herlp for a list of commands.",
			MaxTrackDuration:   0,
			AutomaticShuffleOn: false,
		},
		Connection: ConnectionConfig{
			ServerAddress:  "127.0.0.1",
			ServerPort:     64738,
			ServerPassword: "",
			RetryEnabled:   true,
			RetryAttempts:  10,
			RetryInterval:  10,
		},
		Cache: CacheConfig{
			Enabled:              false,
			MaximumSizeMebibytes: 512,
			ExpireTime:           24,
			CheckInterval:        5,
		},
		Volume: VolumeConfig{
			DefaultVolume: 0.2,
			LowestVolume:  0.01,
			HighestVolume: 0.8,
		},
		Aliases: AliasConfig{
			AddAliases:               []string{"add", "a"},
			SkipAliases:              []string{"skip", "s"},
			SkipPlaylistAliases:      []string{"skipplaylist", "sp"},
			AdminSkipAliases:         []string{"forceskip", "fs"},
			AdminSkipPlaylistAliases: []string{"forceskipplaylist", "fsp"},
			HelpAliases:              []string{"help", "h"},
			VolumeAliases:            []string{"volume", "v"},
			MoveAliases:              []string{"move", "m"},
			ReloadAliases:            []string{"reload", "r"},
			ResetAliases:             []string{"resetqueue", "reset", "rq"},
			NumTracksAliases:         []string{"numtracks", "numsongs", "nt", "ns"},
			NextTrackAliases:         []string{"nexttrack", "nextsong", "next"},
			CurrentTrackAliases:      []string{"currenttrack", "currentsong", "current"},
			SetCommentAliases:        []string{"setcomment", "comment"},
			NumCachedAliases:         []string{"numcached", "nc"},
			CacheSizeAliases:         []string{"cachesize", "cached"},
			KillAliases:              []string{"kill", "k"},
			ShuffleAliases:           []string{"shuffle", "shuf"},
			ShuffleOnAliases:         []string{"shuffleon", "shufon"},
			ShuffleOffAliases:        []string{"shuffleoff", "shufoff"},
		},
		Permissions: PermissionsConfig{
			AdminsEnabled:      true,
			Admins:             []string{"Matt"},
			AdminAdd:           false,
			AdminAddPlaylists:  false,
			AdminSkip:          false,
			AdminHelp:          false,
			AdminVolume:        false,
			AdminMove:          true,
			AdminReload:        true,
			AdminReset:         true,
			AdminNumTracks:     false,
			AdminNextTrack:     false,
			AdminCurrentTrack:  false,
			AdminSetComment:    true,
			AdminNumCached:     true,
			AdminCacheSize:     true,
			AdminKill:          true,
			AdminShuffle:       true,
			AdminShuffleToggle: true,
		},
	}
}
