package corpus

import "testing"
import "fmt"

type testpair struct{
  source []byte
  spec map[string]int
}

var tests = []testpair{
  
  { []byte("The quick brown fox"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    } },

 { []byte("The"), 
    map[string]int{
      "the": 1,
    } },

  { []byte("?The quick brown fox"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    } },

  { []byte("The quick brown fox?"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    } },

  { []byte("?The quick.brown fox!"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    } },

  { []byte("The quick the tHE"), 
    map[string]int{
      "the": 3,
      "quick": 1,
    } },

  { []byte(" The quick brown fox "), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    } },

}

func TestCountWordFreq(t *testing.T){
  for _, pair := range tests{
    result := Analyze(pair.source)
    fmt.Println("result =", result)
    for word, count := range result {
      if (count != pair.spec[word]) {
        t.Error(
          "For word", word,
          "expected count", pair.spec[word],
          "got", count,  
        )
      }
    }
  }
}




