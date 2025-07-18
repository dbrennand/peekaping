package executor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"peekaping/src/utils"
)

func GenericValidator[T any](cfg *T) error {
	return utils.Validate.Struct(cfg)
}

func GenericUnmarshal[T any](configJSON string) (*T, error) {
	var cfg T
	dec := json.NewDecoder(bytes.NewReader([]byte(configJSON)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return &cfg, nil
}
