package ascii

import ("fmt"
	)
//import "encoding/hex"

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// Oppgave 1b
// Implementer en funksjon som eksportere const ascii
func GetASCIIStringLiteral() string{
	return ascii;
}

// Funksjon tar en "string literal" med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
//
// Eksempel (utskriften kommer fra en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...

func IterateOverASCIIStringLiteral(stringLiteral string) {
	for x:= 0; x < len(stringLiteral); x++ {
		fmt.Printf("%X%6s%12b\n", stringLiteral[x], string(stringLiteral[x]), stringLiteral[x])
	}
}



// Unix-like operating systems are known to use it as erase control character, i.e. to delete the previous character in the line mode. 

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen GreetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1c
func greetingASCII() string {
	toStr := []byte("\"HellÆØoæø :-)\"")
	var retStr string
	for i := 0; i < len(toStr) ; i++ {
		fmt.Println(toStr[i])
		if toStr[i] <= 127 {
		retStr = retStr+string(toStr[i])
		}
	}
	fmt.Println(retStr)
	return retStr
}