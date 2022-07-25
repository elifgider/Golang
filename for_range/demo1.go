package for_range

import "fmt"

func Demo1() {

	sehirler := []string{"Ankara", "İzmir", "İstanbul"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}
}
