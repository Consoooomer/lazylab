package headers

import "sync"

type (
	HeaderKey   string
	HeaderValue string
)

type Headers struct {
	headers map[HeaderKey]HeaderValue
	mu      sync.RWMutex
}

func New(values map[string]string) *Headers {
	return &Headers{
		mu: sync.RWMutex{},
	}
}

func (h *Headers) Add(key, value string) {
}

func (h *Headers) ToString() string {
	return ""
}
