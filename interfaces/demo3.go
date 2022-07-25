package interfaces

import "fmt"

func Demo3() {
	dogrula(5)
}
func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}
