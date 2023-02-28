package main

// import of the exercise test didn't work.
// i'll run it in the playground to check if it works.
import (
	// "fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	// wc := WordCount("hello world hello")
	// fmt.Println(wc)
	// prints map[hello:2 world:1]
	// wc.Test(WordCount)
	// returns an error normally because local file doesn't have a wc import
	// uncomment it to run on the playground.
}

/*
the result i got is:

​
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

  ALL PASSES!

*/
