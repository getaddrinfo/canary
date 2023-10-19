package matcher

import "github.com/getaddrinfo/canary/pkg/downstream"

type Matcher[T any] interface {
	func(v T) downstream.Target
}
