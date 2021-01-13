/*
Package scrabble implements utility routines
 useful when dealing with the Scrabble game.
*/
package scrabble

import (
                "unicode"
        )

// Impl. note: const map.
var letterValueMap = func(letter rune) int { return 0 }
func init() {
        scoringData := map[uint8]string{
                         1: "AEIOU" +
                            "LNRST",
                         2: "DG",
                         3: "BCMP",
                         4: "FHVWY",
                         5: "K",
                         8: "JX",
                        10: "QZ",
                }
        lValue := map[rune]uint8{}
        lValueInit := func() {
                        for score, letters := range scoringData {
                                for _, letter := range letters {
                                    lValue[letter] = score
                                }
                        }
                }
        lValueInit()
        letterValueMap = func(letter rune) int {
                        return int(lValue[letter])
                }
}

// Score returns the Scrabble score for word.
func Score(word string) int {
        var sum int = 0
        for _, c := range word {
                sum += letterValueMap(unicode.ToUpper(c))
        }
        return sum
}
