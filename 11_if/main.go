package main

import "fmt"

func main() {
	buah := "Jeruk"

	if buah == "Jeruk" {
		fmt.Println("Ini Buah Jeruk")
	} else if buah == "Jambu" {
		fmt.Println("Ini Buah Jambu")
	} else {
		fmt.Println("Bukan Buah Jeruk dan Bukan Buah Jambu")
	}
}
