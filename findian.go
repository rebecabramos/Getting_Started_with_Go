/* The program checks if a word starts with (I) is concluded with (n) and has (a) in the middle */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Enter a word for verification> ")
	in := bufio.NewReader(os.Stdin) // use bufio.NewReader(os.Stdin) and in.ReadString('\n') for read string with spaces
	my_string, err := in.ReadString('\n')
	if err != nil {
		fmt.Errorf("some error happened - %v", err)
		return
	}

	// remove '\n'
	my_string = my_string[:len(my_string)-1]
	count := 0

	for position, char := range strings.ToLower(my_string) { //ToLower() function for transform all letter in lowercase
		if position == 0 && string(char) == "i" { // it is necessary to convert from rune to string to do the verification
			count++
		}
		if (position == len(my_string)-1) && (string(char) == "n") {
			count++
		}
		if strings.Contains(string(char), "a") {
			count++
		}
	}
	if count < 3 {
		fmt.Printf("Not Found!\n")
	} else {
		fmt.Printf("Found!\n")
	}
}
