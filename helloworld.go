package main

import "fmt"

func main() {

	type User struct {
		nome  string
		idade string
	}

	var pessoa User = User{
		nome:  "Douglas Goncalves de Souza",
		idade: "56",
	}

	const msg string = "Hello Word. %s\n."
	// var total int = 1
	// idade := 56

	/*
			  soma := func(x int, y int) int {
					return x + y
				}(1, 2)

		soma := func(x int, y int) int {
			return x + y
		}
	*/

	fmt.Printf(msg, pessoa.nome+", "+pessoa.idade)
}
