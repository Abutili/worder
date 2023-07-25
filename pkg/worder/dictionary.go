package worder

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type LetterNode struct {
	Letter rune
	Parent *LetterNode
	Child  *LetterNode
}

type Word struct {
	Word  string
	Runes []rune
	Graph *LetterNode
}

type Dictionary struct {
	Name        string
	Words       map[string]Word
	SortedWords []string
}

// func (d *Dictionary) createSortedSlice() {
// 	for _, w := range d.Words {

// 	}
// }

// Check if the word exists in this dictionary
func (d *Dictionary) IsValidWord(test_word string) bool {
	if _, ok := d.Words[strings.ToUpper(test_word)]; ok {
		return true
	}
	return false
}

func NewDictionary(name, path string) (d Dictionary) {

	words := map[string]Word{}

	// load words in file rd.Path
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	lineno := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineno++
		// remove leading and trailing space and control characters
		line := strings.TrimSpace(scanner.Text())

		// comment
		if strings.HasPrefix(line, "#") {
			log.Println(lineno, " Found an comment line, skipping.")
			continue
		}

		// multiple words on a single line means skip it
		if len(strings.Fields(line)) > 1 {
			log.Println(lineno, " Found a line with multiple words, skipping")
			continue
		}

		// skip empty lines
		if line == "" {
			log.Println(lineno, " Found an empty line, skipping.")
			continue
		}

		//log.Println("Building word: ", line)
		// build word graph here?
		//wg := &LetterNode{}
		wg := CreateWordGraph(line)

		//log.Println(wg.Show())

		words[line] = Word{
			Word:  line,
			Runes: []rune(line),
			Graph: wg,
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	log.Println(path, " File read successfully!")

	d = Dictionary{
		Name:  name,
		Words: words,
	}
	return d
}
