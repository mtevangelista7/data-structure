package queue

import "fmt"

type Queue struct {
	itens []int
}

// o append é muito melhor para esse caso
func Enqueue(queue *Queue, value int) {
	newSlice := make([]int, len(queue.itens)+1)
	copy(newSlice, queue.itens)
	newSlice[len(newSlice)-1] = value
	queue.itens = newSlice
}

func Dequeue(queue *Queue) int {
	first := queue.itens[0]
	queue.itens = queue.itens[1:]
	return first
}

func First(queue *Queue) int {
	return queue.itens[0]
}

func IsEmpty(queue *Queue) bool {
	return len(queue.itens) == 0
}

// usar o len() é melhor pois é O(1) e dessa forma fica O(n) visto que o slice possui a propriedade len
// mas resolvi fazer dessa só para realizar as operações principais sem função embutida
func Size(queue *Queue) int {
	c := 0
	for range queue.itens {
		c++
	}
	return c
}

func ExecuteAll() {
	var test Queue

	Enqueue(&test, 100)
	Enqueue(&test, 110)
	Enqueue(&test, 115)
	Dequeue(&test)
	Dequeue(&test)
	fmt.Println(test)
}
