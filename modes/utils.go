package modes

import (
	"github.com/fatih/color"
	"strings"
)

var figureColor = color.New(color.FgHiGreen, color.Bold)
var darkItalic = color.New(color.FgHiBlack, color.Italic)
var dark = color.New(color.FgHiBlack)
var bold = color.New(color.Bold)

var endBox = "╰─"
var line = " │  "
var continueBox = dark.Sprint("├─")
var darkEndBox = dark.Sprint(endBox)
var darkLine = dark.Sprint(line)

var margin = 1

func s(n int) string {
	return strings.Repeat(" ", n)
}
