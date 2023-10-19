package server

import (
	"github.com/getaddrinfo/canary/pkg/downstream"
	"github.com/getaddrinfo/canary/pkg/self"
	"github.com/getaddrinfo/canary/pkg/shedder"
	"github.com/getaddrinfo/canary/pkg/strategy"
)

type HttpConfig struct {
	Port uint16
	Host string
}

type Opts struct {
	// HTTP server config
	HttpConfig HttpConfig

	// Strategy, DefaultTarget and LoadShedder get combined into a single
	// strategy.Strategy, which handles all of these behaviours.
	Strategy      strategy.Strategy
	DefaultTarget downstream.Target
	LoadShedder   shedder.Handler

	// Generates request ids for requests to the downstream target,
	// ignored if nil.
	RequestIdGenerator self.RequestIdStrategy

	// Slice of functions that mutate the inbound http.Request
	// before being forwarded to the downstream target.
	InboundMutatorFuncs []InboundMutator

	// Slice of functions that mutate the received http.Response
	// from the downstream service before being written back to the
	// client.
	OutboundMutatorFuncs []OutboundMutator
}

func DefaultOpts() *Opts {
	return &Opts{
		HttpConfig: HttpConfig{
			Port: 19945,
			Host: "localhost",
		},
		Strategy:           defaultStrategy{},
		LoadShedder:        nil,
		RequestIdGenerator: nil,
	}
}
