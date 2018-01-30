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
  var map_frequencies map[string]int
  map_frequencies = corpus.Analyze(content)
  var frequencies corpus.ByFreq
  frequencies = corpus.MapToWordCount(map_frequencies)
  sort.Sort(frequencies)
 
  //display on stdout 
  for _, freq := range frequencies {
    fmt.Println(freq)
  }
 
}
