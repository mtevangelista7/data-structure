package stack

import "fmt"

type Stack struct {
	itens []int
}

func Push(stack *Stack, item int) {
	newItens := make([]int, Size(stack)+1)
	copy(newItens, stack.itens)

	newItens[len(newItens)-1] = item
	stack.itens = newItens
}

func Pop(stack *Stack) int {
	top := Top(stack)
	stack.itens = stack.itens[0 : len(stack.itens)-1]
	return top
}

func Top(stack *Stack) int {
	return stack.itens[len(stack.itens)-1]
}

func IsEmpty(stack *Stack) bool {
	return Size(stack) == 0
}

func Size(stack *Stack) int {
	return len(stack.itens)
}

func ExecuteAll() {
	var test Stack

	Push(&test, 100)
	Push(&test, 110)
	Push(&test, 115)

	fmt.Println(test)

	fmt.Println(Pop(&test))
	fmt.Println(Pop(&test))
	Push(&test, 200)
}
