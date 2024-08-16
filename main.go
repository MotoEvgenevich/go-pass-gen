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

	var result []byte
	rand.Seed(uint64(time.Now().UnixNano()))
	for len(result) <= passLen {
		time.Sleep(time.Millisecond * 10)
		randomNum := rand.Intn(12) + 1
		switch {
		case randomNum < 4:
			result = append(result, randomChar(97, 122))
		case randomNum > 3 && randomNum < 7 && digitsParam:
			result = append(result, randomChar(97, 122))
		case randomNum > 6 && randomNum < 10 && uppercaseLetterParam:
			result = append(result, randomChar(65, 90))
		case randomNum > 9 && specSymbolParam:
			result = append(result, generateSpecialSymbol())
		}
	}
	fmt.Println(string(result))
	return string(result)
}

func randomChar(min, max int) byte {
	return byte(rand.Intn(max-min+1) + min)
}

func generateSpecialSymbol() byte {
	specialRanges := [][]int{
		{33, 47}, {58, 64}, {91, 96}, {123, 126},
	}
	selectedRange := specialRanges[rand.Intn(len(specialRanges))]
	return randomChar(selectedRange[0], selectedRange[1])
}
