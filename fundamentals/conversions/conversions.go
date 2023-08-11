package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	fmt.Println("Test concat " + strconv.Itoa(200))

	num, _ := strconv.Atoi("200")
	fmt.Println(num - 199)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Is bool")
	}

}
