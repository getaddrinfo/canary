package server

import (
	"net/http"

	"github.com/getaddrinfo/canary/pkg/self"
)

type InboundMutator func(r *http.Request)

func makeRequestId(
	req *http.Request,
	strat self.RequestIdStrategy,
) string {
	if strat == nil {
		return ""
	}

	id := strat.Generate(req)
	req.Header.Set("X-Request-ID", id)

	return id
}

func addResponseId(
	res *http.Response,
	id string,
) {
	if id == "" {
		return
	}

	if res.Header.Get("X-Request-ID") == "" {
		res.Header.Set("X-Request-ID", id)
	}
}
