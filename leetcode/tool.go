package dance


type ListNode struct {
	Val  int
	Next *ListNode
}


func constructorList(list []int) *ListNode {
	var head, previous, node *ListNode
	for _, num := range list {
		node = &ListNode{
			Val:  num,
			Next: nil,
		}
		if head == nil {
			head = node
		}
		if previous != nil {
			previous.Next = node
		}
		previous = node
	}
	return head
}