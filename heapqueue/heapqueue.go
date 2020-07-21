package heapqueue


type EntryI interface {
	Cmp(otherEntry interface{}) int
}


type HeapQueue interface {
	Pop() EntryI
	Push(newEntry EntryI)
	Size() int
}

type heapqueue struct {
	Array []EntryI
}

func (q *heapqueue) Size() int {
	return len(q.Array)
}

func (q *heapqueue) Pop() EntryI {
	temp := q.delete(0)
	return temp
}

func (q *heapqueue) Push(newEntry EntryI) {
	q.Array = append(q.Array, newEntry)
	q.heapifyUp(len(q.Array) - 1)
}



func (q *heapqueue) swap(parentIndex int, childIndex int) bool {
	parent := q.Array[parentIndex]
	child := q.Array[childIndex]

	if parent.Cmp(child) == 1 {
		q.Array[parentIndex] = child
		q.Array[childIndex] = parent
		return true
	}
	return false
}

func (q *heapqueue) heapifyDown(i int) {
	for 2*i+1 < len(q.Array) {
		leftChildI := 2*i + 1
		rightChildI := 2*i + 2
		bigChildI := leftChildI
		if rightChildI < len(q.Array) && q.Array[leftChildI].Cmp(q.Array[rightChildI]) >= 0 {
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
	deleted := q.Array[i]
	oldLeaf := q.Array[len(q.Array)-1]
	q.Array[i] = oldLeaf

	q.Array = q.Array[:len(q.Array)-1]

	if deleted.Cmp(oldLeaf) == 1 && i < len(q.Array) {
		q.heapifyUp(i)
	} else {
		q.heapifyDown(i)
	}
	return deleted
}



func SetupHeapQueue() HeapQueue {
	return &heapqueue{}
}
