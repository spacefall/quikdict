package modes

import (
	"fmt"
	"quikdict/utils"
	"strings"
)

func PrintDictionary(info utils.WordInfo, tw int) {
	bold.Print(info.Word)
	// print phonetics in a dark color with brackets and the phonetics in italic
	if len(info.Phonetics) > 0 {
		dark.Printf(" [%s]\n", strings.Join(info.Phonetics, ", "))
	} else {
		println()
	}

	// print meaning
	for _, meaning := range info.Meanings {
		// print part of speech in green
		figureColor.Printf("%s\n", meaning.PartOfSpeech)

		defLen := len(meaning.Definitions)
		for i, def := range meaning.Definitions {

			// print definition
			if i != defLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", continueBox, utils.WordWrap(def.Definition, tw-4-margin, darkLine))
				if def.Example != "" {
					darkItalic.Printf("%s %s %s\n", line, endBox, utils.WordWrap(def.Example, tw-8-margin, line+s(4)))
				}

			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", darkEndBox, utils.WordWrap(def.Definition, tw-4-margin, s(4)))
				if def.Example != "" {
					darkItalic.Printf("%s%s %s\n", s(5), endBox, utils.WordWrap(def.Example, tw-8-margin, s(8)))
				}
			}
		}
		// new line to better separate figure of speech
		println()
	}
	dark.Printf("Source: %s\n", strings.Join(info.Sources, ", "))
}
