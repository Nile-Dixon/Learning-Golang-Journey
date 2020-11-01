package main;

import (
	"fmt"
)

//A STRUCT THAT STORES INFORMATION OF STACK VALUE
type Node struct {
	Prev *Node
	Value string
}

//A STRUCT THAT STORE INFORMATION ABOUT THE QUEUE
type Queue struct {
	Back *Node
	Front *Node
	Size int
}

//FACTORY FUNCTION TO CREATE QUEUE STRUCTS
func makeQueue() Queue {
	queue := Queue{nil, nil, 0}
	return queue
}

//FUNCTION TO ADD TO QUEUE
func addToQueue(q *Queue, n *Node) {
	queue := q
	if queue.Size == 0 {
		queue.Back = n
		queue.Front = n
		queue.Size = queue.Size + 1
	} else if queue.Size == 1 {
		queue.Back = n
		queue.Front.Prev = n
		queue.Size = queue.Size + 1
	} else {
		queue.Back.Prev = n
		queue.Back = n
		queue.Size = queue.Size + 1
	}
}

//FUNCTION TO REMOVE FROM QUEUE
func removeFromQueue(q *Queue) *Node {
	queue := q
	if queue.Size == 0 {
		return nil
	} else if queue.Size == 1{
		returnNode := queue.Front
		queue.Front = nil
		queue.Back = nil
		queue.Size = 0
		return returnNode
	} else {
		returnNode := queue.Front
		queue.Front = returnNode.Prev
		queue.Size = queue.Size - 1
		return returnNode
	}
}

//FUNCTION TO TEST QUEUE
func main () {
	fmt.Println("This is a test of my Queue")

	queue := makeQueue()

	//CREATE NODES TO ADD TO QUEUE
	firstNode := Node{nil, "This is the first element of the stack"}
	secondNode := Node{nil, "This is the second element of the stack"}
	thirdNode := Node{nil, "This is the third element of the stack"}

	//ADD NODES TO THE QUEUE
	addToQueue(&queue,&firstNode)
	addToQueue(&queue,&secondNode)
	addToQueue(&queue,&thirdNode)

	//PRINT THE SIZE OF THE QUEUE
	fmt.Println("Size of the queue:",queue.Size)

	//REMOVE NODE FROM THE QUEUE
	returnNode := removeFromQueue(&queue)
	fmt.Println("Size of the queue:", queue.Size)
	fmt.Println("Contents of the node:", returnNode.Value)

}