package heapqueue

import (
	"errors"
	"fmt"
	"reflect"
)

// IntEntry is an example implementation for the EntryI interface for storing
// entries containing integer priorities and keys

// Entries can be retrieved like so:
/*
func exampleRetrieve(myHeap HeapQueue){
	retrievedEntry := myHeap.Pop().(IntEntry)
	fmt.Println(retrievedEntry)
}
*/

type IntEntry struct {
	Priority int
	Key      int
}

func (e IntEntry) Cmp(otherEntry interface{}) (int, error) {
	b, ok := otherEntry.(IntEntry)
	if !ok {
		return 0, errors.New("unable to compare IntEntry to " + reflect.TypeOf(otherEntry).String())
	}
	if e.Priority > b.Priority || (e.Priority == b.Priority && e.Key > b.Key) {
		return 1, nil
	}
	if e.Priority == b.Priority && e.Key == b.Key {
		return 0, nil
	}
	return -1, nil
}

func (e IntEntry) String() string {
	return fmt.Sprintf("Priority: %d Key: %d", e.Priority, e.Key)
}
