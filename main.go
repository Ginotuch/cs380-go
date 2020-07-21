package main

import (
	"fmt"
	"github.com/Ginotuch/heapqueue/heapqueue"
)

func main() {
	newheap := heapqueue.SetupHeapQueue()
	newheap.Push(heapqueue.IntEntry{Priority: 1, Key: 2})
	newheap.Push(heapqueue.IntEntry{})
	for newheap.Size() > 0{
		fmt.Println(newheap.Pop().(heapqueue.IntEntry))
	}

}