package corpus

import (//"fmt"
        //"sort"
        "strings"
        "unicode"
)


func Analyze(content []byte) map[string]int{
  result := make(map[string]int)
  current_word := ""
  for _, character := range content {
    if unicode.IsLetter(rune(character)) {
      current_word += strings.ToLower(string(character))
    } else {
      count(current_word, result)
      current_word = ""
    }
  }
  count(current_word, result)
  return result
}

func count(word string, result map[string]int) {
  if (word == "") { return }
  if _, ok := result[word]; ok {
    result[word]++   
  } else {
    result[word] = 1    
  }
}


type wordCount struct{
  word string
  count int
}


func MapToWordCount(countMap map[string]int) []wordCount {
  var result []wordCount
  for word, count := range countMap {
    result = append(result, wordCount{word, count})
  }
  return result
}


type ByFreq []wordCount

func (this ByFreq) Len() int {
  return len(this)
}

func (this ByFreq) Less(i, j int) bool {
  if (this[i].count == this[j].count) {
    return (this[i].word < this[j].word)
  } else { 
    return (this[i].count > this[j].count)     
  }
}

func (this ByFreq) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}



