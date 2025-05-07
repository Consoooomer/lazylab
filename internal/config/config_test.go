package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigValidate(t *testing.T) {
	testCases := []struct {
		config Config
		isErr  bool
	}{}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			err := test.config.validate()
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLabConfigValidate(t *testing.T) {
	testCases := []struct {
		config LabConfig
		isErr  bool
	}{}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			err := test.config.validate()
			if test.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
