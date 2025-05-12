package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_ToString(t *testing.T) {
	testCases := []struct {
		name     string
		query    *Query
		expected string
	}{
		{
			"no query keys",
			New(map[string]string{}),
			"",
		},
		{
			"one query arg",
			New(map[string]string{
				"key1": "value1",
			}),
			"?key1=value1",
		},
		{
			"many query arg",
			New(map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			}),
			"?key1=value1&key2=value2&key3=value3",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, _ := test.query.ToString()

			assert.Equal(t, test.expected, got)
		})
	}
}
