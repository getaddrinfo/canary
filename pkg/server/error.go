package server

import (
	"errors"
	"net/http"

	canaryerr "github.com/getaddrinfo/canary/internal/error"
	"github.com/getaddrinfo/canary/pkg/strategy"
)

func handleError(
	rw http.ResponseWriter,
	err error,
) {
	if err == nil {
		return
	}

	if errors.Is(err, strategy.ErrNoDownstreamMatch) {
		canaryerr.ErrDestinationUnknown.Write(rw)
		return
	}

	if errors.Is(err, ErrNoTargeter) {
		canaryerr.ErrNotConfigured.Write(rw)
		return
	}

	canaryerr.ErrUnknownFatal.Write(rw)
}
