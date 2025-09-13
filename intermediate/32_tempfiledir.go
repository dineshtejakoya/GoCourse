package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)
	fmt.Println("Temporary file created:", tempFile.Name())
	checkError(err)

	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkError(err)

	//defer os.Remove(tempDir)
	defer os.RemoveAll(tempDir)
	fmt.Println("Temporary Directory created:", tempDir)
}
