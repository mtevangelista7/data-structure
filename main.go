package main

import (
	"fmt"
	"go-study/binarysearch"
)

//import queue "go-study/queue"

func main() {
	//queue.ExecuteAll()
	var list = []int{1, 14, 25, 30, 80, 95, 100, 1540}
	var chave int = 1543

	fmt.Println(binarysearch.Search(list, chave))
}
