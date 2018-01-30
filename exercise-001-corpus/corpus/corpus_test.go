package corpus

import "testing"
import "sort"

type testpair struct{
  source []byte
  spec map[string]int
}

var countTests = []testpair{
  
  { []byte(" "),
    map[string]int{},
  },

  { []byte("The quick brown fox"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    }, 
  },

  { []byte("The"), 
    map[string]int{
      "the": 1,
    },
  },

  { []byte("?The quick brown fox"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    }, 
  },

  { []byte("The quick brown fox?"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    }, 
  },

  { []byte("?The quick.brown fox!"), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    }, 
  },

  { []byte("The quick the tHE"), 
    map[string]int{
      "the": 3,
      "quick": 1,
    }, 
  },

  { []byte(" The quick brown fox "), 
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    }, 
  },

  { []byte(" The the quick the brown The fox Fox "), 
    map[string]int{
      "the": 4,
      "quick": 1,
      "brown": 1,
      "fox": 2,
    },
  },

}

func TestCountWordFreq(t *testing.T){
  for _, pair := range countTests{
    result := Analyze(pair.source)
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


type toWordTestPair struct{
  mp map[string]int
  arr []wordCount
}

var toWordTests = []toWordTestPair{
  
  {
    map[string]int{},
    []wordCount{},
  },
 
  {
    map[string]int{
      "the": 1,
    },
    []wordCount{
      {"the", 1},
    },   
  },

  {
    map[string]int{
      "the": 4,
      "quick": 1,
      "brown": 1,
      "fox": 2,
    },
    []wordCount{
      {"the", 4},
      {"fox", 2},
      {"brown", 1},
      {"quick", 1},
    },
  },
    
  {
    map[string]int{
      "the": 1,
      "quick": 1,
      "brown": 1,
      "fox": 1,
    }, 
    []wordCount{
      {"brown", 1},
      {"fox", 1},
      {"quick", 1},
      {"the", 1},
    },
  },
}


func TestMapToWordCount(t *testing.T){
  for i, pair := range toWordTests {
    newarr := MapToWordCount(pair.mp)
    sort.Sort(ByFreq(newarr))
    for j, count := range newarr {
      if (count != pair.arr[j]) {
        t.Error(
          "For count", i,j,
          "expected ", pair.arr[j],
          "got", count,
        )       
      } 
    }
  }
}        
  
