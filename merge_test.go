package errorx

import (
	"testing"
	"errors"
)

func TestMergeError(t *testing.T) {
	errorMerged := MergeError([]error{
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	}...)

	if errorMerged == nil {
		t.Errorf("Cannot merge errors")
	}
}