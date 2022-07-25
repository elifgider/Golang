package functions

import "fmt"

func Topla(a, b int) {
	var toplam = a + b
	fmt.Println(toplam)

}

func Cikar(a, b int) {

	var fark = a - b
	fmt.Println(fark)

}

func Carp(a, b int) {
	var carpim = a * b
	fmt.Println(carpim)
}

func Bol(a, b int) {
	var bolum = a / b
	fmt.Println(bolum)
}
