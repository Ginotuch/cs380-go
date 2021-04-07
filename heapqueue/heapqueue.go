package heapqueue

// The interface HeapQueue uses to handle elements in the heap.
// Use either a pre-made implementation or write your own.
// Look at intentry.go for an example.
type EntryI interface {
	Cmp(otherEntry interface{}) int
	GetKey() interface{}
}

type HeapQueue interface {
	Peek() (EntryI, error)
	Pop() EntryI
	Push(newEntry EntryI)
	Size() int
}

type heapqueue struct {
	array     []EntryI
	positions map[interface{}]int
}

func (q *heapqueue) Size() int {
	return len(q.array)
}

func (q *heapqueue) Pop() EntryI {
	temp := q.delete(0)
	return temp
}

func (q *heapqueue) Push(newEntry EntryI) {
	if position, ok := q.positions[newEntry.GetKey()]; ok {
		if newEntry.Cmp(q.array[position]) == 1 {
			q.array[position] = newEntry
			q.heapifyDown(position)
		} else if newEntry.Cmp(q.array[position]) == -1 {
			q.array[position] = newEntry
			q.heapifyUp(position)
		}
	} else {
		q.positions[newEntry.GetKey()] = q.Size()
		q.array = append(q.array, newEntry)
		q.heapifyUp(q.Size() - 1)
	}
}

func (q *heapqueue) Peek() (EntryI, error) {
	if q.Size() > 0 {
		return q.array[0], nil
	}
	return nil, ErrorEmpty
}

func (q *heapqueue) swap(parentIndex int, childIndex int) bool {
	parent := q.array[parentIndex]
	child := q.array[childIndex]

	if parent.Cmp(child) == 1 {
		q.positions[parent.GetKey()] = childIndex
		q.positions[child.GetKey()] = parentIndex

		q.array[parentIndex] = child
		q.array[childIndex] = parent
		return true
	}
	return false
}

func (q *heapqueue) heapifyDown(i int) {
	for 2*i+1 < len(q.array) {
		leftChildI := 2*i + 1
		rightChildI := 2*i + 2
		bigChildI := leftChildI
		if rightChildI < len(q.array) && q.array[leftChildI].Cmp(q.array[rightChildI]) >= 0 {
			bigChildI = rightChildI
		}
		if !q.swap(i, bigChildI) {
			return
		}
		i = bigChildI
	}

}

func (q *heapqueue) heapifyUp(i int) {
	for i > 0 {
		parentIndex := (i - 1) >> 1
		if !q.swap(parentIndex, i) {
			return
		}
		i = parentIndex
	}
}

func (q *heapqueue) delete(i int) EntryI {
	deleted := q.array[i]
	oldLeaf := q.array[len(q.array)-1]
	q.array[i] = oldLeaf

	q.positions[q.array[i].GetKey()] = i
	delete(q.positions, deleted.GetKey())

	q.array = q.array[:len(q.array)-1]

	if deleted.Cmp(oldLeaf) == 1 && i < len(q.array) {
		q.heapifyUp(i)
	} else {
		q.heapifyDown(i)
	}
	return deleted
}

func SetupHeapQueue() HeapQueue {
	return &heapqueue{positions: make(map[interface{}]int)}
}
