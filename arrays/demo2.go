package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "Adana"
	sehirler[2] = "İzmir"
	sehirler[3] = "Bursa"
	sehirler[4] = "Denizli"

	fmt.Println(sehirler)

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
}
