package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "birthday_001.txt"
	// => Birthday - 1 of 4.txt
	newName, err := match(fileName, 4)
	if err != nil {
		fmt.Println("no match")
		os.Exit(1)
	}
	fmt.Println(newName)
}
// match returns the new filenam, or an error if the file name
// didn't match our pattern
func match(fileName string, total int) (string, error) {
	// "birthday", "001", "txt"
	pieces := strings.Split(fileName, ".")
	// getting extention "txt". In case there are more dots in the 
	// filename, we use the last item in the array to make sure we
	// get the extension
	ext := pieces[len(pieces)-1]
	// in case there are more dots, we have to get these pieces
	// back together. For example, "birth.day" would be ["birth", "day"],
	// and we don't want to break up the words
	tmpFileName := strings.Join(pieces[0:len(pieces)-1], ".")
	// same steps as above, but now we want the "001" part
	pieces = strings.Split(tmpFileName, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", fileName)
	}

	// Birthday - 1.txt
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}