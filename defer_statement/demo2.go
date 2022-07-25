package defer_statement

import "fmt"

func Kontrol(sayi int) string {

	if sayi%2 == 0 {
		return "Çift sayıdır"
	}
	if sayi%2 != 0 {
		return "tek sayıdır"
	}
	defer DeferFonksiyonu()
	return "belli değil"
}

func DeferFonksiyonu() {
	fmt.Println("Bitti")
}

func Test() {
	sonuc := Kontrol(5)
	fmt.Println(sonuc)
}
 