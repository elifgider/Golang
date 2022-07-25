package conditionals

import "fmt"

func Workshop1() {
	//Üç adet int değişkeni tanımlayınız.
	//Ekrana en büyük olanı yazdırınız.

	var degisken1 int = 10
	var degisken2 int = 5
	var degisken3 int = 15

	if degisken1 > degisken2 {

		if degisken1 > degisken3 {
			fmt.Println("EN BUYUK :", degisken1)
		}

	} else if degisken2 > degisken3 {
		fmt.Println("en büyük: ", degisken2)

	} else {
		fmt.Println("en buyuk:", degisken3)
	}

	// var sayi1,sayi2,sayi3 int=10,5,15
	// var enBuyuk int=sayi1
	//if sayi2>enBuyuk{
	//enBuyuk=sayi2}
	//if sayi3>enBuyuk {
	//enBuyuk=sayi3}
	//fmt.Printf("en büyük sayı: %v",enBuyuk)

}
