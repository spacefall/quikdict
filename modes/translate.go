package modes

import "fmt"

func PrintTranslation(word, translation, origin, destination string) {
	fmt.Printf("%s %s ⟶ %s %s", word, dark.Sprintf("(%s)", origin), bold.Sprint(translation), dark.Sprintf("(%s)", destination))
}
