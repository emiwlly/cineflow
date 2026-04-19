package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type movie struct {
	id int 
	name string
	quantity int 
}

const space = "================================"
var scanner = bufio.NewScanner(os.Stdin)

var movies []movie
func main () {

	movies = []movie{
		{id: 1, name: "The Crow", quantity: 5},
		{id: 2, name: "Wuthering Heights", quantity: 3},
		{id: 3, name: "Before sunrise", quantity: 0},
	} 

	fmt.Print(space)
	get_movies(movies)

	fmt.Println(space)
	movies = add_movies(movies)

	fmt.Println(space)
	movies = delete_movies(movies)

	fmt.Println(space)
	get_movies(movies)

}

func add_movies(m []movie) []movie {
	id := 4

	fmt.Print("which movie do you want to add (name): ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("quntity available (number): ")
	scanner.Scan()
	qtdStr := scanner.Text()

	quantity, err := strconv.Atoi(qtdStr)
	if err != nil {
		fmt.Print("Something went wrong")
		return m
	}

	m = append(m, movie{id, name, quantity})
	
	return m
}

func get_movies(m []movie) {
	for _, v := range m {
		fmt.Println("id:", v.id, "\nname:", v.name, "\nquantity available for rent:", v.quantity, "\n")
	}
}

func delete_movies(m []movie) []movie {
	fmt.Print("Which movie do you want to delete (id): ")
	scanner.Scan()
	id_to_remove := scanner.Text()

	id, err := strconv.Atoi(id_to_remove)
	if err != nil {
		fmt.Print("Something went wrong")
		return m
	}

	index := -1
	for i, v := range m {
		if v.id == id {
			index = i
			break
		}
	}

	if index != -1 {
		m = append(m[:index], m[index+1:]...)
	} else {
		fmt.Println("movie not found")
	}

	return m
}
