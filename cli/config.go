package cli

import (
	"fmt"
	"os"
	"quikdict/modes"
)

var HelpStr = fmt.Sprintf(`Usage: %s [word] [options]
Options:
  -th, --thesaurus    Print synonyms and antonyms
  -tr, --translate    Translate the word to another language
  -h, --help          Print this help message`, os.Args[0])

type Params struct {
	Word        string
	Mode        modes.Mode
	Translation string
}
