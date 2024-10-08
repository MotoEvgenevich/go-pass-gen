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
	var result string
	rand.Seed(uint64(time.Now().UnixNano()))

	for len(result) < passLen {
		randomNum := rand.Intn(12) + 1
		switch {
		case randomNum < 4:
			result += universalGenerate(97, 122) // lowercase letters
		case randomNum > 3 && randomNum < 7 && digitsParam:
			result += universalGenerate(48, 57) // digits
		case randomNum > 6 && randomNum < 10 && uppercaseLetterParam:
			result += universalGenerate(65, 90) // uppercase letters
		case randomNum > 9 && specSymbolParam:
			result += universalGenerateRanges([][]int{
				{33, 47},
				{58, 64},
				{91, 96},
				{123, 126},
			})
		}
	}
	return result
}

func universalGenerate(min int, max int) string {
	randomNum := rand.Intn(max-min+1) + min
	return string(byte(randomNum))
}

func universalGenerateRanges(ranges [][]int) string {
	rangeIndex := rand.Intn(len(ranges))
	selectedRange := ranges[rangeIndex]
	return universalGenerate(selectedRange[0], selectedRange[1])
}
