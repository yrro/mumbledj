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
	AvailableServices []string
}

// AddService adds a service name to the AvailableServices struct.
func (h *ServiceHandler) AddService(service string) {
	h.AvailableServices = append(h.AvailableServices, service)
}

// GetAvailableServices returns a slice containing the names of the enabled services.
func (h *ServiceHandler) GetAvailableServices() []string {
	return h.AvailableServices
}

// GetTracks returns the audio tracks associated with the given URL, if any
// exist.
func (h *ServiceHandler) GetTracks(url string) ([]Track, error) {
	return nil, nil
}
