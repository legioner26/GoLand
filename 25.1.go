package main

import (
	"flag"
	"fmt"
)

func main() {
	var str, substr string

	flag.StringVar(&str, "str", "привет мир!", "set string")
	flag.StringVar(&substr, "substr", "вет ", "set string fot search")
	flag.Parse()

	fmt.Println("str:", str)
	if isSubstr(&substr, &str) {
		fmt.Printf("substr: \"%s\" найдена \n", substr)
	} else {
		fmt.Printf("substr: \"%s\" не найдена \n", substr)
	}
}

func isSubstr(substr *string, str *string) bool {
	substrRunes := []rune(*substr)
	strRunes := []rune(*str)

	count := 0
	lenSubstrRunes := len(substrRunes)
	for _, r := range strRunes {

		if r == substrRunes[count] {
			count++
			if count == lenSubstrRunes {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
