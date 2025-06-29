package main

import "fmt"

func main() {
	// Declaring an array using shorthand syntax
	names := [4]string{"apple", "banana", "cherry", "date"}

	// Displaying array elements using a loop
	fmt.Println("Array elements:")

	for index := 0; index < 3; index++ {
		fmt.Println(names[index])
	}

	for i, fruits := range names {
		fmt.Println(i,fruits)
	}
}