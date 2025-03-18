package modes

import (
	"fmt"
	"quikdict/utils"
	"strings"
)

func PrintThesaurus(info utils.WordInfo, tw int) {
	printed := false

	for _, meaning := range info.Meanings {
		if len(meaning.Synonyms) == 0 {
			continue
		} else if len(meaning.Synonyms) > 0 && !printed {
			bold.Println("Synonyms")
			printed = true
		}

		// print part of speech in green
		figureColor.Printf("%s\n", meaning.PartOfSpeech)

		synLen := len(meaning.Synonyms)
		for i, syn := range meaning.Synonyms {
			// print definition
			if i != synLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", continueBox, utils.WordWrap(syn, tw-4-margin, dark.Sprint(line)))
			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", darkEndBox, utils.WordWrap(syn, tw-4-margin, s(4)))
			}
		}
		// new line to better separate figure of speech
		println()
	}

	printed = false
	for _, meaning := range info.Meanings {
		if len(meaning.Antonyms) == 0 {
			continue
		} else if len(meaning.Antonyms) > 0 && !printed {
			bold.Println("Antonyms")
			printed = true
		}

		// print part of speech in green
		//goland:noinspection GoUnhandledErrorResult
		figureColor.Printf("%s\n", meaning.PartOfSpeech)

		antLen := len(meaning.Antonyms)
		for i, ant := range meaning.Antonyms {
			if i != antLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", continueBox, utils.WordWrap(ant, tw-4-margin, darkLine))
			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", darkEndBox, utils.WordWrap(ant, tw-4-margin, s(4)))
			}
		}
		// new line to better separate figure of speech
		println()
	}
	dark.Printf("Source: %s\n", strings.Join(info.Sources, ", "))
}
