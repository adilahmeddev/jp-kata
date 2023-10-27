package jpservice

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode/utf16"
)

type LanguageProcessor interface {
	GetUnique(characters string) []string
}

type JpProcessor struct {
}

func (jpProcessor *JpProcessor) GetUnique(characters string) []string {
	u := utf16.Decode([]uint16{0x4e00})
	y := utf16.Decode([]uint16{0x9fcf})
	r, err := regexp.Compile("[" + string(u) + "-" + string(y) + "]")
	if err != nil {
		log.Fatal(err.Error())
	}

	characterSet := map[string]struct{}{}

	characterArray := strings.Split(characters, "")
	uniqueChars := []string{}

	for i := 0; i < len(characterArray); i++ {
		fmt.Println("YO DAWG _ " + characterArray[i])
		if r.MatchString(characterArray[i]) {
			_, exists := characterSet[characterArray[i]]
			fmt.Println("YYOMAMA _ " + characterArray[i])
			if !exists {
				uniqueChars = append(uniqueChars, characterArray[i])
				characterSet[characterArray[i]] = struct{}{}
				fmt.Println("z _ " + characterArray[i])
			}
		}
	}

	return uniqueChars
}
