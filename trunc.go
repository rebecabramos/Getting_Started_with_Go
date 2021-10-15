/* The program truncates a number of type float to integer*/

package main

import "fmt"

func main() {
	var number float32

	fmt.Printf("Enter a float number, please: ")
	fmt.Scan(&number)

	fmt.Printf("The floating point number (%f) corresponds to the integer number(%d).", number, int(number))
}
