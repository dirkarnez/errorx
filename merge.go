package errorx

import (
	"strings"
	"errors"
)

func MergeError(errList ...error) error {
	if len(errList) < 1 {
		return nil
	}

	s := make([]string, len(errList))

	for idx, currentError := range errList {
		s[idx] = currentError.Error()
	}

	return errors.New(strings.Join(s, ", "))
}