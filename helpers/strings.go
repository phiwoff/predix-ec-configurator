package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReplaceAll ...
func ReplaceAll(filename string, old []string, new []string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	newContent := string(content)
	for index, element := range old {
		newContent = strings.Replace(string(newContent), element, new[index], -1)
	}
	return newContent
}

// GetStringAsInt ...
func GetStringAsInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return result
}
