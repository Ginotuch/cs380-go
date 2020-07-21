package heapqueue

import (
	"fmt"
	"log"
)

type IntEntry struct {
	Priority int
	Key int
}

func (e IntEntry) Cmp(otherEntry interface{}) int {
	b, ok := otherEntry.(IntEntry)
	if !ok {
		log.Fatal("Unable to convert otherEntry to ")
	}
	if e.Priority > b.Priority || (e.Priority == b.Priority && e.Key > b.Key) {
		return 1
	}
	if e.Priority == b.Priority && e.Key == b.Key {
		return 0
	}
	return -1
}

func (e IntEntry) String() string{
	return fmt.Sprintf("Priority: %d Key: %d", e.Priority, e.Key)
}