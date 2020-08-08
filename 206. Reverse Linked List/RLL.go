package main

type ListNode struct {
	val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {

// 	var previous *ListNode
// 	current := head
// 	for current != nil {
// 		nextTemp := current.Next
// 		current.Next = previous
// 		previous = current
// 		current = nextTemp
// 	}

// 	return previous

// }



func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	curr := head

	for ;curr!=nil; {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev

}
func main() {
	input := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	reverseList(&input)

}
