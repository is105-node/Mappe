package frequence
	
import ("testing"
		"io/ioutil")

func BenchmarkCountRunes(b *testing.B) {
        file, _ := ioutil.ReadFile("text1.txt")
        fl:=string(file)
        for i := 0; i < b.N; i++ {
        	CountRunes(fl)
    }
}

func BenchmarkCountLines(b *testing.B){
	file, _ := ioutil.ReadFile("text1.txt")
        for i := 0; i < b.N; i++ {
        	CountLines(file)
    }
}

func BenchmarkWriteFile10(b *testing.B){
	file, _ := ioutil.ReadFile("text1.txt")
	fl:=string(file)
        for i := 0; i < b.N; i++ {
        	writeFile("fnameTest.txt", fl)
    }
}
