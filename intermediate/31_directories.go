package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

// instead of typing err!=nil everything lets keep it in a function
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//this is one way
	// err := os.Mkdir("subdir",0755)
	// checkError(err)

	//checkError(os.Mkdir("subdir1", 0755))

	//we can remove the files that are created using os.RemoveAll where we need to pass directories as strings
	//defer os.RemoveAll("subdir2")

	//byteslice we are keeping it empty to create an empty file
	//os.WriteFile("subdir1/file", []byte(""), 0755)

	//MkdirAll creates multiple directories at once
	//mkdirAll is designed to not to fail if the directory already exists, but mkdir will return if file already exist
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))

	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	//result is a slice

	for _, entry := range result {
		//here entry variable also has some methods associated with it
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	//now when os.readDir is subdir/parent the cursor of os is in that path only
	//in order to move the cursor to different directory we use chdir

	checkError(os.Chdir("subdir/parent/child"))
	//we use . because the we want to read/traverse from the current directory
	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result {
		fmt.Println(entry)
	}

	fmt.Println("Go three levels above")
	//Go three levels above the current path
	checkError(os.Chdir("../../.."))
	//to get the os current working directory
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	//filepath.Walk and filepath.WalkDir
	//WalkDir is preferred due to performance reasons
	//WalkDir uses os.Directory entry under the hood
	//filepath.Walk uses os.fileinfo which provides more details but maybe less efficient
	pathfile := "subdir/"
	//try out to see all contents in gocourse
	//pathfile := "."

	fmt.Println("Walking Directory")
	//walkDir function takes another function as its argumetns
	//type WalkDirFunc func(path string, d DirEntry, err error) error
	//for DirEntry we need to use os.DirEntry as it is part of os package
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(path)
		return nil
	})

	checkError(err)

	//remove all the directories we created
	//checkError(os.RemoveAll("./subdir"))
	// checkError(os.RemoveAll("./subdir"))

	//if we use only os.Remove() and if the file is not empty it would return error
	//in this case to delete contents in directories use removeall

}
