package interfaces

import "fmt"

type krediHesaplayici interface {
	kredi() float64
}

type konut struct {
	fiyat float64
	oran  float64
}

type araba struct {
	km   float64
	oran float64
}

type ihtiyac struct {
	gelir float64
	oran  float64
}

func (k1 konut) kredi() float64 {
	return k1.fiyat * k1.oran / 100 / 12
}

func (k2 araba) kredi() float64 {
	return k2.km * k2.oran / 100 / 12
}
func (k3 ihtiyac) kredi() float64 {
	return k3.gelir * k3.oran / 100 / 12
}

func AylikOdemeHesaplama(krediler []krediHesaplayici) float64 {
	odemeToplami := 0.0
	for i := 0; i < len(krediler); i++ {
		odemeToplami += float64(krediler[i].kredi())
	}
	return odemeToplami
}

func Demo2() {
	k1 := konut{fiyat: 200000, oran: 10}
	k2 := araba{km: 50000, oran: 15}
	k3 := ihtiyac{gelir: 5000, oran: 5}

	krediler := []krediHesaplayici{k1, k2, k3}
	odemeToplami := AylikOdemeHesaplama(krediler)
	fmt.Println("Toplam ödeme: ", odemeToplami)

}
