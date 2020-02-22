package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func count(field string, fields []string) int{
	count := 0
  for _, value := range fields {
    if field == value {
      count++
    }
  }
  return count;
}

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	fields := strings.Fields(s)
	for i := range fields {
		wordCountMap[fields[i]] = count(fields[i], fields)
	}
	fmt.Println(wordCountMap)
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
