package strategy

import (
	"errors"
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
)

var ErrNoDownstreamMatch = errors.New("no downstream for given request")

type ErrorIfNoDestination struct {
	Strategy Strategy
}

func (e *ErrorIfNoDestination) For(r *http.Request) (downstream.Target, error) {
	ds, err := e.Strategy.For(r)

	// this may be the case if the fallback
	// handler has made sure there is a destination
	if err != nil {
		return ds, err
	}

	if ds == nil {
		return nil, ErrNoDownstreamMatch
	}

	return ds, nil
}

var _ Strategy = (*ErrorIfNoDestination)(nil)

func NewErrorIfNoDestination(strategy Strategy) Strategy {
	return &ErrorIfNoDestination{Strategy: strategy}
}
