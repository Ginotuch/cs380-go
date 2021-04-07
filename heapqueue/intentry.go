package heapqueue

import (
	"fmt"
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

func (e IntEntry) Cmp(otherEntry interface{}) int {
	b, ok := otherEntry.(IntEntry)
	if !ok {
		panic(IncompatibleTypesError{SelfType: e, OtherType: otherEntry})
	}
	if e.Priority > b.Priority || (e.Priority == b.Priority && e.Key > b.Key) {
		return 1
	}
	if e.Priority == b.Priority && e.Key == b.Key {
		return 0
	}
	return -1
}

func (e IntEntry) GetKey() interface{} {
	return e.Key
}

func (e IntEntry) String() string {
	return fmt.Sprintf("Priority: %d Key: %d", e.Priority, e.Key)
}
