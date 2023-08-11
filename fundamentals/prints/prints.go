package main

import "fmt"

func main() {

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("The value of X is", xs)
	fmt.Printf("The value of X is %f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa!"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
