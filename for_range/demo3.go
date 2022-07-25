package for_range

import "fmt"

func Demo3() {

	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for key, value := range sozluk {
		fmt.Println("key değerleri: ", key)
		fmt.Println("value değerleri: ", value)
	}

}
