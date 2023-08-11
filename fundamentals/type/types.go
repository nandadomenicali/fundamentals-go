package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal is:", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("Byte is:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("The value maximum is:", i1)

	var i2 rune = 'a'
	fmt.Println("The rune is:", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("The type of x is:", reflect.TypeOf(x))
	fmt.Println("The literal type is:", reflect.TypeOf(49.99))

	bo := true
	fmt.Println(!bo)

	s1 := "Hello World!"
	fmt.Println(s1 + "!")
	fmt.Println("The length of the string is:", len(s1))

	s2 := `Hello
	my 
	name
	is....`
	fmt.Println("The length of the string is:", len(s2))
}
