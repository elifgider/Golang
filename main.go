package main

import (
	"golesson/project"
)

//------------------POINTERS---------------------------------------------------------------------------------------------------------------------------------------------------------------
// sayilar := []int{1, 2, 3, 4}
// pointers.Demo2(sayilar)
// fmt.Println("Maindeki ilk eleman: ", sayilar[0]
//------------------STRUCTS-----------------------------------------------------------------------------------------------------------------------------------------------------------------
//structs.Demo2()
//------------------CHANNELS---------------------------------------------------------------------------------------------------------------------------------------------------------------
// ciftSayiToplamCn := make(chan int)
// tekSayiToplamCn := make(chan int)
// go channels.CiftSayilar(ciftSayiToplamCn)
// go channels.TekSayilar(tekSayiToplamCn)

// ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamCn, <-tekSayiToplamCn
// carpim := ciftSayiToplam * tekSayiToplam
// fmt.Println("Çarpım", carpim)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
func main() {

	//fmt.Println(error_handling.TahminEt2(-5))

	project.GetAllProducts()
	project.AddProduct()

}
