package strategy

import (
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
)

// Strategy represents a way of resolving to a target.
//
// Generally, user implemented targets are then wrapped by an internal
// Strategy that handles the case that there is an error determining
// the strategy, or where you want to route the user to none of the
// canaries specified (i.e., to the non-canary instance or lb)
//
// Strategies can be chained together, where each Strategy involves a small
// behaviour (e.g., returning a default strategy if none is found) to make
// composable strategies - see `pkg/strategy/fallback.go` as an example.
type Strategy interface {

	// For takes a *http.Request and resolves a
	// target to send it to. Implementers can decide
	// their behaviour e.g., sticky canary, cookie canary, random, etc/
	//
	// If the strategy wants to send to the default destination,
	// return `nil` for `downstream.Target` - the handler calling it
	// will return a default destination.
	For(*http.Request) (downstream.Target, error)
}
