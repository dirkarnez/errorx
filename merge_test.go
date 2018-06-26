package errorx

import (
	"testing"
	"errors"
	"fmt"
)

func TestMergeErrorNoParam(t *testing.T) {
	fmt.Println("TestMergeErrorNoParam")

	errorMerged := MergeError()

	if errorMerged != nil {
		t.Errorf("Should not return error")
	}
}

func TestMergeErrorNil(t *testing.T) {
	fmt.Println("TestMergeErrorNil")

	errorMerged := MergeError(nil)

	if errorMerged != nil {
		t.Errorf("Should not return error")
	}
}

func TestMergeErrorSingleError(t *testing.T) {
	fmt.Println("TestMergeErrorSingleError")

	errorMerged := MergeError([]error{
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	}...)

	if errorMerged == nil {
		t.Errorf("Cannot merge errors")
	}
}

func TestMergeErrorMultipleErrors(t *testing.T) {
	fmt.Println("TestMergeErrorMultipleErrors")

	errorMerged := MergeError(
		errors.New("error 1"),
	)

	if errorMerged == nil {
		t.Errorf("Cannot merge errors")
	}
}