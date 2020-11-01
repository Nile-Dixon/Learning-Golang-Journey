package main;

import (
	"fmt"
)

//A STRUCT THAT STORES INFORMATION OF STACK VALUE
type Node struct {
	Prev *Node
	Value string
}

type Stack struct {
	Top *Node
	Size int
}

//FACTORY FUNCTION TO CREATE STACK STRUCTS
func makeStack() Stack {
	stack := Stack{nil, 0}
	return stack
}

//FUNCTION TO ADD A NODE TO TOP OF STACK
func addToStack(s *Stack, n *Node) {
	stack := s
	if stack.Top == nil { 
		stack.Top = n
		stack.Size = 1
	} else {
		currentTop := stack.Top
		stack.Top = n
		n.Prev = currentTop
		stack.Size = stack.Size + 1
	}
}

//FUNCTION TO REMOVE THE TOP NODE FROM A STACK
func popFromStack(s *Stack) (*Node, error) {
	topNode := s.Top
	if topNode == nil {
		return nil, errors.New("Stack is empty.")
	}
	nextNode := topNode.Prev
	s.Top = nextNode
	s.Size = s.Size - 1
	return topNode, nil
}

func main() {
	fmt.Println("This is a test of my Stack")

	//CREATE THE STACK
	stack := makeStack()

	//CREATE NODES FOR STACK
	firstNode := Node{nil, "This is the first element of the stack"}
	secondNode := Node{nil, "This is the second element of the stack"}
	thirdNode := Node{nil, "This is the third element of the stack"}

	//ADD NODES TO THE STACK
	addToStack(&stack, &firstNode)
	addToStack(&stack, &secondNode)
	addToStack(&stack, &thirdNode)

	//REMOVE ELEMENT FROM STACK
	topNode, err := popFromStack(&stack)
	fmt.Println(topNode.Value)

	//REMOVE ANOTHER ELEMENT FROM STACK
	nextNode, err := popFromStack(&stack)
	fmt.Println(nextNode.Value)
}