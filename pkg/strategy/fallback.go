package strategy

import (
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
)

type Fallback struct {
	Default  downstream.Target
	Strategy Strategy
}

func (f *Fallback) For(r *http.Request) (downstream.Target, error) {
	ds, err := f.Strategy.For(r)

	if err != nil {
		return f.Default, err
	}

	if ds == nil {
		return f.Default, nil
	}

	return ds, nil
}

var _ Strategy = (*Fallback)(nil)

func NewFallback(
	default_ downstream.Target,
	strategy Strategy,
) Strategy {
	return &Fallback{
		Default:  default_,
		Strategy: strategy,
	}
}
