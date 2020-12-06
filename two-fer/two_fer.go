/*
Package twofer implements "two for one".
*/
package twofer

import ( "fmt" )

const twoForOneFormat = "One for %s, one for me."

// ShareWith returns "One for name, one for me",
//  with name equals to "you" if the argument name is empty.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf(twoForOneFormat, name)
}
