package main

import (
  "html/template"
  "net/http"
  "fmt"
)

var signups = map[string]int{}
var c chan string = make(chan string)

var homeT = template.Must(template.ParseFiles("my_server/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  
  c <- username
   
  homeT.Execute(w, &signups)
}

func count(c <-chan string) {
  for username := range c {
    if (username != "") { 
      if _, ok := signups[username]; ok {
        signups[username]++
      } else {
        signups[username] = 1
      }
    }
  } 
}


func main() {
  //bind handler
  go count(c)
  http.HandleFunc("/", home)


  //start server
  fmt.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", nil)
}
