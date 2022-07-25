package variables

import "fmt"

func Demo1() {
	var metin string = "merhaba Dunya!"
	fmt.Println(metin)

	var kdv float32 = 15
	fmt.Println("kdv oranı: ", kdv)

	var eskiFiyat float32 = 50
	var yeniFiyat float32 = eskiFiyat + (eskiFiyat * kdv / 100)
	fmt.Println("eskiFiyat: ", eskiFiyat)
	fmt.Println("yeniFiyat: ", yeniFiyat)

	kdv2 := 8.5
	fmt.Printf("kdv2 veri tipi: %T", kdv2)
	fmt.Print("\n")

	var metin1 string = "Elif"
	var metin2 string = "Esma"
	var sonuc bool = metin1 == metin2
	fmt.Println(sonuc)
}
