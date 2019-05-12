package bfrequence
	
import "testing"
	

func BenchmarkCountRunes10(b *testing.B) {
        // run the Fib function b.N times
        for i := 0; i < b.N; i++ {
        	CountRunes("text1.txt")
    }
}

func BenchmarkCountLines10(b *testing.B){
	for i := 0; i < b.N; i++ {
		CountLines("text1.txt")
	}
}

func BenchmarkWriteFile10(b *testing.B){

        for i := 0; i < b.N; i++ {
        	WriteFile("fnameTest.txt", "Lines in the file: 124787\nRune: ' ' - 1285884 Instances\nRune: 'e' - 406157 Instances\nRune: 'o' - 282560 Instances\nRune: 'h' - 218875 Instances\nRune: 's' - 215605 Instances")
    }
}
