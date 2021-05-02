package httpHelper

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	defaultContentType = "application/json; charset=utf-8"
	defaultHTTPCode    = http.StatusInternalServerError

	CtxTransport = "transport"
)

// ctxKey type definision for passing context values
type ctxKey string

// HTTPError satisfies the Headerer and StatusCoder interfaces in
// package kit/transport/http.
type HTTPError struct {
	error
	statusCode int
	headers    http.Header
}

// StatusCode return the status code of the error
func (h HTTPError) StatusCode() int {
	return h.statusCode
}

// Headers return the http headers
func (h HTTPError) Headers() http.Header {
	return h.headers
}

type errorWrapper struct {
	Error string `json:"error"`
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}

	w.Header().Set("Content-Type", defaultContentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}

	code := defaultHTTPCode
	if sc, ok := err.(httptransport.StatusCoder); ok {
		code = sc.StatusCode()
	}

	w.WriteHeader(code)
	w.Write(body)
}

func HeadersToContext(ctx context.Context, r *http.Request) context.Context {
	ctx = context.WithValue(ctx, ctxKey(CtxTransport), "HTTP")

	for k := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, ctxKey(k), r.Header.Get(k))
		ctx = context.WithValue(ctx, ctxKey(strings.ToLower(k)), r.Header.Get(k))
	}

	return ctx
}
