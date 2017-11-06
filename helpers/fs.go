package helpers

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/mholt/archiver"
)

// IsExist ...
func IsExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// exist
		return true
	}
	// not exist
	return false
}

// CreateDirectory ...
func CreateDirectory(directoryPath string) {
	//choose your permissions well
	err := os.MkdirAll(directoryPath, 0777)
	//check if you need to panic, fallback or report
	if IsError(err) {
		return
	}
}

// WriteStringToFile ...
func WriteStringToFile(filepath, s string, executable bool) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()
	if executable == true {
		os.Chmod(filepath, 0755)
	}

	_, err = io.Copy(fo, strings.NewReader(s))
	if IsError(err) {
		return nil
	}
	return nil
}

// ExtractArchiveFileToFolder ...
func ExtractArchiveFileToFolder(pathToFile string, outputFolder string) {
	log.Printf("Extracting %s to %s", pathToFile, outputFolder)
	err := archiver.TarGz.Open(pathToFile, outputFolder)
	if IsError(err) {
		log.Panicf("ERROR extracting %s", pathToFile)
	}
	log.Printf("==> Done extracting %s to %s", pathToFile, outputFolder)
}
