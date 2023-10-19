package server

import (
	"fmt"
	"net/http"

	"github.com/getaddrinfo/canary/pkg/self"
	"github.com/getaddrinfo/canary/pkg/shedder"

	strategy "github.com/getaddrinfo/canary/pkg/strategy"
)

/*
maybe in future we should consider more advanced
routing capabilities
*/
type Server struct {
	conf     HttpConfig
	targeter strategy.Strategy

	requestIdGenerator self.RequestIdStrategy

	inboundMutatorFuncs  []InboundMutator
	outboundMutatorFuncs []OutboundMutator
}

func (s *Server) OnRequest(rw http.ResponseWriter, req *http.Request) {
	destination, err := s.targeter.For(req)

	// if there was an error, and no destination was found (i.e., the fallback
	// was not in place), handle it.
	if err != nil && destination == nil {
		handleError(rw, err)
		return
	}

	if err != nil {
		// TODO: log (destination exists -> non-fatal)
		func() {}()
	}

	id := makeRequestId(req, s.requestIdGenerator)

	// run all the mutators for sent requests
	for i := range s.inboundMutatorFuncs {
		s.inboundMutatorFuncs[i](req)
	}

	// send the request
	res, err := destination.Send(addContext(req, ContextData{
		RequestID: id,
	}))

	if err != nil {
		handleError(rw, err)
		return
	}

	// run all the response mutators for received response
	for i := range s.outboundMutatorFuncs {
		s.outboundMutatorFuncs[i](res)
	}

	addResponseId(res, id)
	writeResponse(rw, res)
}

func (s *Server) Serve() error {
	return http.ListenAndServe(
		fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port),
		http.HandlerFunc(s.OnRequest),
	)
}

func NewServer(conf *Opts) *Server {
	if conf == nil {
		conf = DefaultOpts()
	}

	var targeter strategy.Strategy

	// TODO: handle this gracefully
	if conf.DefaultTarget == nil {
		func() {}()
	}

	if conf.LoadShedder != nil {
		targeter = shedder.NewAwareTargeter(
			conf.Strategy,
			conf.DefaultTarget,
			conf.LoadShedder,
		)
	} else {
		targeter = strategy.NewFallback(conf.DefaultTarget, conf.Strategy)
	}

	// Wrap them all in a handler that makes sure we get an appropriate error
	// if no destination is found.
	targeter = strategy.NewErrorIfNoDestination(targeter)

	// shutup go build for now
	_ = targeter

	return &Server{
		conf: conf.HttpConfig,

		inboundMutatorFuncs:  conf.InboundMutatorFuncs,
		outboundMutatorFuncs: conf.OutboundMutatorFuncs,

		requestIdGenerator: conf.RequestIdGenerator,
		targeter:           defaultStrategy{},
	}
}
