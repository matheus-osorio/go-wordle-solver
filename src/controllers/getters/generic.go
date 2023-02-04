package getters

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Struct that holds generic helper methods
type WordPicker struct {
	WordList []string
	WordSize uint
}

func (picker *WordPicker) SelectWordsBySize() {
	selectedWords := []string{}
	for _, name := range picker.WordList {
		if len(name) == int(picker.WordSize) {
			selectedWords = append(selectedWords, name)
		}
	}

	picker.WordList = selectedWords
}

func (picker *WordPicker) ReplaceLatinCharacters() {
	runeTransformer := runes.Remove(runes.In(unicode.Mn))
	transformer := transform.Chain(norm.NFD, runeTransformer, norm.NFC)
	newWordList := []string{}
	for _, word := range picker.WordList {
		newWord, _, _ := transform.String(transformer, word)

		newWordList = append(newWordList, newWord)
	}

	picker.WordList = newWordList
}
