package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// locating, referencing files and directing within filesystem
// Absolute path specifies the complete path from the root directory (C:\\)
// Relative path is in position with the current working directory or .. to go above the current working directory
// path separators, in unix like systems the paths are typically separated by //
// windows paths use \\
// Go's file path/filepath package handles both styles seamlessly
func main() {
	relativePath := "./data/file.txt"
	//as we are in windows we are using c: drive like path which will help in isabs() function
	absolutePath := `C:\Users\me\file.txt`

	//join paths using filepath.join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	//remove the redundant path
	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	//Base generates last component fo the filepath
	fmt.Println(filepath.Base("/home/user/docs/file.txt"))

	fmt.Println("Is relative path variable absolute", filepath.IsAbs(relativePath))
	fmt.Println("Is absolute path variable absolute", filepath.IsAbs(absolutePath))

	//we can extract the extension of a file name
	fmt.Println("Extension of file", filepath.Ext(file))
	//Trimsuffix takes two arguments filename and extension of file
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	//we want to find a path that is relative to another path
	//if we are inside b folder we need to go to t folder and access
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	//if you see in both paths b and c are the childs of a
	//so right now we are in c directory and to move one level above ..
	//and then go to b/t/fil
	//therefore ..\b\t\file
	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	//Convert a relative path to an absolute path
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Absolute Path: ", absPath)
	}
}
