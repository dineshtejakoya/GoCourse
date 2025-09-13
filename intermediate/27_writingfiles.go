package intermediate

import (
	"fmt"
	"os"
)

// when writing large amounts of data, consider using bufferedwriters for better performance
func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	//Write data to file
	data := []byte("Hello world!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Data has been written to file successfully")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Dinesh Teja Koya!\n")
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Writing to writeString.txt completed")
}
