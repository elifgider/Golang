package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"Id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body) //byte formatnda verir.
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todo Todo
	json.Unmarshal(bodyBytes, &todo) //string formatını json formatına çevirir.
	fmt.Println(todo)                //struct yapısında okuma yaparız jsondan todo ya çevirme

}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo) //marshaldan todo ya
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://jsonplaceholder.typicode.com/todos", "applicaton/json;charset-utf-8",
		bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body) //byte formatnda verir.
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todo) //string formatını json formatına çevirir.
	fmt.Println(todoResponse)
}
