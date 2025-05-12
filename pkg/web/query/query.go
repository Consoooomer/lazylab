package query

import (
	"fmt"
	"strings"
	"sync"
)

const (
	queryStarter   = '?'
	querySeparator = '&'
	queryEqual     = '='
)

type Query struct {
	queryArgs map[string]string
	mu        sync.RWMutex
}

func New(query map[string]string) *Query {
	return &Query{
		mu:        sync.RWMutex{},
		queryArgs: query,
	}
}

func (q *Query) Reset() {
	q.mu.Lock()
	q.queryArgs = make(map[string]string)
	q.mu.Unlock()
}

func (q *Query) Add(key, value string) {
	q.mu.Lock()
	q.queryArgs[key] = value
	q.mu.Unlock()
}

func (q *Query) ToString() (string, error) {
	if len(q.queryArgs) == 0 {
		return "", nil
	}

	builder := strings.Builder{}
	defer builder.Reset()

	err := builder.WriteByte(queryStarter)
	if err != nil {
		return "", fmt.Errorf("can't write query starter to string: %w", err)
	}

	queryArgsCount := len(q.queryArgs)
	position := 1
	q.mu.RLock()
	for queryName, queryValue := range q.queryArgs {
		_, err = builder.WriteString(queryName)
		if err != nil {
			return "", fmt.Errorf("can't write query key(%s): %w", queryName, err)
		}

		err = builder.WriteByte(queryEqual)
		if err != nil {
			return "", fmt.Errorf("can't write query equal: %w", err)
		}

		_, err = builder.WriteString(queryValue)
		if err != nil {
			return "", fmt.Errorf("can't write query value(%s): %w", queryValue, err)
		}

		if position == queryArgsCount {
			break
		}
		position++
		err = builder.WriteByte(querySeparator)
		if err != nil {
			return "", fmt.Errorf("can't write query separator: %w", err)
		}
	}
	q.mu.RUnlock()

	return builder.String(), nil
}

func (q *Query) Build() (string, error) {
	result, err := q.ToString()
	if err != nil {
		return result, fmt.Errorf("can't build queries: %w", err)
	}
	q.Reset()
	return result, nil
}
