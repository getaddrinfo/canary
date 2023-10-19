package downstream

import (
	"net/http"
)

// Target represents a downstream http server that may be sent to.
type Target interface {
	// Send sends a request to the downstream target
	// the error returned represents if making the request
	// failed (e.g., tls failed), _not_ if the request
	// was bad - this is the responsibility of downstream.ErrorHandler
	Send(*http.Request) (*http.Response, error)

	// Identity returns a unique string (e.g., the IP address, the hostname) for
	// this Target. Must be unique within the context of a Strategy
	Identity() string
}
