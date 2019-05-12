package main

import (
	"fmt"
	"compress/gzip"
	"io"
	"io/ioutil"
	"bytes"
	"log"
)

// Write gzipped data to a Writer
func gzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	return err
}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}

func main() {
	s := "some data"
	fmt.Println("original:\t", s, len(s))
	
	var buf bytes.Buffer
	err := gzipWrite(&buf, []byte(s))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("compressed:\t", buf.String(), len(buf.String()))
	
	var buf2 bytes.Buffer
	err = gunzipWrite(&buf2, buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("decompressed:\t", buf2.String(), len(buf2.String()))
}