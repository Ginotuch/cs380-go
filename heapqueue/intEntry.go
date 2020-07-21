package heapqueue

import (
	"fmt"
	"log"
)

// IntEntry is an example implementation for the EntryI interface for storing
// entries containing integer priorities and keys

// Entries can
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
		log.Fatal("Unable to convert otherEntry to IntEntry")
	}
	if e.Priority > b.Priority || (e.Priority == b.Priority && e.Key > b.Key) {
		return 1
	}
	if e.Priority == b.Priority && e.Key == b.Key {
		return 0
	}
	return -1
}

func (e IntEntry) String() string {
	return fmt.Sprintf("Priority: %d Key: %d", e.Priority, e.Key)
}
