package error_handling

import (
	"errors"
	"fmt"
)

//kullanıcıdan 1-100 aralığında sayı girmesini iste
//o sayıyı tahmin ederken aşağı yukarı diye bizi yönlendir

func TahminEt(tahmin int) (string, error) {
	sayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında sayı giriniz.")
	}

	if tahmin < sayi {
		return "yukarı ", nil

	}
	if tahmin > sayi {
		return "aşağı", nil
	}
	return "Bildiniz", nil
}

func Demo3() {
	mesaj, _ := TahminEt(50)
	fmt.Println(mesaj)

}
