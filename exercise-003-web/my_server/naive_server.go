package main

import (
  "html/template"
  "net/http"
  "fmt"
)

var signups = map[string]int{}

var homeT = template.Must(template.ParseFiles("my_server/home.html"))

func home(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form.Get("username")
  fmt.Printf("username: %s", username)
  count(username, signups)
  homeT.Execute(w, &signups)
}

func count(user string, signups map[string]int) {
  if (user == "") { return }
  if _, ok := signups[user]; ok {
    signups[user]++
  } else {
    signups[user] = 1
  }
}

func main() {
  http.HandleFunc("/", home)
  fmt.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", nil)
}
