package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

func main() {

	leight := flag.Int("len", 8, "leight of password")
	digitsParam := flag.Bool("digit", true, "digits in password")
	uppercaseLetterParam := flag.Bool("upletter", true, "uppercase letter in password")
	specSymbolParam := flag.Bool("spec", true, "spec symbol in password")
	save := flag.Bool("save", true, "Save to file $FILENAME")

	flag.Parse()

	pass := randomPassGenerator(*leight, *digitsParam, *uppercaseLetterParam, *specSymbolParam)

	if *save {
		file, err := os.OpenFile("password.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(pass + "\n")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("file with pass create success. Pass is:", pass)
	}
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
			result += lowerCaseGenerate()
		case randomNum > 3 && randomNum < 7 && digitsParam:
			result += digitsGenerate()
		case randomNum > 6 && randomNum < 10 && uppercaseLetterParam:
			result += uppercaseLetterGenerate()
		case randomNum > 9 && specSymbolParam:
			result += specSymbolGenerate()
		}
	}
	fmt.Println(result)
	return result
}

func lowerCaseGenerate() string {
	min := 97
	max := 122
	randomNum := rand.Intn(max-min+1) + min
	return string(byte(randomNum))
}

func digitsGenerate() string {
	min := 48
	max := 57
	randomNum := rand.Intn(max-min+1) + min
	return string(byte(randomNum))
}

func uppercaseLetterGenerate() string {
	min := 65
	max := 90
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

/*
все эти функции можно в принципе выделить в одну куда параметром передавать min и max или даже rangesm как в случае спецсимволов
*/

func universallGenerate() string {
	result := ""
	return result
}
