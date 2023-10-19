package server

import (
	"io"
	"net/http"
)

type OutboundMutator func(r *http.Response)

func writeResponse(w http.ResponseWriter, res *http.Response) error {
	for name, values := range res.Header {
		w.Header().Set(name, values[0])
	}

	w.WriteHeader(res.StatusCode)

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	_, err = w.Write(body)

	if err != nil {
		return err
	}

	return nil
}
