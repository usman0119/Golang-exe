package main

import (
	"fmt"
	"math"
)

func hitung(d float64) (float64, float64) {
	var luas = math.Pi * math.Pow(d/2, 2)

	var keliling = math.Pi * d

	return luas, keliling
}

func main() {

	var diameter float64 = 15

	var luas, keliling = hitung(diameter)

	fmt.Println(" hasil luas nya : ", luas)
	fmt.Println("hasil kelilingnya : ", keliling)
}
