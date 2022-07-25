package arrays

import "fmt"

//dizideki sayılardan en büyüğünü bulan program yazınız

func Demo3() {

	sayilar := [5]int{20, 30, 50, 10, 2}
	enBuyuk := sayilar[0]
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}

	}
	fmt.Println(enBuyuk)
}
