package main

import "fmt"

var sblm = 0
var sedh = 1
var hasil int


func main() {
	for j := 0; j < 10; j++{
		hasil = sblm + sedh
		sblm = sedh
		sedh = hasil
		
        fmt.Println(hasil)
    }
}