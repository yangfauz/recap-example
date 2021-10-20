package main

import "fmt"

func main() {
	//Declaration Variables
	a := ""
	var b string
	var c string = ""

	fmt.Println(a, b, c)

	//Declaration Multivariable
	var d, e, f int
	var g, h, i = true, 2.3, "four"

	fmt.Println(d, e, f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	//Shorthand Variables
	j := 100
	var k float64 = 100
	var l []string
	var err error

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(err)
}
