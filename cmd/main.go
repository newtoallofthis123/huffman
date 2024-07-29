package main

import (
	"fmt"

	"github.com/newtoallofthis123/huffman/tree"
)

func main() {
	fmt.Println("Hello World")

	chars := []interface{}{'A', 'A', 'A', 'B', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D', 'E', 'E'}

	t := tree.FromList(chars)
	t.DebugPrint()
}
