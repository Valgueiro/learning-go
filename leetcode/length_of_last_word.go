package main

import "fmt"

func main() {
	str := "   fly me   to   the moon  "
	for _, s := range filter(split(str), isNotEmptyStringOrHasAlphabeticChars) {
		fmt.Println(s)
	}
	fmt.Println(lengthOfLastWord(str))
}

func isNotEmptyStringOrHasAlphabeticChars(s string) bool {
	for _, c := range s {
		if c != ' ' {
			return true
		}
	}

	return false
}

func lengthOfLastWord(s string) int {
	splitString := split(s)
	filteredSplitStrings := filter(splitString, isNotEmptyStringOrHasAlphabeticChars)
	lastWord := filteredSplitStrings[len(filteredSplitStrings)-1]

	count := 0
	for _, c := range lastWord {
		if c != ' ' {
			count++
		}
	}
	return count
}

func split(s string) []string {
	out := []string{}
	lastSpaceIndex := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || i == len(s)-1 {
			out = append(out, s[lastSpaceIndex+1:i+1])
			lastSpaceIndex = i
		}
	}

	return out
}

func filter(strings []string, filterFunc func(string) bool) []string {
	out := []string{}
	for _, s := range strings {
		if filterFunc(s) {
			out = append(out, s)
		}
	}
	return out
}
