package pointers

import (
	"fmt"
)

func PrintPointer() {
	// Operador & é usado para obter o endereço de memória de uma variável
	var x int = 10

	// p agora é um pointeiro para x
	p := &x

	// imprime o endereço de memória de x
	fmt.Printf("endereço de memória de x: %p\n", p)

	// Operador *, usado com ponteiros serve para "Desreferenciação" que é usado para acessar o valor armazenado
	// no endereço de memória a qual esse ponteiro aponta

	fmt.Printf("valor de x antes de ser alterado pelo ponteiro: %d\n", x)

	// p ainda é um ponteiro, dessa forma, não estou alterando o endereço de memória e sim o valor armazenado no mesmo
	*p = 100

	fmt.Printf("valor de x alterado pelo ponteiro p: %d\n", x)

	// é possível criar um ponteiro que não aponta para nenhum lugar, ele irá possui o valor nil
	// tentar desreferenciar um ponteiro nil causa erro
	// *int = significa que o tipo é um ponteiro para um int
	var p1 *int

	// atribuindo valor ao ponteiro
	p1 = &x

	fmt.Printf("endereço de memória para de p1 (mesmo que x) %p", p1)
}
