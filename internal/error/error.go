package error

import (
	"fmt"
	"net/http"
)

const Template = `Something went wrong while canary processed the request: %s

If you are a user, try again in a few minutes.

If you are an administrator, refer to the documentation for more information (code = %s)`

type Code string

type Error struct {
	Code    Code
	Message string
}

const (
	codeDestinationUnknown         Code = "DestinationUnknown"
	codeNotConfigured              Code = "NotConfigured"
	codeDestinationDerivationError Code = "DestinationDerivationFailed"
	codeFatalUnknown               Code = "FatalUnknown"
)

var ErrDestinationUnknown = Error{codeDestinationUnknown, "No destination matches"}
var ErrDestinationDerivationError = Error{codeDestinationDerivationError, "Error while deriving destination"}
var ErrUnknownFatal = Error{codeFatalUnknown, "Unknown fatal exception"}
var ErrNotConfigured = Error{codeNotConfigured, "Not configured yet"}

func Format(e Error) string {
	return fmt.Sprintf(Template, e.Message, e.Code)
}

func (e Error) String() string {
	return Format(e)
}

func (e Error) Write(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write([]byte(e.String()))
}
