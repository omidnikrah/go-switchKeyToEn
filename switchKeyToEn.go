package switchKeyToEn

import (
	"fmt"
	"strings"
)

var (
	persianChar = []string{"ض", "ص", "ث", "ق", "ف", "غ", "ع", "ه", "خ", "ح", "ج", "چ", "ش", "س", "ی", "ب", "ل", "ا", "ت", "ن", "م", "ک", "گ", "ظ", "ط", "ز", "ر", "ذ", "د", "پ", "و", "؟"}
	englishChar = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "[", "]", "a", "s", "d", "f", "g", "h", "j", "k", "l", ";", "'", "z", "x", "c", "v", "b", "n", "m", ",", "?"}
)

func switchKeyToEn(text string) {
	for i := 0; i <= len(persianChar)-1; i++ {
		text = strings.ReplaceAll(text, persianChar[i], englishChar[i])
	}
	fmt.Printf("%s", text)
}
