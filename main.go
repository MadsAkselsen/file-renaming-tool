package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// fileName := "birthday_001.txt"
	// // => Birthday - 1 of 4.txt
	// newName, err := match(fileName, 4)
	// if err != nil {
	// 	fmt.Println("no match")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)

	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	count := 0
	// type rename struct {
	// 	filename 	string
	// 	path		string
	// }
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Dir: ", file.Name())
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}

	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
	}

	// origPath := fmt.Sprintf("%s/%s")
	// newPath := fmt.Sprintf("%s/%s")
}
// match returns the new filename, or an error if the file name
// didn't match our pattern
func match(filename string, total int) (string, error) {
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
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}