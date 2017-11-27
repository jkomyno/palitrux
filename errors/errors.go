package errors

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	NotFound uint8 = iota
	NotAllowed
	Unsupported
	Unavailable
	BadRequest
	NotImplemented
	InternalError
)

var (
	ErrorNotFound            = New("Not found", NotFound)
	ErrorMethodNotAllowed    = New("Method not allowed", NotAllowed)
	ErrorMIMETypeUnsupported = New("MIME type not supported", Unsupported)
	ErrorLimiter             = New("Rate limit exceed", Unavailable)
	ErrorMissingParam        = New("Missing required params", BadRequest)
	ErrorNotImplemented      = New("This endpoint isn't implemented", NotImplemented)
	ErrorFileSizeExceeded    = New("The uploaded file size is too big to be processed", BadRequest)
	ErrorContentType         = New("The Content-Type header is wrong or not set", BadRequest)
)

// Error describes the error structure
type Error struct {
	Message string `json:"message,omitempty"`
	Code    uint8  `json:"code"`
}

// JSON encode the error to JSON
func (e Error) JSON() []byte {
	eJSON, _ := json.Marshal(e)
	return eJSON
}

// HTTPCode returns the proper HTTP Status Code
func (e Error) HTTPCode() int {
	switch e.Code {
	case BadRequest:
		return http.StatusBadRequest
	case NotAllowed:
		return http.StatusMethodNotAllowed
	case Unsupported:
		return http.StatusUnsupportedMediaType
	case InternalError:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case NotImplemented:
		return http.StatusNotImplemented
	default:
		return http.StatusServiceUnavailable
	}
}

// New returns a new Error
func New(err string, code uint8) Error {
	err = strings.Replace(err, "\n", "", -1)
	return Error{
		Message: err,
		Code:    code,
	}
}

// ReplyWithError replies to the client with a JSON describing the error
func ReplyWithError(w http.ResponseWriter, err Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPCode())
	w.Write(err.JSON())
}
