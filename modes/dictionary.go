package modes

import (
	"fmt"
	"quikdict/utils"
	"strings"
)

func PrintDictionary(info utils.WordInfo) {
	//goland:noinspection GoUnhandledErrorResult
	bold.Print(info.Word)
	// print phonetics in a dark color with brackets and the phonetics in italic
	if len(info.Phonetics) > 0 {
		//goland:noinspection GoUnhandledErrorResult
		dark.Printf(" [%s]\n", strings.Join(info.Phonetics, ", "))
	} else {
		fmt.Print("\n")
	}

	// print meaning
	for _, meaning := range info.Meanings {
		// print part of speech in green
		//goland:noinspection GoUnhandledErrorResult
		figureColor.Printf("%s\n", meaning.PartOfSpeech)

		defLen := len(meaning.Definitions)
		for i, def := range meaning.Definitions {
			// print definition
			if i != defLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("├─"), def.Definition)
				if def.Example != "" {
					//goland:noinspection GoUnhandledErrorResult
					darkItalic.Printf(" │   ╰─ %s\n", def.Example)
				}
			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("╰─"), def.Definition)
				if def.Example != "" {
					//goland:noinspection GoUnhandledErrorResult
					darkItalic.Printf("     ╰─ %s\n", def.Example)
				}
			}
		}
		// new line to better separate figure of speech
		fmt.Print("\n")
	}
}
