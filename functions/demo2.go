package functions

func DortIslem(sayi1, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	fark := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)
	return toplam, fark, carpim, bolum
}
