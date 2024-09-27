package main

import "fmt"

func main() {

	fmt.Println("\n###############")
	fmt.Println("EQUATION SOLVER")
	fmt.Println("###############")

	fmt.Println("\nWelcome to the equation solver.\nWhat type of equation do you want to solve?\n\t1. Linear\n\t2. Quadratic")

	for {
		fmt.Print("\nEnter selection: ")

		var selection int

		fmt.Scan(&selection)

		switch selection {
		case 1:
			solveLinear()
		case 2:
			solveQuadratic()
		default:
			fmt.Println("\nInvalid selection")
			continue
		}

		break
	}

}
