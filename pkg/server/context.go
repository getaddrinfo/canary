package server

import (
	"context"
	"net/http"
)

type ContextKey string

const RequestIdContextKey ContextKey = "requestId"

type ContextData struct {
	RequestID string
}

func addContext(req *http.Request, data ContextData) *http.Request {
	return req.WithContext(
		context.WithValue(req.Context(), RequestIdContextKey, data.RequestID),
	)
}
