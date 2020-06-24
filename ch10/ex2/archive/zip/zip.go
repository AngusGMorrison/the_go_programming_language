// Package zip facilitates the extraction of FileInfo from zip files.
package zip

import (
	gozip "archive/zip"
	"os"

	"github.com/angusgmorrison/the_go_programming_language/ch10/ex2/archive"
)

func init() {
	archive.RegisterFormat("zip", "\x50\x4B\x03\x04", 0, Decode)
	archive.RegisterFormat("zip", "\x50\x4B\x05\x06", 0, Decode) // empty archive
}

// Decode extracts the FileInfo of all files found within the zip archive.
func Decode(file *os.File) ([]os.FileInfo, error) {
	reader, err := gozip.OpenReader(file.Name())
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	foundFiles := make([]os.FileInfo, 0)
	for _, f := range reader.File {
		foundFiles = append(foundFiles, f.FileInfo())
	}
	return foundFiles, nil
}
