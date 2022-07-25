package arrays

import "fmt"

func Demo4() {

	var sayilar [3][3]int
	sayilar[0][0] = 1
	sayilar[0][1] = 2
	sayilar[0][2] = 3
	sayilar[1][0] = 4
	sayilar[1][1] = 5
	sayilar[1][2] = 6
	sayilar[2][0] = 7
	sayilar[2][1] = 8
	sayilar[2][2] = 9

	for i := 0; i < len(sayilar); i++ {
		for j := 0; j < len(sayilar); j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
