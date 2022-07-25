package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {

	for i := 0; i < 10; i += 2 {
		fmt.Println("çift sayilar", i)
		time.Sleep(time.Second)
	}

}

func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("tek sayilar", i)
		time.Sleep(time.Second)
	}
}
