/*
Package acronym implements utility routines
 generating acronyms.
*/
package acronym

import (
		"strings"
		"unicode"
	)

func filter(s []string, test func(string) bool) (sFiltered []string) {
	for _, se := range s {
		if test(se) {
			sFiltered = append(sFiltered, se)
		}
	}
	return
}

func mapf(s []string, mapper func(string) string) (sMapped []string) {
	for _, se := range s {
		sMapped = append(sMapped, mapper(se))
	}
	return
}

// Abbreviate returns the acronym corresponding to string s.
func Abbreviate(s string) string {
	nonWordChars := func(c rune) bool {
			return !unicode.IsLetter(c) && (c != '\'') && !unicode.IsNumber(c)
		}
	words := filter(strings.FieldsFunc(s, nonWordChars), func(ss string) bool { return ss != "" })
	initials := mapf(words, func(ss string) string { return strings.ToUpper(ss[:1]) })
	return strings.Join(initials, "")
}
