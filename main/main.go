package main

import (
	"fmt"
)

type movie struct {
	name string
	quantity int 
}

var movies []movie
func main () {

	movies = []movie{
		{name: "The Crow", quantity: 5},
		{name: "Wuthering Heights", quantity: 3},
		{name: "Before sunrise", quantity: 0},
	} 

	get_movies(movies)

}

func get_movies(m []movie) {
	for _, value := range m {
		fmt.Println("name:", value.name, "\nquantity available for rent:", value.quantity, "\n")
	}
}
