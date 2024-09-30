package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

type wordToCount struct {
	word  string
	count int
}

var r = regexp.MustCompile(`[^\s]+`)

// sortSubSlice сортирует слайс строк между элементами startIndex endIndex.
func sortSubSlice(inSlc []string, startIndex int, endIndex int) {
	tempSlice := inSlc[startIndex:endIndex]
	sort.Slice(tempSlice, func(i, j int) bool {
		return tempSlice[i] < tempSlice[j]
	})
}

// countWordFreq возвращает все встреченные слова и их частоту для строки str.
func countWordFreq(str string) []wordToCount {
	allWords := r.FindAllString(str, -1)
	wordCountMap := make(map[string]int)
	for _, word := range allWords {
		word = strings.ToLower(word)
		runeWord := []rune(word)
		wordLen := len(runeWord)
		if !unicode.IsDigit(runeWord[wordLen-1]) && !unicode.IsLetter(runeWord[wordLen-1]) {
			runeWord = runeWord[:wordLen-1]
		}
		if len(runeWord) > 0 && (!unicode.IsDigit(runeWord[0]) && !unicode.IsLetter(runeWord[0])) {
			runeWord = runeWord[1:]
		}
		word = string(runeWord)
		if word != "" {
			wordCountMap[word]++
		}
	}
	wordsCount := make([]wordToCount, 0, len(wordCountMap))
	for key, value := range wordCountMap {
		wordsCount = append(wordsCount, wordToCount{key, value})
	}
	sort.Slice(wordsCount, func(i, j int) bool {
		return wordsCount[i].count > wordsCount[j].count
	})
	return wordsCount
}

func Top10(str string) []string {
	const returnSliceLen = 10
	strSlice := []string{}
	prevCount := 0
	prevWordIndex := 0
	wordsCount := countWordFreq(str)
	for index, elem := range wordsCount {
		// Лексикографисечки отсортировать слова с одинаковочй частотой.
		if elem.count != prevCount {
			sortSubSlice(strSlice, prevWordIndex, index)
			if index >= returnSliceLen {
				break
			}
			strSlice = append(strSlice, wordsCount[index].word)
			prevCount = elem.count
			prevWordIndex = index
		} else {
			strSlice = append(strSlice, wordsCount[index].word)
		}
		sortSubSlice(strSlice, prevWordIndex, len(strSlice))
	}
	if len(strSlice) > returnSliceLen {
		return append([]string(nil), strSlice[:returnSliceLen]...)
	}
	return strSlice
}
