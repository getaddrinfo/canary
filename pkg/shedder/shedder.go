package shedder

import (
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
)

type Handler interface {
	// Shed returns true if the request should be shedded to a non-canary instance
	Shed(downstream.Target) bool

	// Track reads a response and may move the Handler towards a "shedding" state.
	Track(*http.Response, error) error
}
