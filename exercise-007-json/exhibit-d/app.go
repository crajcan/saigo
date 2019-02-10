package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Phone ...
type Phone struct {
	Name string      `json:"name"`
  Age  int         `json:"age"`
  ID   string      `json:"id"`
  Carrier string   `json:"carrier"`
  ImageUrl string  `json:"imageUrl"`
  Snippet string   `json:"snippet"`
}

var allPhones []Phone

func setup() {
	data, err := ioutil.ReadFile("phones.json")
	if err != nil {
		fmt.Println("Error reading phones.json")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &allPhones)
	if err != nil {
		fmt.Println("Error in unmarshalling phones")
	}
  //fmt.Println(allPhones)
}

func phones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(allPhones)
}

func main() {
	setup()
	http.HandleFunc("/phones", phones)
  fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
