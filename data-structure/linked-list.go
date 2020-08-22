package ds

type Node struct {
	Data int
	Next *Node
}

func (list *Node) Len() (cnt int) {
	tmp := list
	for tmp.Next != nil {
		tmp = tmp.Next
		cnt++
	}
	return
}

func (list *Node) Append(val int) {
	n := &Node{val, nil}
	if list == nil {
		list = n
	}
	tmp := list
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = n
}

func (list *Node) Find(val int) bool {
	tmp := list
	for tmp != nil {
		if tmp.Data == val {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (list *Node) Remove(val int) {
	tmp := list
	var prev *Node
	for tmp != nil {
		if tmp.Data == val {
			if prev == nil { // first node
				list = tmp.Next
			} else {
				prev.Next = tmp.Next
			}
			tmp = nil // remove tmp
			break
		}
		prev = tmp
		tmp = tmp.Next
	}
}
