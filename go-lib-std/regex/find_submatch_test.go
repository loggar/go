package regex_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubmatch(t *testing.T) {
	str := "abc@def.com is the mail address of 8th user with id 124"
	exp := regexp.MustCompile(
		`([a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,})` +
			`|(email|mail)` +
			`|(\d{1,3})`,
	)

	match := exp.FindStringSubmatch(str)
	// log.Println(match) //[abc@def.com abc@def.com  ]
	matches := exp.FindAllStringSubmatch(str, -1)
	// log.Println(matches) // [[abc@def.com abc@def.com  ] [mail  mail ] [8   8] [124   124]]

	assert.Equal(t, 4, len(match), "found sub-match")
	assert.Equal(t, 4, len(matches), "found sub-matches")
}

func TestFindSubmatch_MatchInsideMatch(t *testing.T) {
	str := "abe21@def.com is the mail address of 8th user with id 124"
	exp := regexp.MustCompile(
		`([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,})` +
			`|(mail|email)` +
			`|(\d{1,3})`,
	)

	match := exp.FindStringSubmatch(str)
	// log.Println(match) // [abe21@def.com abe21@def.com 21   ]
	matches := exp.FindAllStringSubmatch(str, -1)
	// log.Println(matches) // [[abe21@def.com abe21@def.com 21   ] [mail    mail ] [8     8] [124     124]]

	// there are a few empty string sub-matches because the second sub-match for the domain name number returns an empty string.
	/*
		for i, s := range match {
			log.Printf("match %d: %q", i, s)
		}
		// match 0: "abe21@def.com"
		// match 1: "abe21@def.com"
		// match 2: "21"
		// match 3: ""
		// match 4: ""
		// match 5: ""
	*/

	assert.Equal(t, 6, len(match), "found sub-match")
	assert.Equal(t, 4, len(matches), "found sub-matches")
}
