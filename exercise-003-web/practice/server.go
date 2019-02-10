package main

import (
    "log"
    "net/http"
    "time"
)

func home(respWriter http.ResponseWriter, request *http.Request) {
    respWriter.Header().Set("Content-Type", "application/json")
    go PrintStuff(request)
    respWriter.Write([]byte("Thanks\n"))
}

func PrintStuff(request *http.Request) {
    request.ParseForm()
    time.Sleep(time.Second * 3)
    log.Println(request.PostForm)
}

func main() {
    log.Println("Starting server...")

    http.HandleFunc("/", home)

    log.Println("Now listening on port 8080")
    http.ListenAndServe(":8080", nil)
}


