package string_functions

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "Elif"
	fmt.Println(strings.HasPrefix(isim, "El"))

	fmt.Println(strings.HasSuffix(isim, "if"))

	fmt.Println(strings.Index(isim, "li"))

	harfler := []string{"e", "l", "i", "f"}
	sonuc := strings.Join(harfler, "-")
	fmt.Println(sonuc)
	fmt.Println(strings.Replace(sonuc, "", "+", 2))
	fmt.Println(strings.Replace(isim, "", "*", -1))
	fmt.Println(strings.Split(isim, "*"))

}
