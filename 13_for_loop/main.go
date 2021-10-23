package main

import "fmt"

func main() {
	count := 1

	for count <= 10 {
		fmt.Println("Count ke ", count)
		count++
	}

	for count := 1; count < 10; count++ {
		fmt.Println("Count ke ", count)
	}

	buah := make(map[string]string)
	buah["nama"] = "Jeruk"
	buah["jenis"] = "Purut"

	for key, value := range buah {
		fmt.Println(key, " = ", value)
	}
}
