package main

import (
	"fmt"
	"strings" //import string package
)

func main() {
	var test = "joe"
	first_name := "syed" //:= short decleration
	last_name := "aftab"
	full_name := first_name + " " + last_name

	fmt.Println(strings.ToUpper(test))
	fmt.Println(full_name)
	fmt.Println(strings.ToUpper(full_name))
}
