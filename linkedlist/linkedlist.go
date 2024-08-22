package apagar

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func NewLinkedList(head *Node) *LinkedList {
	return &LinkedList{
		head: head,
	}
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func InsertBeg(list *LinkedList, value int) {
	node := NewNode(value)

	if list.head == nil {
		list.head = node
	} else {
		list.head.next = list.head
		list.head = node
	}
}

func InsertEnd(list *LinkedList, value int) {
	node := NewNode(value)

	if list.head == nil {
		list.head = node
	} else {
		actualNode := list.head

		for {
			if actualNode.next == nil {
				break
			}

			actualNode = actualNode.next
		}

		actualNode.next = node
	}
}

func Search(list *LinkedList, target int) bool {
	return true
}
