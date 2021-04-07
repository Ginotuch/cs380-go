package heapqueue

import (
	"errors"
	"fmt"
)

var (
	ErrorEmpty = errors.New("heap is empty")
)

type IncompatibleTypesError struct {
	SelfType  interface{}
	OtherType interface{}
}

func (i IncompatibleTypesError) Error() string {
	return fmt.Sprintf("unable to compare %T type to %T", i.SelfType, i.OtherType)
}
