package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"quikdict/cli"
	"quikdict/sources"
	"strings"
)

var figureColor = color.New(color.FgHiGreen, color.Bold)
var darkItalic = color.New(color.FgHiBlack, color.Italic)
var dark = color.New(color.FgHiBlack)
var bold = color.New(color.Bold)

func main() {
	color.NoColor = false
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [word]", os.Args[0])
		handleErrStr("no word was provided")
	}

	args := cli.Parse(os.Args[1:])

	info, err := sources.GetFromDictionaryAPI(args.Word)
	handleErr(err)

	// Print synonyms and antonyms if thesaurus flag is set
	if args.Thesaurus {
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
	} else {
		//goland:noinspection GoUnhandledErrorResult
		bold.Print(args.Word)
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
