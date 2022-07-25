package loops

import "fmt"

//Kullanıcının girdiği sayının asal olup olmadığını kontrol eden program

func Demo4() {
	sayi := 0
	fmt.Println("Pozitif bir sayi giriniz: ")
	fmt.Scanln(&sayi)

	if sayi < 0 {
		fmt.Println("Lütfen pozitif bir sayı giriniz.")
	} else {
		asalMi := true
		for i := 2; i < sayi; i++ {
			if sayi%i == 0 {
				asalMi = false
			}
		}
		if asalMi == true {
			fmt.Println("Girdiginiz sayi asal")
		} else {
			fmt.Println("Girdiginiz sayi asal degil")
		}

	}
	if sayi == 2 {
		fmt.Println("Girdiginiz sayi asal")
	}
}
