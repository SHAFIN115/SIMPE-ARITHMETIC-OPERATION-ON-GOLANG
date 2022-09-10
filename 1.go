package main

import (
	"fmt"
)

func add_values(a int, b int) int {

	var c = a + b

	return c
}

func add_sub(a int, b int) int {

	var c = a - b

	return c
}

func mul(a int, b int) int {

	var c = a * b

	return c
}

func div(a int, b int) int {

	var c = a / b

	return c
}

func main() {
	var i int
	var a int
	var b int
	var c int

	fmt.Print("ENTER WHAT TYPE OF OPERATION DO YOU WANT TO PERFORM?")
	fmt.Print("\n 1.ADD\n 2.SUBTRACT\n 3.MULTIPLY \n 4.DIVISION \n")

	fmt.Scan(&i)
	if i == 1 {
		fmt.Print("PUT TWO INTEGER :")
		fmt.Scan(&a)
		fmt.Scan(&b)
		c = add_values(a, b)
		fmt.Printf("Sum: %d \n", c)

	}
	if i == 2 {
		fmt.Print("PUT TWO INTEGER :")
		fmt.Scan(&a)
		fmt.Scan(&b)
		c = add_sub(a, b)
		fmt.Printf("SUBTRACTION: %d \n", c)

	}
	if i == 3 {
		fmt.Print("PUT TWO INTEGER :")
		fmt.Scan(&a)
		fmt.Scan(&b)
		c = mul(a, b)
		fmt.Printf("MULTIPLICATION: %d \n", c)

	}
	if i == 4 {
		fmt.Print("PUT TWO INTEGER :")
		fmt.Scan(&a)
		fmt.Scan(&b)
		c = div(a, b)
		fmt.Printf("DIVISION: %d \n", c)

	}

}
