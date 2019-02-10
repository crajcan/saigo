package main

import (
  "html/template"
  "net/http"
  "fmt"
  "sync"
)

var homeT = template.Must(template.ParseFiles("my_server/home.html"))
var c chan string = make(chan string)

var signups = struct{
  sync.RWMutex
  m map[string]int
}{m: make(map[string]int)}


func home(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  signupsCopy := map[string]int{}
  signups.RLock()
  for k,v := range signups.m {
    signupsCopy[k] = v
  }
  if _, ok := signupsCopy[username]; ok{
    signupsCopy[username]++
  } else {
    signupsCopy[username] = 1
  }
  signups.RUnlock()

  c <- username

  signups.RLock()
  homeT.Execute(w, &signupsCopy)
  signups.RUnlock()
}

func count(c <-chan string) {
  for {
    username := <- c
    if (username != "") { 
      if _, ok := signups.m[username]; ok {
        signups.Lock()
        signups.m[username]++
        signups.Unlock()
      } else {
        signups.Lock()
        signups.m[username] = 1
        signups.Unlock()
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
