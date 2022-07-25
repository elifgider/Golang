package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("kaydedildi", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")

}
func Demo3() {

	p := product{"Laptop", 8000}

	p = product{"mouse", 200}
	fmt.Println("bitti")
	defer p.Save()
}
