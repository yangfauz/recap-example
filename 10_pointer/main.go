package main

import "fmt"

func main() {
	//Pass by value (mengcopy)
	// buah := "Apel"
	// sayur := buah

	// sayur = "brokoli"

	// fmt.Println(buah)
	// fmt.Println(sayur)

	//Pass by Reference (mengcopy)
	buah := "Apel"
	sayur := &buah

	*sayur = "brokoli"

	fmt.Println(buah)
	fmt.Println(*sayur)
}
