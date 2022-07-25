package loops

import "fmt"

//İki sayı birbirinin kendisi hariç bölenleri toplamına eşitse bu sayılara arkadaş sayılar denir.
//Girilen sayının arkadaş sayı olup olmadığını arkadaş sayı ise çiftini bulan program yazınız.
func Demo5() {

	sayi1 := 0

	fmt.Println(" sayı giriniz: ")
	fmt.Scanln(&sayi1)

	toplam := 0
	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam += i
		}

	}
	toplam2 := 0
	for i := 1; i < toplam; i++ {
		if toplam%i == 0 {
			toplam2 += i
		}

	}
	if sayi1 == toplam2 {
		fmt.Printf(" %v Arkadaş sayısı %v dır.", sayi1, toplam)

	} else {
		fmt.Println("Arkadaş sayı değildir.")
	}

}
