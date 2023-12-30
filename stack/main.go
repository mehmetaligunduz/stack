package main

import (
	"fmt"
	"stack/memory"
)

func main() {
	newStack := memory.NewStack()
	newStack.Push("A")
	newStack.Push("B")
	newStack.Push("C")
	newStack.Push("D")

	newStack.ToString()

	fmt.Println("------------------------------------------")

	newStack.Pop()
	newStack.Pop()

	newStack.ToString()

	fmt.Println("------------------------------------------")

	newStack.Push("E")
	newStack.Push("F")

	newStack.ToString()

}
