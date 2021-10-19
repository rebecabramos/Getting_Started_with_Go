/* The program asks the user for a string of words, and the word they want to search and replace.
If the searched word exists in the list it is replace, otherwise an error message is displayed.*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	/* (s, old, new string, n int). Here, s is the original or given string, old is the
	string that you want to replace. new is the which replaces the old, and n is the
	number of times the old replaced */
	var string_init string
	var searched_string string
	var string_replace string
	var flag bool

	fmt.Printf("Enter a few words, when you think it's enough type stop> ")
	fmt.Scan(&string_init)

	var s []string
	s = append(s, string_init)

	for {
		fmt.Printf("Enter a few words, when you think it's enough type stop> ")
		fmt.Scan(&string_init)

		s = append(s, string_init)

		if string_init == "stop" {
			s = s[:len(s)-1] // remove the last item of list
			break
		}
	}

	fmt.Printf("Type what word are you looking for> ")
	fmt.Scan(&searched_string)

	fmt.Printf("Now, the syllable you want to replace in the previously typed word> ")
	fmt.Scan(&string_replace)

	for i := 0; i < len(s); i++ {
		if searched_string == s[i] {
			s[i] = strings.Replace(s[i], searched_string, string_replace, -1) // if the number integer n < 0 , which in this case is -1, there is no limit on the number of replacements
			flag = true
		}
	}

	if flag {
		fmt.Println("Result word list: ", s)
	} else {
		fmt.Println("Err: Could not find this word")
	}
}
