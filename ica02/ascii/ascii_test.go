package ascii

import (
		"unicode"
		"testing"
	)

func Test_greetingASCII(t *testing.T) {
	toStr := []byte(greetingASCII())
	for i := 0; i < len(toStr) ; i++ {
		if toStr[i] >= unicode.MaxASCII {
			t.Errorf("%q%v\"", "GreetingASCII want < 127, got ", toStr[i])
		}
	}
}
