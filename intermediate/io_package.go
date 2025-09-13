package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

//io package is standard library package in go for input and output operations, it provides interfaces and functions for reading and writing data
//it facilitates interactions with various data sources,like files, networks or in memory buffers etc.,
//core interfaces: io.Reader, io.Writer, io.Closer

func readFromReader(r io.Reader) {
	//1024 bytes
	buf := make([]byte, 1024)
	//r.Read returns number of bytes
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error Reading from reader", err)
	}
	//buf[:n] this buffer will include index 0 to n-1
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
	}
}

// we can close fs.file,os.file any other resource that is using it, in this common function
func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
}

func main() {

}
