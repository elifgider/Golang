package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}
func C() {
	fmt.Println("C fonksiyonu çalıştı")
}
func D() {
	fmt.Println("D fonksiyonu çalıştı")
}
func E() {
	fmt.Println("E fonksiyonu çalıştı")
}
func B() {
	defer A()
	defer C()
	defer E()
	defer D()
	fmt.Println("B fonksiyonu çalıştı")

}
