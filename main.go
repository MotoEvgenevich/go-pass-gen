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

	length := flag.Int("len", 8, "length of password")
	digitsParam := flag.Bool("digit", true, "digits in password")
	uppercaseLetterParam := flag.Bool("upletter", true, "uppercase letter in password")
	specSymbolParam := flag.Bool("spec", true, "spec symbol in password")
	save := flag.Bool("save", true, "Save to file $FILENAME")
	filename := flag.String("filename", "password.txt", "Filename to save the password")

	flag.Parse()

	pass := randomPassGenerator(*length, *digitsParam, *uppercaseLetterParam, *specSymbolParam)

	if *save {
		file, err := os.OpenFile(*filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(pass + "\n")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Password saved to file:", *filename)
		fmt.Println("Generated password:", pass)
	} else {
		fmt.Println("Generated password:", pass)
	}
}

// 97-122 lowercase eng letters
// 65-90 uppercase eng letters
// 48-57 diggits

func randomPassGenerator(passLen int, digitsParam bool, uppercaseLetterParam bool, specSymbolParam bool) string {
	var result string
	rand.Seed(uint64(time.Now().UnixNano()))

	for len(result) < passLen {
		randomNum := rand.Intn(4) + 1
		switch {
		case randomNum == 1:
			result += universalGenerate(97, 122) // lowercase letters
		case randomNum == 2 && digitsParam:
			result += universalGenerate(48, 57) // digits
		case randomNum == 3 && uppercaseLetterParam:
			result += universalGenerate(65, 90) // uppercase letters
		case randomNum == 4 && specSymbolParam:
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
