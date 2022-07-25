package pointers

import "fmt"

func Demo2(sayilar []int) {
	sayilar[0] = 10

	fmt.Println("Demodaki  ilk elaman: ", sayilar[0])
}
