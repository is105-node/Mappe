package main
import (
	"net/http"
	"net/url"
)
import "io/ioutil"
import "fmt"
import "os"
import "strings"
import "bufio"

func main() {
	for {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	st := strings.Fields(text)
	resp, err := http.PostForm("http://localhost:8080/", url.Values{"Name": {string(st[0])}, "Email":{string(st[1])}})
	if err != nil {
		panic(err)
	}
	
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", html)
	
}
	

}