package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	leight := 10
	digitsParam := true
	uppercaseLetterParam := true
	specSymbolParam := true
	randomPassGenerator(leight, digitsParam, uppercaseLetterParam, specSymbolParam)
}

// 97-122 lowercase eng letters
// 65-90 uppercase eng letters
// 48-57 diggits

func randomPassGenerator(passLen int, digitsParam bool, uppercaseLetterParam bool, specSymbolParam bool) string {

	result := ""

	for len(result) <= passLen {
		rand.Seed(uint64(time.Now().UnixNano()))
		time.Sleep(time.Millisecond * 10)
		min := 1
		max := 12
		randomNum := rand.Intn(max-min+1) + min
		switch {
		case randomNum < 4:
			result += getRandomCharBetween(97, 122) // lowerCaseGenerate
		case randomNum > 3 && randomNum < 7 && digitsParam:
			result += getRandomCharBetween(48, 57) // digitsGenerate
		case randomNum > 6 && randomNum < 10 && uppercaseLetterParam:
			result += getRandomCharBetween(65, 90) // uppercaseLetterGenerate
		case randomNum > 9 && specSymbolParam:
			result += specSymbolGenerate()
		}
	}
	fmt.Println(result)
	return result
}

func getRandomCharBetween(min int, max int) string {
	randomNum := rand.Intn(max-min+1) + min
	return string(byte(randomNum))
}

func specSymbolGenerate() string {
	specialRanges := [][]int{
		{33, 47},
		{58, 64},
		{91, 96},
		{123, 126},
	}

	// Выбираем случайный диапазон
	rangeIndex := rand.Intn(len(specialRanges))
	selectedRange := specialRanges[rangeIndex]
	randomChar := rand.Intn(selectedRange[1]-selectedRange[0]+1) + selectedRange[0]
	return string(byte(randomChar))
}
