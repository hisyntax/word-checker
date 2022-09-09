package validate

import (
	"strings"
)

var Words []string

//default validator validates a string of words from a predefined array of words
func DefaultValidator(message string) bool {
	Words = []string{
		"arse", "ass", "bastard", "bitch",
		"bloody", "bollocks", "bugger", "bullshit",
		"cock", "crap", "cunt", "damn", "dick", "dyke",
		"frigger", "fuck", "hell", "shit", "kike", "nigra",
		"piss", "prick", "pussy", "shite", "slut", "spastic",
		"twat", "wanker", "whore",
	}
	var wordMatch bool

	for _, v := range Words {
		if strings.Contains(strings.ToLower(message), v) {
			wordMatch = true
		}
	}

	return wordMatch
}

//custom validator validates a string of words from a defined custom array of words
func CustomValidator(customWords []string, message string) bool {
	var wordMatch bool
	for _, v := range customWords {
		if strings.Contains(strings.ToLower(message), v) {
			wordMatch = true
		}
	}

	return wordMatch
}
