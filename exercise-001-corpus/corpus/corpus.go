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




