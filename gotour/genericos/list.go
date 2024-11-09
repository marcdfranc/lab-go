package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
// se precisar comparar inves de Any use comparablec como o que segue.
// observe o original em https://go-tour-br.appspot.com/tour/generics/2
type List[T comparable] struct {
	next *List[T]
	val  T
}

// Add adds a new element to the end of the list.
func (l *List[T]) Add(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
}

// Remove removes the first occurrence of the value from the list.
func (l *List[T]) Remove(value T) {
	current := l
	for current.next != nil {
		if current.next.val == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// IsEmpty checks if the list is empty.
func (l *List[T]) IsEmpty() bool {
	return l.next == nil
}

// Length returns the length of the list.
func (l *List[T]) Length() int {
	length := 0
	current := l
	for current.next != nil {
		length++
		current = current.next
	}
	return length
}

// Print prints all elements in the list.
func (l *List[T]) Print() {
	current := l.next
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func main() {
	list := &List[int]{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()
	fmt.Printf("Length: %d\n", list.Length())

	list.Remove(2)
	list.Print()
	fmt.Printf("Length: %d\n", list.Length())

	fmt.Printf("Is empty: %v\n", list.IsEmpty())
	list.Remove(1)
	list.Remove(3)
	fmt.Printf("Is empty: %v\n", list.IsEmpty())
}
