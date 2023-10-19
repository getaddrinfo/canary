package shedder

import (
	"net/http"

	"github.com/getaddrinfo/canary/pkg/downstream"
)

type awareTarget struct {
	target  downstream.Target
	handler Handler
}

func (a *awareTarget) Send(r *http.Request) (*http.Response, error) {
	res, err := a.target.Send(r)

	trackerr := a.handler.Track(res, err)

	if trackerr != nil {
		return res, trackerr
	}

	return res, err
}

func (a *awareTarget) Identity() string {
	return a.target.Identity()
}

var _ downstream.Target = (*awareTarget)(nil)
