package server

import (
	"errors"
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
	"github.com/getaddrinfo/canary/pkg/strategy"
)

var ErrNoTargeter = errors.New("no targeter specified")

type defaultStrategy struct{}

func (t defaultStrategy) For(r *http.Request) (downstream.Target, error) {
	return nil, ErrNoTargeter
}

var _ strategy.Strategy = (*defaultStrategy)(nil)
