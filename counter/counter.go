package counter

import (
	"container/heap"
	"sync"
)

type counterItem struct {
	value interface{}
	// the number of times this item is in the counter.
	// since we're using a heap, this can be thought of as the priority
	count int
	// the index in the heap array
	index int
}

// A "max heap", where the largest item is at the "top" of the heap
type counterHeap []*counterItem

func (h counterHeap) Len() int {
	return len(h)
}

func (h counterHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, so we use greater than here
	return h[i].count > h[j].count
}

func (h counterHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *counterHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*counterItem)
	item.index = n
	*h = append(*h, item)
}

func (h *counterHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1
	*h = old[0 : n-1]
	return item
}

func (h *counterHeap) update(item *counterItem, value interface{}, count int) {
	item.value = value
	item.count = count
	heap.Fix(h, item.index)
}

type Counter struct {
	mapping map[interface{}]*counterItem
	heap    counterHeap
	mutex   sync.Mutex
}

func New() *Counter {
	return &Counter{
		mapping: map[interface{}]*counterItem{},
		heap:    make(counterHeap, 0),
	}
}

type CounterItem struct {
	Value interface{}
	Count int
}

// TODO: is it possible to get the most common without popping and pushing?
// O(n log k)
func (c *Counter) MostCommon(n int) (items []CounterItem) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var toPushBack []*counterItem

	for n > 0 && len(c.heap) > 0 {
		item := heap.Pop(&c.heap).(*counterItem)
		items = append(items, CounterItem{
			Value: item.value,
			Count: item.count,
		})
		toPushBack = append(toPushBack, item)
		n--
	}

	for _, item := range toPushBack {
		heap.Push(&c.heap, item)
	}

	return
}

func (c *Counter) Get(item interface{}) int {
	if counterItem, ok := c.mapping[item]; ok {
		return counterItem.count
	}
	return 0
}

func (c *Counter) AddItems(items ...interface{}) {
	for _, item := range items {
		c.addItem(item)
	}
}

func (c *Counter) addItem(item interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if cItem, ok := c.mapping[item]; ok {
		c.heap.update(cItem, cItem.value, cItem.count+1)
	} else {
		cItem := &counterItem{
			value: item,
			count: 1,
		}
		c.mapping[item] = cItem
		heap.Push(&c.heap, cItem)
	}
}

