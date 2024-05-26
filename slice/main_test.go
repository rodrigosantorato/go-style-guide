package main_test

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	/*
		1.
		nil vs empty slice
		refs:
		https://go.dev/wiki/CodeReviewComments#declaring-empty-slices
		https://github.com/uber-go/guide/blob/master/style.md#nil-is-a-valid-slice
	*/
	t.Run("slice nil vs slice vazio", func(t *testing.T) {
		// ruim
		emptySlice := []int{}
		fmt.Printf("empty slice: %v, type: %T, len: %v, cap: %v\n", emptySlice, emptySlice, len(emptySlice), cap(emptySlice))

		// bom
		var nilSlice []int
		fmt.Printf("nil slice: %v, type: %T, len: %v, cap: %v\n\n", nilSlice, nilSlice, len(nilSlice), cap(nilSlice))

		// os dois são quase idênticos, mas quando comparados com nil tem comportamento diferente
		fmt.Printf("empty slice == nil: %v\n", emptySlice == nil)
		fmt.Printf("nil slice == nil: %v\n\n", nilSlice == nil)

		// nil é um slice válido
		nilSlice = append(nilSlice, 1)
		fmt.Printf("vc pode dar append em um slice nil. nilSlice depois do append: %v\n\n", nilSlice)

		// prefira usar len(slice) para checar se o slice está vazio
		// ruim
		if nilSlice == nil {
			// lógica aqui
			// (nilSlice cairia aqui, mas o emptySlice não)
		}

		// bom
		if len(nilSlice) == 0 {
			// lógica aqui
			// (tanto nil quanto slice vazio vão ter length de 0)
		}
	})

	/*
		2.
		slice length and capacity
		refs:
		https://github.com/uber-go/guide/blob/master/style.md#specifying-slice-capacity
	*/
	t.Run("tamanho e capacidade de um slice", func(t *testing.T) {
		/*
			o slice funciona como uma referência para um array
			o tamanho (length) do slice reflete a quantidade de elementos desse slice
			a capacidade reflete o tamanho do array por trás desse slice

			ao inicializar um slice populado, seu tamanho e capacidade serão igual ao número de elementos
		*/
		slice := []int{1, 2}
		fmt.Printf("slice: %v\nlength: %v, capacity: %v\n", slice, len(slice), cap(slice))

		/*
			Quando tentamos adicionar um elemento além da capacidade máxima do slice,
			a linguagem fará uma alocação de um novo array em memória para que o slice consiga aumentar em tamanho
			Exp: length == 2 e cap == 2 então ao adicionar o numero 3 a capacidade precisa aumentar
		*/
		slice = append(slice, 3)
		fmt.Printf("slice: %v\nlength: %v, capacity: %v\n", slice, len(slice), cap(slice))

		// perceba que normalmente o novo array de apoio do slice vem como o dobro da capacidade anterior
		for num := 4; num < 13; num++ {
			slice = append(slice, num)
			fmt.Printf("slice: %v\nlength: %v, capacity: %v\n", slice, len(slice), cap(slice))
		}
		fmt.Println("-------------------")

		/*
			sempre que possível inicialize um slice com capacidade prealocada
			isso evita alocações em memória desnecessárias e faz diferença no consumo de memória da sua aplicação
		*/
		capacity := 10
		preAlloc := make([]int, 0, capacity)
		fmt.Printf("preAlloc: %v, length: %v, capacity: %v\n", preAlloc, len(preAlloc), cap(preAlloc))
		for i := 1; i <= capacity; i++ {
			preAlloc = append(preAlloc, i)
		}
		fmt.Printf("preAlloc: %v, length: %v, capacity: %v\n", preAlloc, len(preAlloc), cap(preAlloc))
	})

	/*
		3.
		slices are like references to arrays
		refs:
		https://go.dev/tour/moretypes/8
		https://github.com/uber-go/guide/blob/master/style.md#copy-slices-and-maps-at-boundaries
		https://github.com/uber-go/guide/blob/master/style.md#returning-slices-and-maps
	*/
	t.Run("lembre-se de que slices funcionam como ponteiros para arrays", func(t *testing.T) {
		games := []string{"castlevania", "final fantasy"}

		type person struct {
			games []string
		}
		var rodrigo person
		// vamos atribuir o slice games para a propriedade games de rodrigo
		rodrigo.games = games
		fmt.Printf("games :%v, rodrigo.games: %v\n\n", games, rodrigo.games)

		// alterar o slice games causará mudanças na propriedade rodrigo.games
		games[0] = "ELDEN RING"
		fmt.Printf("games :%v, rodrigo.games: %v\n\n", games, rodrigo.games)

		// vamos voltar games ao normal e tentar de novo, usando uma variavel temp
		games[0] = "castlevania"
		temp := games
		rodrigo.games = temp
		games[0] = "METAL GEAR SOLID"
		/*
			temos o mesmo resultado não desejado, pois games contém a referência de um array
			essa mesma referência é passada para temp
			essa mesma referência é passada para rodrigo.games
			quando o slice original games é alterado, todas as referências sofrem alteração

			por isso gosto de pensar que os slices tem comportamento de ponteiro
		*/
		fmt.Printf("games :%v, rodrigo.games: %v\n\n", games, rodrigo.games)

		games[0] = "castlevania"
		/*
			para fugir desse problema:
			ao atribuir o valor de um slice para uma propriedade é recomendado criar um novo slice
			com o tamanho do slice origem e então copiar o valor do slice origem para o destino
		*/
		rodrigo.games = make([]string, len(games))
		copy(rodrigo.games, games)
		games[0] = "WORLD OF WARCRAFT"
		fmt.Printf("games :%v, rodrigo.games: %v\n\n", games, rodrigo.games)
	})

	/*
		4.
		append returns a new slice
	*/
	t.Run("append retorna um novo slice", func(t *testing.T) {
		// vamos criar um slice que contem uma lista de palavras
		words := []string{"cat", "dog", "mouse"}
		// vamos criar um struct "person" que será uma pessoa que possui um vocabulário
		type person struct {
			vocab []string
		}

		var sara, karen person
		/*
			para economizar memória, vamos adicionar os elementos novos no slice words
			então vamos usar pedaços do slice words para atribuir os vocabulários de sara e karen

			sara recebe só até o quarto elemento "eagle", karen recebe todos os elementos (até whale)
		*/
		words = append(words, "eagle", "whale")
		sara.vocab = words[:4]
		karen.vocab = words
		fmt.Printf(
			"words: %v\nsara's vocab: %v\nkaren's vocab: %v\n\n",
			words,
			sara.vocab,
			karen.vocab,
		)

		/*
			quando alteramos o primeiro elemento de words, o vocabulário de sara e karen serão afetados
			isso pode ocasionar bugs escondidos em produção
		*/
		words[0] = "GOTCHA"
		fmt.Printf(
			"words: %v\nsara's vocab: %v\nkaren's vocab: %v\n\n",
			words,
			sara.vocab,
			karen.vocab,
		)

		// vamos voltar words ao normal e criar novas pessoas
		words[0] = "cat"
		var john, steve person
		/*
			queremos atribuir a lista de words + eagle para john
			queremos atribuir a lista de words + eagle + whale para steve
			o append retorna um novo slice, então o vocabulário de john e steve são slices diferentes de words
		*/
		john.vocab = append(words, "eagle")
		steve.vocab = append(words, "eagle", "whale")
		fmt.Printf(
			"words: %v\njohn's vocab: %v\nsteve's vocab: %v\n\n",
			words,
			john.vocab,
			steve.vocab,
		)

		// modificando o slice words, não afeta john e steve
		words[0] = "GOTCHA"
		fmt.Printf(
			"words: %v\njohn's vocab: %v\nsteve's vocab: %v\n\n",
			words,
			john.vocab,
			steve.vocab,
		)
	})
}
