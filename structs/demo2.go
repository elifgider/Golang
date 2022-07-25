package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	id        int
	age       int
}

func (c1 customer) save() {
	fmt.Println("Eklendi", c1.id)
}

func Demo2() {
	c1 := customer{firstName: "Elif", lastName: "Gider", id: 0000, age: 22}
	c1.save()
	c1.update()

}

func (c1 customer) update() {
	fmt.Println("Müşteri bilgileri güncellendi.", c1.id)
}
