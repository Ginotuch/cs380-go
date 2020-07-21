package heapqueue

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

var seed int64 = 1337420

func TestHeapqueue_Push(t *testing.T) {
	heapInstance := SetupHeapQueue()
	listToTest := []IntEntry{
		{1, 1},
		{3, 3},
		{5, 5},
		{6, 6},
		{8, 8},
	}
	for _, entry := range listToTest {
		heapInstance.Push(entry)
		checkHeap(t, heapInstance)
	}
}

func TestHeapqueue_PushRandom(t *testing.T) {
	rand.Seed(seed)
	heapInstance := SetupHeapQueue()
	listToTest, _ := getRandomList(t)
	for _, entry := range listToTest {
		heapInstance.Push(entry)
		checkHeap(t, heapInstance)
	}
}

func TestHeapqueue_Pop(t *testing.T) {
	heapInstance := SetupHeapQueue()
	listToTest := []IntEntry{
		{1, 1},
		{3, 3},
		{5, 5},
		{6, 6},
		{8, 8},
	}
	for _, entry := range listToTest {
		heapInstance.Push(entry)
	}
	for _, entry := range listToTest {
		if heapInstance.Pop().Cmp(entry) != 0 {
			t.Errorf("Not popped in correct order")
		}
		checkHeap(t, heapInstance)
	}
}

func TestHeapqueue_PopRandom(t *testing.T) {
	rand.Seed(seed)
	heapInstance := SetupHeapQueue()
	listToTest, _ := getRandomList(t)
	for _, entry := range listToTest {
		heapInstance.Push(entry)
	}
	for _, entry := range sortIntEntries(t, listToTest) {
		if heapInstance.Pop().Cmp(entry) != 0 {
			t.Errorf("Not popped in correct order")
			t.FailNow()
		}
		checkHeap(t, heapInstance)
	}
}

func TestHeapqueue_Peek(t *testing.T) {
	heapInstance := SetupHeapQueue()
	listToTest := []IntEntry{
		{1, 1},
		{3, 3},
		{5, 5},
		{6, 6},
		{8, 8},
	}
	for _, entry := range listToTest {
		heapInstance.Push(entry)
	}
	if entry, _ := heapInstance.Peek(); entry.Cmp(IntEntry{1, 1}) != 0 {
		t.Errorf("peek did not return the root entry, instead returned %d", entry.GetKey())
	}
}

func TestHeapqueue_PeekRandom(t *testing.T) {
	rand.Seed(seed)
	heapInstance := SetupHeapQueue()
	listToTest, _ := getRandomList(t)
	for _, entry := range listToTest {
		heapInstance.Push(entry)
	}
	for _, entry := range sortIntEntries(t, listToTest) {
		if peekedEntry, _ := heapInstance.Peek(); peekedEntry.Cmp(entry) != 0 {
			t.Errorf("Not popped in correct order")
			t.FailNow()
		}
		heapInstance.Pop()
	}
}

func sortIntEntries(t *testing.T, entryArray []IntEntry) []IntEntry {
	t.Helper()
	sort.SliceStable(entryArray, func(i, j int) bool {
		return entryArray[i].Cmp(entryArray[j]) == -1
	})
	var sortedIntEntries []IntEntry

	for _, entry := range entryArray {
		sortedIntEntries = append(sortedIntEntries, entry)
	}
	return sortedIntEntries
}

func getRandomList(t *testing.T) ([]IntEntry, int) {
	t.Helper()
	max := int(math.Pow(10, 3))
	min := int(math.Pow(10, 2))
	listSize := rand.Intn(max-min+1) + min
	var listToTest []IntEntry

	smallestInt := 1001
	for i := 0; i < listSize; i++ {
		num := rand.Intn(1000)
		listToTest = append(listToTest, IntEntry{num, num})
		if num < smallestInt {
			smallestInt = num
		}
	}
	return listToTest, smallestInt
}

func checkHeap(t *testing.T, heapInstance HeapQueue) {
	t.Helper()
	heapArray := heapInstance.(*heapqueue).array
	size := len(heapArray)
	for i, entry := range heapArray[:size/2] {
		if 2*i+1 < size {
			if !(entry.Cmp(heapArray[2*i+1]) <= 0) {
				t.Errorf("heap invariant does not hold. parent: %d, child: %d", i, 2*i+1)
				t.FailNow()
			}
		}
		if 2*i+2 < size {
			if !(entry.Cmp(heapArray[2*i+2]) <= 0) {
				t.Errorf("heap invariant does not hold. parent: %d, child: %d", i, 2*i+2)
				t.FailNow()
			}
		}
	}
}
