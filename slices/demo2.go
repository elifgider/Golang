package slices

import "fmt"

func Demo2() {

	sehirler := []string{"Ankara", "İzmir", "Bursa", "Adana"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	copy(sehirlerKopya, sehirler)

	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirler[:])
}
