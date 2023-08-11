package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var ray = 3.2

	area := PI * math.Pow(ray, 2)

	fmt.Println("The area of the circumference is", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Opa!"
	fmt.Println(g, h, i)


}
