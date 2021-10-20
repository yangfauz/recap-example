package main

import "fmt"

func main() {
	//integer
	var u uint8 = 255
	fmt.Println(u)
	var i int8 = 127
	fmt.Println(i)
	var a int32 = 10
	var b int32 = 20
	fmt.Println(a, b)

	//floating point
	var decimalNumber float64 = 2.62
	fmt.Printf("Decimal Number: %f\n", decimalNumber)   // 15 digit precision
	fmt.Printf("Decimal Number: %.3f\n", decimalNumber) // 3 angka dibelakang koma
}
