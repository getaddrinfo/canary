package shedder

import (
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
	"github.com/getaddrinfo/canary/pkg/strategy"
)

type awareTargeter struct {
	strat    strategy.Strategy
	default_ downstream.Target
	shedder  Handler
}

func (f *awareTargeter) For(r *http.Request) (downstream.Target, error) {
	target, err := f.strat.For(r)

	// something went wrong, but we don't want to
	// throw away this request, so just route it
	// to the default target
	if err != nil {
		return f.default_, err
	}

	// if the shedder wants us away from this target,
	// return the default and a nil error
	if f.shedder.Shed(target) {
		return f.default_, nil
	}

	// if the strategy returned nil (no canary for this request),
	// then return the default strategy
	if target == nil {
		return f.default_, nil
	}

	return &awareTarget{
		target:  target,
		handler: f.shedder,
	}, nil
}

func NewAwareTargeter(
	strat strategy.Strategy,
	fallback downstream.Target,
	shed Handler,
) *awareTargeter {
	return &awareTargeter{strat, fallback, shed}
}

var _ strategy.Strategy = (*awareTargeter)(nil)
