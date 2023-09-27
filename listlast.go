package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(link *List) interface{} {
	for link.Head != nil {
		if link.Head.Next == nil {
			return link.Head.Data
		}
		link.Head = link.Head.Next
	}

	return nil
}
