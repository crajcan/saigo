package corpus

import ("strings"
        "unicode"
)


//Populate map with (word, frequency) pairs from input text
func Analyze(content []byte) map[string]int{
  freqs := make(map[string]int)
  current_word := ""
  for i, character := range content {
    if isAlphanumeric(character) || isPartOfNumeral(i, character, content) {
      current_word += strings.ToLower(string(character))
    } else {
      count(current_word, freqs)
      current_word = ""
    }
  }
  count(current_word, freqs)
  return freqs
}


//helper for Analyze
func count(word string, freqs map[string]int) {
  if (word == "") { return }
  if _, ok := freqs[word]; ok {
    freqs[word]++   
  } else {
    freqs[word] = 1    
  }
}

func isAlphanumeric(character byte) bool {
  return ((unicode.IsLetter(rune(character))) || (unicode.IsDigit(rune(character))))
}

func nextAlphanumeric(i int, content []byte) bool {
  return (((i+1) < len(content)) && isAlphanumeric(content[i+1]))
}

func isSeparator(character byte) bool {
  return (character == 46 || character == 44)
}

func isPartOfNumeral(i int, character byte, content []byte) bool {
  return (isSeparator(character) && nextAlphanumeric(i, content)) 
}


//Change map to array of structs to make use of sorting

type wordCount struct{
  word string
  count int
}

type FreqArr []wordCount


func MapToWordCount(countMap map[string]int) []wordCount {
  var result []wordCount
  for word, count := range countMap {
    result = append(result, wordCount{word, count})
  }
  return result
}


//Implement sorting interface


func (this FreqArr) Len() int {
  return len(this)
}


func (this FreqArr) Less(i, j int) bool {
  if (this[i].count == this[j].count) {
    return (this[i].word < this[j].word)
  } else { 
    return (this[i].count > this[j].count)     
  }
}

func (this FreqArr) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}



