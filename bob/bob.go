/*
Package bob implements the conversation abilities of Bob,
 a lackadaisical teenager.
*/
package bob

import (
	"strings"
	"unicode"
)

// MsgKind represents the kinds of a message/remark.
// Kinds can be combined (happens only for Asking+Yelling).
type MsgKind int
const ( 
	mkA = (1 << (iota + 1))  // Asking     10 = 2
	mkY                      // Yelling   100 = 4
	mkN                      // Nothing  1000 = 8
)

// ContainsAlpha returns whether the string s contains at least one alpha character.
func ContainsAlpha(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true 
		}
	}
	return false
}

// IsUpper returns whether the string s is uppercase.
func IsUpper(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) && !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

func msgKind(msg string) MsgKind {
	var msgKind MsgKind = 0
	var msgTrimmed = strings.TrimSpace(msg)
	if msgTrimmed == "" {
		msgKind |= mkN
	} else {
		if strings.HasSuffix(msgTrimmed, "?") {
			msgKind |= mkA
		}
		if ContainsAlpha(msgTrimmed) && IsUpper(msgTrimmed) {
			msgKind |= mkY
		}
	}
	return msgKind
}

// const map:
var replyToMap = func() map[string]string {
		return map[string]string{
				"asking_question":  "Sure.",
				"yelling":          "Whoa, chill out!",
				"yelling_question": "Calm down, I know what I'm doing!",
				"saying_nothing":   "Fine. Be that way!",
				"default":          "Whatever.",
			}
	}
func replyTo(msgKey string) string {
	return replyToMap()[msgKey]
}

// Hey returns Bob's answer to a remark made to him.
func Hey(remark string) string {
	switch msgKind(remark) {
		case mkA:
			return replyTo("asking_question")
		case mkY:
			return replyTo("yelling")
		case mkA | mkY:
			return replyTo("yelling_question")
		case mkN:
			return replyTo("saying_nothing")
		default:
			return replyTo("default")
	}
}
