package string_functions

import (
	"fmt"
	"strings"
)

func Demo1() {
	isim := "Elif"
	fmt.Println(strings.Count("hello", "l"))
	fmt.Println(strings.Contains("hello", "h"))
	fmt.Println(strings.Index(isim, "f"))
	fmt.Println(strings.ToLower(isim))
	fmt.Println(strings.ToUpper(isim))
}
