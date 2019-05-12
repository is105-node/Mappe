// https://www.socketloop.com/references/golang-io-pipe-function-example
package pipe

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Reputf8(filename string) []byte {
	
	var fileinfo string
	fi, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%s %s \n %s", "Error opening the file", filename, err)
		fileinfo = ""
	} else {
			for x:=0;x<100;x++ {
			fileinfo = fileinfo + fmt.Sprintf("%X ", fi[x])	
		}
		
	}
	return []byte(fileinfo)

	
}

func Repb64(filename string) []byte {
	
	var b64ret string
	fi, err := ioutil.ReadFile(filename)

	fi = []byte(fi)
	
	if err != nil {
		fmt.Printf("%s %s \n %s", "Error opening the file", filename, err)
		b64ret = ""
	} else {
		
		b64 := base64.StdEncoding.EncodeToString(fi)
		for x:=0;x<100;x++ {
			b64ret = b64ret + fmt.Sprintf("%X ", b64[x])	
		}
		
	}
	return []byte(b64ret)
}

func Repgzip(filename string) {
	fi, _ := ioutil.ReadFile(filename)
	var st []byte
	for x:=0;x<len(fi);x++ {
		st = append(st, fi[x])
	}
	pReader, pWriter := io.Pipe()

	
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter)

	gWriter := gzip.NewWriter(b64Writer)

	go func() {
		gWriter.Write([]byte(st))
		
		gWriter.Close()
		b64Writer.Close()
		pWriter.Close()
	}()


	b64Reader := base64.NewDecoder(base64.StdEncoding, pReader)
	gReader, err := gzip.NewReader(b64Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	text, err := ioutil.ReadAll(gReader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i:=0;i<100;i++ {
		fmt.Printf("%X ", text[i])	
	}
	fmt.Printf("\nBytes Written: %v", len(text))

}

func Retmain(filebyte string) {

	// read and write with pipe
	pReader, pWriter := io.Pipe()
	// use base64 encoder
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter)
	gWriter := gzip.NewWriter(b64Writer)

	// write text to be gzipped and push to pipe
	go func() {
		fmt.Println("Start writing")
		n, err := gWriter.Write([]byte(filebyte))
		fmt.Printf("len = %d\n", n)

		if err != nil {
			fmt.Println(err)
			fmt.Println("F")
			os.Exit(1)
		}

		gWriter.Close()
		b64Writer.Close()
		pWriter.Close()
		fmt.Printf("Written %d bytes \n", n)
	}()

	// start reading from the pipe(reverse of the above process)

	// use base64 decoder to wrap pipe Reader
	b64Reader := base64.NewDecoder(base64.StdEncoding, pReader)

	// read gzipped text and decompressed the text
	gReader, err := gzip.NewReader(b64Reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Look at the final output at the other side of the pipe

	// print out the text
	text, err := ioutil.ReadAll(gReader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", text)

}