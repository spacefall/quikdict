package modes

import (
	"fmt"
	"quikdict/utils"
)

func PrintThesaurus(info utils.WordInfo) {
	printed := false

	for _, meaning := range info.Meanings {
		if len(meaning.Synonyms) == 0 {
			continue
		} else if len(meaning.Synonyms) > 0 && !printed {
			//goland:noinspection GoUnhandledErrorResult
			bold.Println("Synonyms")
			printed = true
		}

		// print part of speech in green
		//goland:noinspection GoUnhandledErrorResult
		figureColor.Printf("%s\n", meaning.PartOfSpeech)

		synLen := len(meaning.Synonyms)
		for i, syn := range meaning.Synonyms {
			// print definition
			if i != synLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("├─"), syn)
			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("╰─"), syn)
			}
		}
		// new line to better separate figure of speech
		fmt.Print("\n")
	}

	printed = false
	for _, meaning := range info.Meanings {
		if len(meaning.Antonyms) == 0 {
			continue
		} else if len(meaning.Antonyms) > 0 && !printed {
			//goland:noinspection GoUnhandledErrorResult
			bold.Println("Antonyms")
			printed = true
		}

		// print part of speech in green
		//goland:noinspection GoUnhandledErrorResult
		figureColor.Printf("%s\n", meaning.PartOfSpeech)

		antLen := len(meaning.Antonyms)
		for i, ant := range meaning.Antonyms {
			// print definition
			if i != antLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("├─"), ant)
			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("╰─"), ant)
			}
		}
		// new line to better separate figure of speech
		fmt.Print("\n")
	}
}
