package queue

type Queue struct {
	itens []int
}

func Enqueue(queue *Queue, value int) {
	size := Size(queue)
	queue.itens = queue.itens[:size+1]
	queue.itens[size] = value
}

func Dequeue(queue *Queue) int {
	primeiro := queue.itens[0]
	queue.itens = queue.itens[1:]
	return primeiro
}

func First(queue *Queue) int {
	return queue.itens[0]
}

func IsEmpty(queue *Queue) bool {
	return Size(queue) == 0
}

// usar o len() é melhor pois é O(1) e dessa forma fica O(n) visto que o slice possui a propriedade len
// mas resolvi fazer dessa só para não usar funções embutidas
func Size(queue *Queue) int {
	c := 0
	for range queue.itens {
		c++
	}
	return c
}
