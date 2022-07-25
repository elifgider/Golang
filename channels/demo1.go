package channels

func CiftSayilar(CiftSayiCn chan int) {
	toplam := 0
	for i := 0; i < 10; i += 2 {
		toplam += i
	}
	CiftSayiCn <- toplam
}

func TekSayilar(TekSayilarCn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {
		toplam += i

	}
	TekSayilarCn <- toplam
}
