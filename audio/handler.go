/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/handler.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package audio

// Handler is an interface for a struct that handles an incoming URL and
// returns a slice of tracks associated with it.
type Handler interface {
	GetTracks(string) ([]Track, error)
}
