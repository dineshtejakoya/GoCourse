package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//buffering means it reads and writes data in chunks
	//it can be used in chat s/w or streaming platfor or audio playing platform
	//transferring datacommunicating between two ends, utilizing small chunks of data

	//initiate a new Reader
	//bufio.NewReader accepts other readers as input
	//bufio.NewReader is a wrapper around an existing reader
	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageeeee!\nHello bufio package"))
	//fmt.Println(reader)

	//Reading Byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	//we can start reading after 20 bytes until new line character
	fmt.Println("Read string: ", line)

	fmt.Println("=====bufio Writer=======")
	//bufio.Writer is a struct that wraps around an io.writer and provides buffering for efficient writing of data

	writer := bufio.NewWriter(os.Stdout)
	//os.Stdout is a writer because it impletements write method and io.writerinterface tells that any
	//Writing byte slice
	data1 := []byte("Hello, bufio package!\n")
	n1, err := writer.Write(data1)
	if err != nil {
		fmt.Println("Error Writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n1)
	//to be able to get the output onto the console we have to flush the buffer
	//what happens to bufio.writer is all the data that is written to the writer is stored in an internal buffer and it is not written immediately to the os.stdout or any other writer that we pass as an argument

	//Flush the buffer to ensure all data is written to os.stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	//Writing String
	str := "This is a string\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	//Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	//It seems there is not much difference between Write and WriteString
	//Write takes bytes slice , for string we take writestring
}
