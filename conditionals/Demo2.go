package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenilen float64 = 900

	if cekilmekIstenilen > hesap {
		fmt.Println("Yetersiz Bakiye")
	} else if cekilmekIstenilen == hesap {
		fmt.Printf("hesapta para kalmadı : %v", hesap-cekilmekIstenilen)
	} else {
		fmt.Printf("kalan tutar: %v", hesap-cekilmekIstenilen)
	}

}
