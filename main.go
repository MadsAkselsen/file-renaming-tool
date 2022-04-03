package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
}

func main() {
	dir := "sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, file{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})
	for _, f := range toRename {
		fmt.Printf("%q\n", f)
	}
	for _, orig := range toRename {
		var n file
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			fmt.Println("Error renaming:", orig.path, err.Error())
		}
	}
}
// match returns the new filename, or an error if the file name
// didn't match our pattern
func match(filename string) (string, error) {
	// "birthday", "001", "txt"
	pieces := strings.Split(filename, ".")
	// getting extention "txt". In case there are more dots in the 
	// filename, we use the last item in the array to make sure we
	// get the extension
	ext := pieces[len(pieces)-1]
	// in case there are more dots, we have to get these pieces
	// back together. For example, "birth.day" would be ["birth", "day"],
	// and we don't want to break up the words
	tmpfilename := strings.Join(pieces[0:len(pieces)-1], ".")
	// same steps as above, but now we want the "001" part
	pieces = strings.Split(tmpfilename, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", filename)
	}

	// Birthday - 1.txt
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}