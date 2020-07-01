package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("\tError! Empty input.")
	} else if len(args) > 0 {
		arg := os.Args[1]
		style := "standard.txt"
		if len(args) >= 2 {
			style = args[1]
			style += ".txt"
		}
		for _, val := range arg {
			if isValid(val) == false {
				log.Fatal("\tError! Not ASCII.")
			}
		}
		ascii := rangeLines(arg, style)
		if len(args) == 1 || len(args) == 2 {
			fmt.Println(string(ascii))
		} else if len(args) >= 3 {
			thirdFlag := args[2]
			if thirdFlag[:9] == "--output=" {
				fileName := thirdFlag[9:]
				for i := 3; i < len(args); i++ {
					fileName += string(" " + args[i])
				}
				output(fileName, ascii)
			}
		}
	}
}

func isValid(val rune) bool {
	if val >= ' ' && val <= '~' {
		return true
	}
	return false
}

func getBorder(content []byte, fileLine int) (int, int) {
	borderFileLine := 0
	borderEnd := 0
	b := -1
	for idx := range content {
		if content[idx] == '\n' {
			borderFileLine++
		}
		if borderFileLine == fileLine-1 {
			borderEnd = idx + 1
			b++
		}
	}
	return (borderEnd - b), borderEnd
}

func rangeLines(arg string, style string) []byte {
	content, err := ioutil.ReadFile(style)
	if err != nil {
		log.Fatal(err)
	}

	splited := strings.Split(arg, "\\n")
	fileLine := 0
	var slice []string
	for _, word := range splited {
		for line := 1; line <= 8; line++ {
			for _, val := range word {
				fileLine = (int(val)-32)*9 + 1 + line
				start, end := getBorder(content, fileLine)
				slice = append(slice, string(content[start:end]))
			}
			slice = append(slice, string("\n"))
		}
	}
	var sliceByte []byte
	for idxSlice, slicedString := range slice {
		for idxBytes, bytes := range slicedString {
			if !(len(slice)-1 == idxSlice && len(slicedString)-1 == idxBytes && bytes == '\n') {
				sliceByte = append(sliceByte, byte(bytes))
			}
		}
	}
	return sliceByte
}

func output(fileName string, ascii []byte) {
	err := ioutil.WriteFile(fileName, ascii, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
