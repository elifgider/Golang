package loops

import "fmt"

func Demo1() {

	i := 1
	for i < 11 {
		fmt.Printf(fmt.Sprintf("%v", i) + " hello world \n")
		i++
	}
}
