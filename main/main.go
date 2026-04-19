package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type movie struct {
	name string
	quantity int 
}

const space = "================================"

var movies []movie
func main () {

	movies = []movie{
		{name: "The Crow", quantity: 5},
		{name: "Wuthering Heights", quantity: 3},
		{name: "Before sunrise", quantity: 0},
	} 

	fmt.Print(space)
	get_movies(movies)
	fmt.Println(space)
	movies = add_movies(movies)
	fmt.Println(space)
	get_movies(movies)

}

func add_movies(m []movie) []movie {
	scanner := bufio.NewScanner(os.Stdin)
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

	m = append(m, movie{name, quantity})
	
	return m
}

func get_movies(m []movie) {
	for _, value := range m {
		fmt.Println("name:", value.name, "\nquantity available for rent:", value.quantity, "\n")
	}
}
