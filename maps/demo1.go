package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)

	sozluk["table"] = "masa"
	sozluk["pencil"] = "kalem"
	sozluk["map"] = "harita"
	fmt.Println(sozluk["table"])
	fmt.Println("sozluk: ", sozluk)
	delete(sozluk, "map")
	fmt.Println(sozluk)

	deger, varMi := sozluk["t"]
	fmt.Println(deger)
	fmt.Println(varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}
