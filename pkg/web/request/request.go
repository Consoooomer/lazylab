package request

import (
	"net/http"

	"github.com/loikx/lazylab/pkg/web/endpoint"
	"github.com/loikx/lazylab/pkg/web/headers"
	"github.com/loikx/lazylab/pkg/web/query"
)

type Request struct{}

func New() *Request {
	return nil
}

func (r *Request) ToURL() string {
	return ""
}

func (r *Request) ToHttpRequest() *http.Request {
	return nil
}

func (r *Request) WithEndpoint(endpointURL endpoint.Endpoint) {
}

func (r *Request) WithHeaders(headers *headers.Headers) {
}

func (r *Request) WithQuery(query *query.Query) {
}
