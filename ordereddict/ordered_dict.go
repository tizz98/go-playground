package ordereddict

import "github.com/tizz98/go-playground/linkedlist"

type OrderedDict struct {
	lookup map[string]*linkedlist.LinkedListNode
	list   *linkedlist.LinkedList
}

func New() *OrderedDict {
	return &OrderedDict{
		lookup: make(map[string]*linkedlist.LinkedListNode),
		list:   linkedlist.New(),
	}
}

func (d *OrderedDict) Set(key string, value interface{}) {
	if n, ok := d.lookup[key]; ok {
		d.list.Remove(n)
	}

	d.lookup[key] = d.list.Append(value)
}

func (d *OrderedDict) Get(key string) interface{} {
	return d.lookup[key].Value()
}

func (d *OrderedDict) Remove(key string) bool {
	if n, ok := d.lookup[key]; ok {
		if ok := d.list.Remove(n); !ok {
			return false
		}
		delete(d.lookup, key)
		return true
	}
	return false
}

func (d *OrderedDict) Iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		for v := range d.list.Iterate() {
			ch <- v.Value()
		}

		close(ch)
	}()

	return ch
}

