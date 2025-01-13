package main

import (
	"fmt"
)

func reverse(k string) string {

	runes := []rune(k)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	//string to reverse

	kalimat := "KASUR INI RUSAK"

	pembalik := reverse(kalimat)

	if pembalik == kalimat {
		fmt.Println("true :  ", pembalik)
	} else {
		fmt.Println("false")
	}

}
