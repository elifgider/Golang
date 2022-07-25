package loops

import "fmt"

func Demo3() {
	// iki oyunculu sayı tahmin etme oyunu
	//ilk oyuncu aklından bir sayı tutar
	//ikinci oyuncunun  bu sayıyı tahmin etmek için 3 hakkı var.

	aklimdakiSayi := 0
	tahmin := 0

	fmt.Println("1.oyuncu aklından bir sayı tut.")
	fmt.Scanln(&aklimdakiSayi)
	fmt.Println("2.oyuncu tahmin et 3 hakkın var.")

	var tahminHakki int = 3
	for i := 0; i < tahminHakki; {
		fmt.Scanln(&tahmin)

		if tahmin == aklimdakiSayi {

			fmt.Print(i + 1)
			fmt.Println(" KEREDE BİLDİNİZ")
			if tahminHakki == 1 {
				fmt.Println("mükemmel")
			} else {
				fmt.Println("Tebrikler!")
			}

			break
		}
		if tahmin < aklimdakiSayi {

			if i == 2 {
				fmt.Println("Üzgünüm deneme hakkın bitti.")
			} else {
				fmt.Println("Bir daha dene")
				fmt.Println("yukarı")
			}
			i++

		}
		if tahmin > aklimdakiSayi {

			if i == 2 {
				fmt.Println("Üzgünüm deneme hakkın bitti.")
			} else {
				fmt.Println("aşağı")
				fmt.Println("Bir daha dene")
			}
			i++
		}

	}

}
