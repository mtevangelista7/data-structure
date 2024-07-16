package sequentiallinear

import "fmt"

// implementação de uma lista linear sequencial

// define a capacidade max
const max int = 10

// estrutura de exemplo para simular registros
type register struct {
	key int
}

// estrutura lista, possui um array e um número de elementos
type list struct {
	register         [max + 1]register
	numberOfElements int
}

// inicializa a lista
func InitList(l *list) {
	l.numberOfElements = 0
}

// retorna o tamanho da lista
func Size(l *list) int {
	return l.numberOfElements
}

// exibe toda a lista
func ShowList(l *list) {
	for _, register := range l.register {
		fmt.Println(register.key)
	}
}

// realiza uma busca dentro da lista
func SequencialSearch(l *list, key int) int {
	// seto que o último elemento válido do array é a chave buscada
	l.register[l.numberOfElements].key = key

	var i int

	// passo por cada chave do array
	// comparando a chave
	for l.register[i].key != key {
		i++
	}

	// verifico se cheguei ao final do array
	if i == l.numberOfElements {
		return -1
	}

	// caso tenha chego aqui significa que saiu do for antes de chegar no final do array
	// logo a chave foi localizada e não é a sentinela adicionada
	return i
}

// adiciona um novo elemento numa posição
func AddElement(l *list, register register, position int) bool {
	if l.numberOfElements == max || position < 0 || position > l.numberOfElements {
		return false
	}

	for i := l.numberOfElements; i > position; i-- {
		l.register[i] = l.register[i-1]
	}

	l.register[position] = register
	l.numberOfElements++
	return true
}

// remove um elemento da liista
func RemoveElement(l *list, element int) bool {

	// verifica se o elemento existe na lista
	position := SequencialSearch(l, element)

	// caso não retorna -1
	if position == -1 {
		return false
	}

	// a partir da posição do item que deverá ser removido
	// desloco o proximo item para a sua posição (uma posição para tras)
	// e assim faço para o resto da lista
	for i := position; i < l.numberOfElements; i++ {
		l.register[i] = l.register[i+1]
	}

	// removo 1 da quantidade
	l.numberOfElements--
	return true
}

func RestoreList(l *list) {
	l.numberOfElements = 0
}

// função para teste, visto que esses métodos não estão exatamente dentro da struct
func ExecuteAll() {
	var test list

	InitList(&test)

	// fmt.Printf("%d \n", Size(&test))

	// ShowList(&test)

	var r1 register = register{42}
	var r2 register = register{75}
	var r3 register = register{21}

	AddElement(&test, r1, 0)
	AddElement(&test, r2, 1)
	AddElement(&test, r3, 2)

	RemoveElement(&test, 42)

	ShowList(&test)
}
