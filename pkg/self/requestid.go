package self

import (
	"net/http"

	"github.com/google/uuid"
)

type RequestIdStrategy interface {
	Generate(r *http.Request) string
}

// UUID v4
type UUIDv4RequestIdStrategy struct {
	strict bool
}

func (strat UUIDv4RequestIdStrategy) Generate(r *http.Request) string {
	val := r.Header.Get("X-Request-ID")

	if val != "" {
		return uuid.NewString()
	}

	if strat.strict && !isValidUuid(val) {
		return uuid.NewString()
	}

	return val
}

// NewUUIDv4RequestIdStrategy returns a new UUIDv4RequestIdStrategy
// If strict is set to true, any passed `X-Request-ID` must be a valid
// UUID.
func NewUUIDv4RequestIdStrategy(strict bool) *UUIDv4RequestIdStrategy {
	return &UUIDv4RequestIdStrategy{strict}
}

var _ RequestIdStrategy = (*UUIDv4RequestIdStrategy)(nil)

// helper fns
func isValidUuid(val string) bool {
	_, err := uuid.Parse(val)

	return err == nil
}
