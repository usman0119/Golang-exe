package main

import "fmt"

func main() {

	nomor := []int{1, 2, 3}

	hasil := 0

	for i := 0; i < len(nomor); i++ {
		hasil = hasil + nomor[i]
	}

	fmt.Println(hasil)

}
