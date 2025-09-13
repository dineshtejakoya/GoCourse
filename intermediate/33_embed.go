package intermediate

//if we are not using something directly but indirectly use _
import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

// Embed directive is to embed static files or directories into go binaries at build time
// this directive provides convenient and efficient way to include assets directly with in your go programs eliminating the need to manage these assets separately
//
//go:embed output.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

	fmt.Println("Embedded content:", content)
	content, err := basicsFolder.ReadFile("basics/hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Embedded file content:", string(content))

	//Now we have access to the basics folder, we'll walk through the folder
	//earlier in DirEntry we used os, but we fs also imported we can use it as well
	//one advantage to use fs is it is cross platform independent
	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

//Note : embedded files cannot be modified at runtime
//we need to rebuild the binary for any updates
