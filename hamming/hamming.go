/*
Package hamming implements a routine
 calculating, given 2 strands of DNA,
 the Hamming Distance between them.
*/
package hamming

import (
		"errors"
	)

func zip(s1, s2 string) ([][2]string, error) {

        if len(s1) != len(s2) {
                return nil, errors.New("strings not of equal length")
        }

        zipS := make([][2]string, len(s1))
        for i := range zipS {
                zipS[i] = [2]string{string(s1[i]), string(s2[i])}
        }

        return zipS, nil
}

func countIf(bools []bool, countCond func(bool) bool) int {
        c := 0
        for _, v := range bools {
                if countCond(v) {
                        c++
                }
        }
        return c
}

func nFalse(bools []bool) int {
        return countIf(bools, func(b bool) bool { return b == false })
}

// Distance returns,
//  given 2 strings representing 2 strands of DNA,
//  the Hamming Distance between them.
func Distance(a, b string) (int, error) {

	zipAB, err := zip(a, b)
        if err != nil {
               return -1, errors.New("sequences not of equal length")
        }

	zIdent := make([]bool, len(zipAB))
	for i := range a {
		zIdent[i] = (zipAB[i][0] == zipAB[i][1])
	}

	return nFalse(zIdent), nil;
}
