package linkedlist

type LinkedListNode struct {
	next *LinkedListNode
	prev *LinkedListNode

	value interface{}
}

func (n *LinkedListNode) Value() interface{} {
	return n.value
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func New() *LinkedList {
	list := &LinkedList{
		head: &LinkedListNode{},
		tail: &LinkedListNode{},
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (l *LinkedList) Append(value interface{}) *LinkedListNode {
	n := &LinkedListNode{
		prev:  l.tail.prev,
		next:  l.tail,
		value: value,
	}

	l.tail.prev.next = n
	l.tail.prev = n

	return n
}

func (l *LinkedList) Remove(n *LinkedListNode) bool {
	if n == l.head || n == l.tail {
		return false
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	return true
}

func (l *LinkedList) Iterate() chan *LinkedListNode {
	ch := make(chan *LinkedListNode)

	go func() {
		n := l.head

		for n.next != l.tail {
			ch <- n.next
			n = n.next
		}

		close(ch)
	}()

	return ch
}

