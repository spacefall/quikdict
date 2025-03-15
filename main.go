package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"quikdict/sources"
	"strings"
)

var figureColor = color.New(color.FgHiGreen, color.Bold)
var darkItalic = color.New(color.FgHiBlack, color.Italic)
var dark = color.New(color.FgHiBlack)

func main() {
	color.NoColor = false
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [word]", os.Args[0])
		handleErrStr("no word was provided")
	}

	word := os.Args[1]

	info, err := sources.GetFromDictionaryAPI(word)
	handleErr(err)

	// print word in bold
	_, _ = color.New(color.Bold).Print(word)

	// print phonetics in a dark color with brackets and the phonetics in italic
	if len(info.Phonetics) > 0 {
		_, _ = dark.Printf(" [%s]/n", darkItalic.Sprint(strings.Join(info.Phonetics, ", ")))
	}

	// print meaning
	for _, meaning := range info.Meanings {
		// print part of speech in green
		_, _ = figureColor.Printf("%s\n", meaning.PartOfSpeech)

		defLen := len(meaning.Definitions)
		for i, def := range meaning.Definitions {
			// print definition
			if i != defLen-1 {
				// print with continuing box characters if not the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("├─"), def.Definition)
				if def.Example != "" {
					_, _ = darkItalic.Printf(" │   ╰─ %s\n", def.Example)
				}
			} else {
				// print with ending box characters if the last definition
				fmt.Printf(" %s %s\n", dark.Sprint("╰─"), def.Definition)
				if def.Example != "" {
					_, _ = darkItalic.Printf("     ╰─ %s\n", def.Example)
				}
			}
		}
		// new line to better separate figure of speech
		fmt.Print("\n")
	}
}

func handleErrStr(err string) {
	fmt.Printf("%s %s\n", color.HiRedString("Error:"), err)
	os.Exit(1)
}

func handleErr(err error) {
	if err != nil {
		handleErrStr(err.Error())
	}
}
