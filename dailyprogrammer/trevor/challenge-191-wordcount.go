package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"sort"
)

type Word struct {
  value string
  frequency int
}

type ByFrequency []Word
func (slice ByFrequency) Len() int {
  return len(slice)
}
func (slice ByFrequency) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
func (slice ByFrequency) Less(i, j int) bool {
	return slice[i].frequency > slice[j].frequency
}

// regex for filtering out non-alpha characters
var alpha, _ = regexp.Compile("[^a-z]")

// regex for tokenising the input text file
var separator, _ = regexp.Compile("[ \\-\n]")

func main() {

  if body, err := ioutil.ReadFile("./book.txt"); err == nil {

	  word_count := make(map[string]int)
	  for _, word := range separator.Split(string(body), -1) {
		if filtered := alpha.ReplaceAllString(strings.ToLower(word), ""); filtered != "" {
			word_count[filtered]++
		}
	  }

	  words := make([]Word, 0, len(word_count))
	  for key := range word_count {
		  words = append(words, Word{key, word_count[key]})
	  }
	  sort.Sort(ByFrequency(words))

	  for _, word := range words {
		fmt.Println(fmt.Sprintf("%v: %v", word.value, word.frequency))
	  }

  } else {
    fmt.Println(err)
  }

}

