package errorx

import (
	"strings"
	"errors"
)

func MergeError(errList ...error) error {
	if errList == nil {
		return nil
	}

	length := len(errList)

	if length < 1 || (length == 1 && errList[0] == nil) {
		return nil
	}

	s := make([]string, len(errList))

	for idx, currentError := range errList {
		s[idx] = currentError.Error()
	}

	return errors.New(strings.Join(s, ", "))
}