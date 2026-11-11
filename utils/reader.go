package utils

import (
	"encoding/json"
	"strings"

	"github.com/sebundefined/chirpctl/internal/validation"
)

func ReadOptionalBool(index int, row []string) (bool, bool, error) {
	if index == -1 {
		return false, false, nil
	}
	v := strings.TrimSpace(row[index])
	if err := validation.ValidateBoolStrict(v); err != nil {
		return false, false, err
	}
	return v == "true", true, nil
}

func ReadOptionalMap(index int, row []string) (map[string]string, bool, error) {
	if index < 0 || index >= len(row) {
		return nil, false, nil
	}
	v := strings.TrimSpace(row[index])
	mapsTmp := make(map[string]string)
	if err := json.Unmarshal([]byte(v), &mapsTmp); err != nil {
		return mapsTmp, false, err
	}
	return mapsTmp, true, nil
}
