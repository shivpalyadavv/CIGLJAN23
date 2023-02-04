package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// monorepo
type Book struct {
	Title  string `json:"booktitle"`
	Author string `json:"author-1"`
	Page   int
	Price  float32
}

func main() {
	// book := Book{}
	// book.Title = "Game of thrones"
	// book.Author = "Ashish"
	// book.Page = 1000
	// book.Price = 2.5

	// //book object -> book.json
	// // Marshal ====  Golang object ---> byteArray --> write in a file or print(that will be in json format)
	// byteArray, err := json.Marshal(book)
	// if err != nil {
	// 	log.Fatal("got error", err)
	// }
	// log.Println(string(byteArray))

	// json -> book object
	// Unmarshal ==> data-->bytearray----> golang struct
	book2 := Book{}
	//jsondata := `{"booktitle":"some title","author-1":"Ashish","Page":1000,"Price":2.5}`

	fileDataArr, err := ioutil.ReadFile("./book.json")
	if err != nil {
		log.Fatal("got error while ReadFile", err)
	}
	//log.Println(string(fileDataArr))
	jsondata := string(fileDataArr)
	err = json.Unmarshal([]byte(jsondata), &book2)
	if err != nil {
		log.Fatal("got error while Unmarshal", err)
	}
	fmt.Println(book2) //some title
}
