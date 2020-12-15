/*
Package hamming implements a routine
 calculating, given 2 strands of DNA,
 the Hamming Distance between them.
*/
package hamming

import (
		"errors"
	)

// Distance returns,
//  given 2 strings representing 2 strands of DNA,
//  the Hamming Distance between them.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("sequences not of equal length")
	}

	zipAB := make([][2]string, len(a))
	zIdent := make([]bool, len(zipAB))
	d := 0
	for i := range a {
		zipAB[i] = [2]string{string(a[i]), string(b[i])}
		zIdent[i] = (zipAB[i][0] == zipAB[i][1])
		if zIdent[i] == false {
			d++
		}
	}

	return d, nil;
}
