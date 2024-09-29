package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

type wordToCount struct {
	word  string
	count int
}

var r = regexp.MustCompile(`[^\s]+`)

func Top10(str string) []string {
	allWords := r.FindAllString(str, -1)
	wordCountMap := make(map[string]int)
	for _, word := range allWords {
		wordCountMap[word]++
	}
	wordsCount := make([]wordToCount, 0, len(wordCountMap))
	for key, value := range wordCountMap {
		wordsCount = append(wordsCount, wordToCount{key, value})
	}
	sort.Slice(wordsCount, func(i, j int) bool {
		return wordsCount[i].count > wordsCount[j].count
	})
	strSlice := []string{}
	prevCount := 0
	prevWordIndex := 0
	for index, elem := range wordsCount {
		if elem.count != prevCount {
			tempSlice := strSlice[prevWordIndex:index]
			sort.Slice(tempSlice, func(i, j int) bool {
				return tempSlice[i] < tempSlice[j]
			})
			if index < 10 {
				strSlice = append(strSlice, wordsCount[index].word)
				prevCount = elem.count
				prevWordIndex = index
				continue
			}
			break
		} else {
			if index < 10 {
				strSlice = append(strSlice, wordsCount[index].word)
				if index == len(wordsCount)-1 || index == 9 {
					tempSlice := strSlice[prevWordIndex:index]
					sort.Slice(tempSlice, func(i, j int) bool {
						return tempSlice[i] < tempSlice[j]
					})
				}
				continue
			}
			break
		}
	}
	return strSlice
}
