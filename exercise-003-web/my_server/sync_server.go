package main

import (
  "html/template"
  "net/http"
  "fmt"
  "sync"
)

var homeT = template.Must(template.ParseFiles("my_server/home.html"))

var signups = struct{
  sync.RWMutex
  m map[string]int
}{m: make(map[string]int)}


func home(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  
  signups.Lock()
  count(username)
  signups.Unlock()

  signups.RLock()
  homeT.Execute(w, &signups.m)
  signups.RUnlock()
}

func count(username string) {
  if (username == "") { return }
  if _, ok := signups.m[username]; ok {
    signups.m[username]++
  } else {
    signups.m[username] = 1
  }
}



func main() {
  //bind handler
  http.HandleFunc("/", home)


  //start server
  fmt.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", nil)
}
