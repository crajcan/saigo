package main

import (
	"fmt"
  "os"
//	"sort"
//	"strings"
)



func main() {

   //get filename
  filename := os.Args[1] 
  if filename == nil {
    fmt.Println("Please enter a filename")
    return
  }

  //open file
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  //get file size
  stat, err := fileStat()
  if err != nil {
    return
  }

  //read file
  content := make([]byte, stat.Size())
  _, err = file.Read(content)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  //analyze
  var result map[string]int
  result := Analyze(content)   
  fmt.Println(result)
 
}
