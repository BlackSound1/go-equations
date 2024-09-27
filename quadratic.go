package main

import (
	"fmt"
	"math"
)

func solveQuadratic() {
	fmt.Println("\n----------------")
	fmt.Println("Quadratic Solver")
	fmt.Println("----------------")

	fmt.Println("\nWhat form does the quadratic function take?\n\t1. Standard Form [y = ax^2 + bx + c]\n\t2. Vertex Form [y = a(x - h)^2 + k]")

	for {
		fmt.Print("\nEnter selection: ")
		var formOfFunction int

		fmt.Scan(&formOfFunction)

		switch formOfFunction {
		case 1:
			// Solve standard
			solveStandard()
		case 2:
			// Solve vertex
			solveVertex()
		default:
			fmt.Println("\nInvalid selection")
			continue
		}

		break

	}

}

func solveVertex() {
	fmt.Println("\n~~~~~~~~~~~~~~~~~~")
	fmt.Println("y = a(x - h)^2 + k")
	fmt.Println("~~~~~~~~~~~~~~~~~~")

	var a, h, k float64

	fmt.Print("\nWhat is the 'a' parameter? ")
	fmt.Scan(&a)

	fmt.Print("\nWhat is the 'h' parameter? ")
	fmt.Scan(&h)

	fmt.Print("\nWhat is the 'k' parameter? ")
	fmt.Scan(&k)

	fmt.Printf("\nYour equation is: y = %v(x - %v)^2 + %v\n", a, h, k)

	delta := -k / a

	posOrNeg(a)
	vertexVF(h, k)
	horizTrans(h)
	vertTrans(k)
	rootsVF(a, h, k, delta)
}

func solveStandard() {
	fmt.Println("\n~~~~~~~~~~~~~~~~~")
	fmt.Println("y = ax^2 + bx + c")
	fmt.Println("~~~~~~~~~~~~~~~~~")

	var a, b, c float64

	fmt.Print("\nWhat is the 'a' parameter? ")
	fmt.Scan(&a)

	fmt.Print("\nWhat is the 'b' parameter? ")
	fmt.Scan(&b)

	fmt.Print("\nWhat is the 'c' parameter? ")
	fmt.Scan(&c)

	fmt.Printf("\nYour equation is: y = %vx^2 + %vx + %v\n", a, b, c)

	delta := descriminant(a, b, c)
	posOrNeg(a)
	yIntercept(c)
	rootsSF(a, b, delta)
}

func descriminant(a float64, b float64, c float64) float64 {
	delta := math.Pow(b, 2) - 4*a*c

	var message string = ""

	if delta < 0 {
		message = "there are no real roots"
	} else if delta == 0 {
		message = "there is exactly 1 real root"
	} else {
		message = "there are exactly 2 real roots"
	}

	fmt.Printf("\nThe descriminant is %.2f, therefore %v\n", delta, message)

	return delta
}

func posOrNeg(a float64) {
	if a < 0 {
		fmt.Println("\nThe parabola opens down and has a maximum")
	} else if a > 0 {
		fmt.Println("\nThe parabola opens up and has a minimum")
	} else {
		fmt.Printf("\nBecause the 'a' paramter is %v, the function is linear, not parabolic", a)
	}
}

func yIntercept(c float64) {
	fmt.Printf("\nThe y-intercept is %v\n", c)
}

func rootsSF(a float64, b float64, delta float64) {
	radical := math.Sqrt(delta)
	denom := 2 * a

	if delta < 0 {
		fmt.Println("\nBoth roots are complex. Will not solve here")
	} else if delta == 0 {
		root := (-b + radical) / denom
		fmt.Printf("\nThe real root is %v, the other is complex (won't solve here)\n", root)
	} else {
		root1 := (-b + radical) / denom
		root2 := (-b - radical) / denom
		fmt.Printf("\nThe real roots are %.2f and %.2f\n", root1, root2)
	}
}

func vertexVF(h float64, k float64) {
	fmt.Printf("\nThe vertex is: (%v, %v)\n", h, k)
}

func horizTrans(h float64) {
	if h > 0 {
		fmt.Printf("\nThe horizontal translation is %v units right\n", h)
	} else if h < 0 {
		fmt.Printf("\nThe horizontal translation is %v units left\n", h)
	} else {
		fmt.Println("\nThere is no horizontal translation")
	}
}

func vertTrans(k float64) {
	if k > 0 {
		fmt.Printf("\nThe vertical translation is %v units up\n", k)
	} else if k < 0 {
		fmt.Printf("\nThe vertical translation is %v units down\n", k)
	} else {
		fmt.Println("\nThere is no vertical translation")
	}
}

func rootsVF(a float64, h float64, k float64, delta float64) {
	radical := math.Sqrt(-k / a)

	if delta < 0 {
		fmt.Println("\nBoth roots are complex. Will not solve here")
	} else if delta == 0 {
		root := h + radical
		fmt.Printf("\nThe real root is %v, the other is complex (won't solve here)\n", root)
	} else {
		root1 := h + radical
		root2 := h - radical
		fmt.Printf("\nThe real roots are %.2f and %.2f\n", root1, root2)
	}
}
