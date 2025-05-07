package config

import (
	"errors"
	"fmt"
)

var ErrEmptyValue = errors.New("empty value")

type Config struct {
	Configs []LabConfig
}

func (c *Config) validate() error {
	var errs error
	for i := range c.Configs {
		if err := c.Configs[i].validate(); err != nil {
			errs = errors.Join(errs, fmt.Errorf("invalid lab config at index %d: %w", i, err))
		}
	}
	return errs
}

type LabConfig struct {
	LabEndpoint string
	LabKey      string
}

func (lc *LabConfig) validate() error {
	var errs error

	if lc.LabKey == "" {
		errs = errors.Join(errs, fmt.Errorf("invalid lab key: %w", ErrEmptyValue))
	}
	if lc.LabEndpoint == "" {
		errs = errors.Join(errs, fmt.Errorf("invalid lab endpoint: %w", ErrEmptyValue))
	}

	return errs
}
