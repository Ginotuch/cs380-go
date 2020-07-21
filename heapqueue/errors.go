package heapqueue

import "errors"

var (
	ErrEmpty             = errors.New("heap is empty")
	ErrIncompatibleTypes = errors.New("incompatible entry types exist in heap")
)
