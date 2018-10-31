package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "ardrone3.xml"
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error: os.Open: %v\n", err)
	}
	defer f.Close()

	//bufio lets us read files line by line
	fReader := bufio.NewReader(f)

	//Start with the first line
	lineNR := 1

	for {
		//read a line
		line, _, err := fReader.ReadLine()
		if err != nil {
			log.Printf("Error: bufio.ReadLine: %v\n", err)
			break
		}

		//Remove leading spaces in line
		tmpLine := strings.TrimSpace(string(line))
		line = []byte(tmpLine)
		//printLine(line)

		//Look for the start tag called <project>
		found := tagProjectStart("<project", line)
		if found {
			fmt.Println("Found project start on lineNR : ", lineNR)
			if err := checkForClosingBracket(line); err != nil {
				log.Fatal("Error: missed closing bracket, incomplete line!")
			}

			lastPosition := findWord(line, "name=")
			if lastPosition > 0 {
				fmt.Println("Looking for word in 'project' Last Position = ", lastPosition)
			} else {
				fmt.Println("Did not find the word you were looking for.")
			}

		}

		//Look for the end tag called </project>
		found = tagProjectEnd(line)
		if found {
			fmt.Println("Found project end on lineNR : ", lineNR)
			if err := checkForClosingBracket(line); err != nil {
				log.Fatal("Error: missed closing bracket, incomplete line!")
			}
		}

		lineNR++
	}

}

func printLine(line []byte) {
	//fmt.Printf("Line : %v \n Type %T\n", line, line)
	for i := 0; i < len(line); i++ {
		character := string(line[i])
		fmt.Print(character)

	}
	fmt.Println()
}

//TODO: Gjør så denne returnerer true/false og ikke type, og
//		flytt type logikken over i main.
//
//tagProjectStart will check if there is a <project> tag in xml
func tagProjectStart(theWord string, line []byte) (found bool) {
	var tag string
	if len(line) > 0 {
		tag = string(line[0:len(theWord)])
		if tag == theWord {
			return true
		}
	}

	//If no found, return an empty struct of type tagProject

	return false
}

//tagProjectEnd will check if there is a <project> tag in xml
func tagProjectEnd(line []byte) (found bool) {
	var tag string
	if len(line) > 0 {
		tag = string(line[0:9])
		if tag == "</project" {
			return true
		}
	}

	//If no found, return an empty struct of type tagProject

	return false
}

//findWord looks for a word, and returns the position the last character found in slice.
// Returns zero if no word was found.
func findWord(line []byte, myWordString string) (lastPosition int) {
	//find word in []byte
	myWordByte := []byte(myWordString)
	foundWord := false

	for linePosition := 0; linePosition < len(line)-len(myWordByte); linePosition++ {
		wordPosition := 0
		for {

			//Since the iteration over the word using wordPosition as a counter will break out
			// if there is a mismatch in the matching, we can be sure that the word was found
			// if word position reaches the same value as the length of the word.
			// And we can then return the result and exit.
			if wordPosition >= len(myWordByte) {
				fmt.Println("Reached the end of the word, breaking out of word loop", linePosition, wordPosition)
				foundWord = true
				lastPosition = linePosition + wordPosition
				return lastPosition
			}

			//If there is no match break out of the loop imediatly, since there is no reason
			// to continue if one fails. Better to break out of the inner for loop and start
			// the iteration of the next charater and see if we are more lucky.
			if line[linePosition+wordPosition] != myWordByte[wordPosition] {
				break
			}

			wordPosition++
		}

		fmt.Println("---- Value of foundWord = ", foundWord)
		if foundWord {
			fmt.Println("Breaking out of outer loop")
			break
		}
	}
	return 0
}

//checkForClosingBracker
//Check for opening and closing angle bracket.
//Will return nil if both start and end bracker was found.
func checkForClosingBracket(line []byte) error {
	for i := 0; i < len(line); i++ {
		character := string(line[i])
		if character == "<" {
			ii := 0
			for {
				if string(line[ii]) == ">" {
					//fmt.Println("Found closing angular bracket at position: ", ii)
					break
				}
				if ii == len(line)-1 {

					return errors.New("Missing ending angular bracket")
				}
				ii++
			}
		}
	}
	return nil
}
