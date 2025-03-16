package cli

import (
	"errors"
	"fmt"
	"os"
	"quikdict/modes"
	"quikdict/utils"
)

func Parse(args []string) (Params, error) {
	var params Params
	params.Mode = modes.ModeDictionary
	skip := false
	for i, arg := range args {
		if skip {
			skip = false
			continue
		}
		switch arg {
		case "-tr", "--translate":
			if len(args) == i+1 || len(args[i+1]) != 2 || !utils.IsLetter(args[i+1]) {
				return Params{}, errors.New("invalid language code")
			}
			params.Translation = args[i+1]
			params.Mode = modes.ModeTranslate
			skip = true

		case "-th", "--thesaurus":
			params.Mode = modes.ModeThesaurus

		case "-h", "--help":
			fmt.Println(HelpStr)
			os.Exit(0)

		default:
			params.Word = arg
		}
	}
	return params, nil
}
