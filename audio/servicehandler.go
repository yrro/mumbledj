/*
 * MumbleDJ
 * By Matthieu Grieger
 * audio/servicehandler.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package audio

// ServiceHandler is a struct that resolves service types via regex on the
// supplied URL, and returns tracks associated with the URL.
type ServiceHandler struct {
}

// GetTracks returns the audio tracks associated with the given URL, if any
// exist.
func (h *ServiceHandler) GetTracks(url string) ([]Track, error) {
	return nil, nil
}
