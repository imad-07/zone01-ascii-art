package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("there is no text\n")
		return
	}
	// check if all of the characters of the input are in the range of 32-126 in the ascii table
	for i := 0; i < len(args[0]); i++ {
		if args[0][i] < 32 || args[0][i] > 126 {
			fmt.Printf("error in input\n")
			return
		}
	}
	// here we split our input with new lines while keeping each one of them in an indexed place in the array
	word := split(args[0])
	fileContent, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("error in standard file")
		return
	}
	// here we get the standard art from the file that they gave us 
	lettres := getLettres(fileContent)
	bl := false
	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if l == len(word)-1 {
				continue
			}
			if bl && word[l+1] != "\n" {
				continue
			}
			fmt.Printf("\n")
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				fmt.Printf(lettres[word[l][j]-32][i])
			}
			fmt.Print("\n")
		}
		bl = true
	}
}

func split(str string) []string {
	word := ""
	splitedword := []string{}
	skip := false
	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			skip = true
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	splitedword = append(splitedword, word)
	return splitedword
}

func getLettres(fileContent []byte) [][]string {
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if i != len(fileContent)-1 && fileContent[i] == '\n' && fileContent[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if fileContent[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, fileContent[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}
