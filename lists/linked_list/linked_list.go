package linkedlist

import "fmt"

const MAX int = 50

type Record struct {
	key int
}

type Element struct {
	rec  Record
	prox int
}

type List struct {
	elements   [MAX]Element
	start      int
	availables int
}

func initList(list *List) {
	// aqui definimos o index do proximo elemento, para cada elemento contido na lista passada pelo usuario
	// dessa forma, todas as posições, menos o inicio e o elemento proximo do ultimo, possuem um proximo elemento
	for i := 0; i < MAX-1; i++ {
		list.elements[i].prox = i + 1
	}

	// aqui definimos que o ultimo elemento do array naõ possui um proximo
	list.elements[MAX-1].prox = -1 // -1 representa o final

	// o inicio não inicia com valor algum ele apenas pode apontar para o proximo elemento
	list.start = -1

	// como estamos iniciando a lista ainda não temos nenhum elemento disponivel
	list.availables = 0
}

func listSize(list *List) {
	// definimos uma variavel de inicio, representando o inicio de elementos válidos
	start := list.start

	// o tamanho inicia com 0
	size := 0

	// interamos até localizar um elemento onde o proximo seja inválido (-1)
	for start != -1 {
		// toda vez que entrar nesse loop é por que temos +1 elemento válido
		size++

		// busca pelo proximo elemento
		start = list.elements[start].prox
	}
}

func showList(list *List) {
	start := list.start

	// Para imprimir os valores das chaves, temos basicamente o mesmo procedimento que o tamanho
	// Temos que entrar no proximo elemento de cada elemento até que esse proximo seja inválido
	for start != -1 {
		fmt.Printf("Item: %d\n", list.elements[start].rec.key)
		start = list.elements[start].prox
	}
}

// atenção, isso é considerando uma linked list ordenada
func searchSequential(list *List, target int) int {
	i := list.start

	// aqui interamos sobre os elemento validos, que não são -1
	// e além disso, como a linked list está ordenada nesse exemplo
	// podemos verificar apenas as que são menores que o número buscado
	// pois caso seja maior, sabemos que não está na lista
	// exemplo: 1 => 5 => 7 => 10 => -1
	// então, enquanto a chave for menor, continuo indo para o proximo elemento
	for i != -1 && list.elements[i].rec.key < target {
		i = list.elements[i].prox
	}

	// após sair do loop, caso o elemento seja válido
	// e caso a chave desse elemento seja a chave buscada pelo usuario
	if i != 1 && list.elements[i].rec.key == target {
		return i
	}

	return -1
}
