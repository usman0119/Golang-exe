package main

import (
	"fmt"
	"strings"
)

func validateHp(nohp string) string {
	if strings.HasPrefix(nohp, "08") {
		return "62" + nohp[1:]
	}
	return nohp
}

func main() {

	hasil := validateHp("081210460854")
	fmt.Println(hasil)

}
