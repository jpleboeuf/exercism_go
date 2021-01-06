/*
Package raindrops implements a converter
 between numbers and raindrop sounds.
*/
package raindrops

import (
        "strings"
        "strconv"
)

// rsConvTblItem represents
//  an item in the conversion table
//  in which each line associates
//  a factor and a raindrop sound.
type rsConvTblItem struct {
        factor int
        sound  string
}

// hasFactor returns if n has f for factor.
func hasFactor(n, f int) bool {
        return n % f == 0
}

// Convert returns
//  the raindrop sounds associated with number,
//  or the number itself
//   if there is no raindrop sound
//   associated with this number.
func Convert(number int) string {
        rsConvTbl := []rsConvTblItem{
                        rsConvTblItem{3, "Pling"},
                        rsConvTblItem{5, "Plang"},
                        rsConvTblItem{7, "Plong"},
                }
        raindropSounds := make([]string, 0, len(rsConvTbl))
        for _, rscti := range rsConvTbl {
                if hasFactor(number, rscti.factor) {
                    raindropSounds = append(raindropSounds, rscti.sound)
                }
        }
        if len(raindropSounds) == 0 {
                return strconv.Itoa(number)
        }
        return strings.Join(raindropSounds, "")
}
