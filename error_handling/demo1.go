package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//nilden farklı olma durumu hata var demektir.
	if err != nil {
		fmt.Println("Dosya bulunamadı.")

	} else {
		fmt.Println(f.Name())
	}

}
