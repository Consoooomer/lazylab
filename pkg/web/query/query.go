package query

import "sync"

type (
	QueryName  string
	QueryValue string
)

type Query struct {
	queryArgs map[QueryName]QueryValue
	mu        sync.RWMutex
}

func New(query map[string]string) *Query {
	return &Query{
		mu: sync.RWMutex{},
	}
}

func (q *Query) Add(key, value string) {
}

func (q *Query) ToString() string {
	return ""
}
