package main

import (
	"fmt"
	"container/list"
)

/*
1. Initialize two pointers (tortoise and hare) that both point to the head of the linked list
2. Loop as long as the hare does not reach null
3. Set tortoise to next node
4. Set hare to next, next node
5. If they are at the same node, reset the tortoise back to the head.
6. Have both tortoise and hare both move one node at a time until they meet again
7. Return the node in which they meet
8. Else, if the hare reaches null, then return null
*/

func hasCycle(head *list.Element) bool {
	if head == nil {
		return false
	}
	t, h := head, head

	for true {
		t = t.Next()
		if h.Next() != nil && h.Next().Next() != nil {
			h = h.Next().Next()
		}else{
			return false
		}
		if t == nil || h == nil {
			return false
		}
		if t == h {
			return true
		}
		fmt.Println(t,h)
	}
	return false
}

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	l.PushBack(4)
	l.PushBack(3)
	// l.PushBack(2)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(hasCycle(l.Front()))
}