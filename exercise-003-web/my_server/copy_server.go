package main

import (
  "html/template"
  "net/http"
  "fmt"
  //"sync"
  //"errors"
  //"strconv"
)

var signups = map[string]int{}
var c chan string = make(chan string)
var i int = 0

var homeT = template.Must(template.ParseFiles("my_server/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
  i++
  r.ParseForm()
  username := r.Form.Get("username")
  //username += strconv.Itoa(i)
  

  /*var signupsCopy map[string]int{}
  for k,v := range signUps {
    signupsCopy[k] = v
  }*/
  //fmt.Println("Parsed username: ", username)
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
