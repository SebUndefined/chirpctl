package validation

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func ValidateHexLen(s string, length int) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return fmt.Errorf("invalid hex string %q: %w", s, err)
	}
	return ValidateSliceLen[byte](b, length)
}

func ValidateUUID(s string) error {
	_, err := uuid.Parse(s)
	return err
}

func ValidateSliceLen[T any](s []T, length int) error {
	if len(s) != length {
		return fmt.Errorf("expected %d elements, got %d", length, len(s))
	}
	return nil
}

func ValidateBoolStrict(s string) error {
	if s != "true" && s != "false" {
		return fmt.Errorf("invalid bool %q: must be \"true\" or \"false\"", s)
	}
	return nil
}

func NotEmptyString(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("empty string is not allowed")
	}
	return nil
}
