package main

import (
	"fmt"
)

func main() {
	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Patricia"
	fmt.Println(ptr, *ptr)

}
