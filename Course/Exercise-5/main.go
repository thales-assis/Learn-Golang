/*
1 - at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
	The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

2 - in func main
	a - this should already be done.
		- print out the value of the variable “x”
		- print out the type of the variable “x”
		- assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
		- print out the value of the variable “x”
	b - now do this
		- now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
			i 	- then use the “=” operator to ASSIGN that value to “y”
			ii	- print out the value stored in “y”
			iii - print out the type of “y

*/

package main

import (
	"fmt"
)

type canary int

var x canary
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("x =", x)

	y = int(x)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", y)
}

/* ----- CONSOLE OUTPUT -----

0
main.canary
x = 42
y = 42
int

-------------------------- */
