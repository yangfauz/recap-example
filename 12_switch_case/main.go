package main

import "fmt"

func main() {
	buah := "Jeruk"
	jum_buah := 3

	switch buah {
	case "Jeruk":
		fmt.Println("Ini Buah Jeruk")
	case "Jambu":
		fmt.Println("Ini Buah Jambu")
	default:
		fmt.Println("Bukan Buah Jeruk dan Bukan Buah Jambu")
	}

	switch jum_buah > 5 {
	case true:
		fmt.Println("Jumlah lebih dari 5")
	case false:
		fmt.Println("Jumlah kurang dari 5")
	}

	switch {
	case jum_buah > 5:
		fmt.Println("Jumlah lebih dari 5")
	case jum_buah < 5:
		fmt.Println("Jumlah kurang dari 5")
	default:
		fmt.Println("Jumlah 5")
	}
}
