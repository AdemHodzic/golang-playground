package main

import "fmt"

func main() {

	fmt.Println("Enter the number of clomuns: ")
	var columns int
	fmt.Scanln(&columns)

	for column := 1; column <= columns; column++ {

		print(" ", columns - column)

		print("#", column)
		
		fmt.Printf(" ")

		print("#", column)

		fmt.Printf("\n")

	} 

}

func print(sign string, number int) {

	for i := 0; i < number; i++ {
		fmt.Printf(sign)
	}

}