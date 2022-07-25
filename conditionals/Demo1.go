package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenilen float64 = 900

	if cekilmekIstenilen > hesap {
		fmt.Println("Yetersiz Bakiye")
	}

	if cekilmekIstenilen < hesap {
		hesap = hesap - cekilmekIstenilen
		fmt.Printf("kalan tutar: %v", hesap)
	}

}
