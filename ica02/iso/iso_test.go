package iso

import ("testing"
		)

func Test_GreetingExtendedASCII(t *testing.T) {
	str := []byte(GreetingExtendedASCII())
	for i := 0; i < len(str) ; i++ {
		if !(str[i] < 255 && str[i] > 127) {
			t.Errorf("%q, %v", "GreetingASCII want < 255 and > 127, got", str[i])
		}
	}
}
