package stack

import "fmt"

type Stack struct {
	itens []int
}

func Push(stack *Stack, item int) {
	newItens := make([]int, Size(stack)+1)
	newItens[len(newItens)-1] = item
	copy(stack.itens, newItens)
}

func Pop(stack *Stack) int {
	stack.itens = stack.itens[:len(stack.itens)-2]
	return stack.itens[Size(stack)-1]
}

func Top(stack *Stack) int {
	return stack.itens[Size(stack)-1]
}

func IsEmpty(stack *Stack) bool {
	return Size(stack) == 0
}

func Size(stack *Stack) int {
	return len(stack.itens)
}

func ExecuteAll() {
	var test Stack

	Push(&test, 1)
	Push(&test, 2)
	Push(&test, 3)

	fmt.Println(test)
}
