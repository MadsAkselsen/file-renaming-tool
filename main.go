package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this should be a real or dry run")
	flag.Parse()
	
	walkDir := "sample"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		currentDir := filepath.Dir(path)
		if matchResult, err := match(info.Name()); err == nil {
			key := filepath.Join(currentDir, fmt.Sprintf("%s.%s", matchResult.base, matchResult.ext))
			toRename[key] = append(toRename[key], info.Name())
		}
		return nil
	})
	
	for key, files := range toRename {
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for i, filename := range files {
			res, _ := match(filename)
			newFilename := fmt.Sprintf("%s - %d of %d.%s", res.base, (i+1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFilename)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			
			if !dry {
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error renaming:", oldPath, newPath, err.Error())
				}
			}
		}
	}
}

type matchResult struct {
	base string
	index int
	ext string
}

// Match returns the new filename, or an error if the file name
// didn't match our pattern
func match(filename string) (*matchResult, error) {
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
		return nil, fmt.Errorf("%s didn't match our pattern", filename)
	}
	return &matchResult{strings.Title(name), number, ext}, nil
}