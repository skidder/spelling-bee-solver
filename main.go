package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	wordsFile        = flag.String("words", "words.txt", "Path to words file")
	requiredLetter   = flag.String("requiredLetter", "", "Letters that must be used")
	availableLetters = flag.String("availableLetters", "", "Letters that can be used")
	minLength        = flag.Int("minLength", 4, "Minimum length of words")
)

func main() {
	flag.Parse()

	var requiredLetterRune rune
	for _, c := range *requiredLetter {
		requiredLetterRune = c
		break
	}
	availableLettersMap := make(map[rune]bool, len(*availableLetters))
	for _, letter := range *availableLetters {
		availableLettersMap[letter] = true
	}

	log.Print("Reading file...")
	file, err := os.Open(*wordsFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// skip words that aren't long enough
		if len(line) < *minLength {
			continue
		}
		letters := findUniqueLetters(line)

		// check for required letter
		matched := false
		for _, letter := range letters {
			if letter == requiredLetterRune {
				matched = true
				break
			}
		}
		if !matched {
			continue
		}

		// check remaining letters
		for _, letter := range letters {
			if letter == requiredLetterRune {
				continue
			}
			if _, found := availableLettersMap[letter]; !found {
				matched = false
				break
			}
		}
		if matched {
			log.Println(line)
		}
	}
	if scanner.Err() != nil {
		log.Fatalf("Error reading words file: %v", err)
	}

}

// findUniqueLetters returns an array of the unique letters in the supplied string.
func findUniqueLetters(word string) []rune {
	lettersMap := map[rune]bool{}
	for _, letter := range word {
		lettersMap[letter] = true
	}

	i := 0
	letters := make([]rune, len(lettersMap))
	for letter := range lettersMap {
		letters[i] = []rune(strings.ToLower(string(letter)))[0]
		i++
	}
	return letters
}
