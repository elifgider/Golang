package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3)

	isimler[0] = "elif"
	isimler[1] = "enes"
	isimler[2] = "esma"
	fmt.Println(isimler)
	isimler = append(isimler, "hulya")
	fmt.Println(isimler)
}
