package main

import (
	"fmt"
)

func solveLinear() {

	fmt.Println("\n-------------")
	fmt.Println("Linear Solver")
	fmt.Println("-------------")

	fmt.Println("\nWhat form does the linear function take?\n\t1. Slope-Intercept Form [y = mx + b]\n\t2. Point-Slope Form [y2 - y1 = m(x2 - x1)]")

	for {
		fmt.Print("\nEnter selection: ")
		var formOfFunction int

		fmt.Scan(&formOfFunction)

		switch formOfFunction {
		case 1:
			// Solve slope intercept
			solveSI()
		case 2:
			// Solve point slope
			solvePS()
		default:
			fmt.Println("\nInvalid selection")
			continue
		}

		break
	}
}

func solveSI() {
	fmt.Println("\n~~~~~~~~~~")
	fmt.Println("y = mx + b")
	fmt.Println("~~~~~~~~~~")

	var m, b float64

	fmt.Print("\nWhat is the 'm' parameter? ")

	fmt.Scan(&m)

	fmt.Print("\nWhat is the 'b' parameter? ")

	fmt.Scan(&b)

	fmt.Printf("\nYour equation is: y = %vx + %v\n", m, b)

	incOrDec(m)

	fmt.Printf("\nThe slope is %.2f\n", m)

	fmt.Printf("\nThe y-intercept is %.2f\n", b)

	rootsSI(m, b)
}

func solvePS() {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("y2 - y1 = m(x2 - x1)")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~")

	var x1, x2, y1, y2 float64

	fmt.Print("\nWhat is the 'x1' coordinate? ")

	fmt.Scan(&x1)

	fmt.Print("\nWhat is the 'y1' coordinate? ")

	fmt.Scan(&y1)

	fmt.Print("\nWhat is the 'x2' coordinate? ")

	fmt.Scan(&x2)

	fmt.Print("\nWhat is the 'y2' coordinate? ")

	fmt.Scan(&y2)

	m := (y2 - y1) / (x2 - x1)

	incOrDec(m)

	fmt.Printf("\nThe slope is %.2f\n", m)

	fmt.Printf("\nThe y-intercept is %.2f\n", yIntPS(x1, y1, m))

	fmt.Printf("\nThe root is %.2f\n", rootPS(x1, y1, m))
}

func incOrDec(m float64) {
	if m > 0 {
		fmt.Println("\nThe function is increasing")
	} else if m < 0 {
		fmt.Println("\nThe function is decreasing")
	} else {
		fmt.Println("\nThe function is constant")
	}
}

func yIntPS(x1 float64, y1 float64, m float64) float64 {
	return m*(0-x1) + y1
}

func rootPS(x1 float64, y1 float64, m float64) float64 {
	x := (1/m)*(0-y1) + x1

	return x
}

func rootsSI(m float64, b float64) {
	y := -b
	zero := y / m

	fmt.Printf("\nThe root is %.2f", zero)
}
