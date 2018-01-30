package main

import (
	"fmt"
  "os"
  "github.com/crajcan/saigo/exercise-001-corpus/corpus"
  "sort"
)



func main() {

   //get filename
  filename := os.Args[1] 
  if filename == "" {
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
  stat, err := file.Stat()
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
  var map_result map[string]int
  map_result = corpus.Analyze(content)
  var result corpus.ByFreq
  result = corpus.MapToWordCount(map_result)
  sort.Sort(result)
 
  //display on stdout 
  for _, freq := range result {
    fmt.Println(freq)
  }
 
}
