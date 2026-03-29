package core

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/AhmedCharfeddine/haraka/core/mapping"
)

const MappingStrategy = "mapping"

func Transliterate(latinWord string, strategy string) (arabicWord string, e error) {
	if len(latinWord) < 3 {
		return latinWord, fmt.Errorf("word has to be at least three characters long")
	}
	if strategy == MappingStrategy {
		return mapping.TransliterateMapping(latinWord), nil
	}
	return latinWord, fmt.Errorf("unknown transliteration strategy")
}

func TransliterateParagraph(latinParagraph string, strategy string) (arabicParagraph string, e error) {
	var output strings.Builder
	var wordBuilder strings.Builder

	for _, r := range latinParagraph {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			wordBuilder.WriteRune(r)
		} else {
			// End of a word
			if wordBuilder.Len() > 0 {
				word := wordBuilder.String()
				arabicWord, err := Transliterate(word, strategy)
				if err != nil {
					return "", err
				}
				output.WriteString(arabicWord)
				wordBuilder.Reset()
			}
			// Add punctuation/whitespace as-is
			output.WriteRune(r)
		}
	}

	// Final word at the end
	if wordBuilder.Len() > 0 {
		word := wordBuilder.String()
		arabicWord, err := Transliterate(word, MappingStrategy)
		if err != nil {
			return "", err
		}
		output.WriteString(arabicWord)
	}

	return output.String(), nil
}
