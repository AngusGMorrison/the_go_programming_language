// Package tar facilitates the extraction of FileInfo from tar files.
package tar

import (
	gotar "archive/tar"
	"io"
	"os"

	"github.com/angusgmorrison/the_go_programming_language/ch10/ex2/archive"
)

func init() {
	archive.RegisterFormat("tar", "ustar\x00", 0x101, Decode)
	archive.RegisterFormat("tar", "ustar ", 0x101, Decode)
}

// Decode extracts the FileInfo of all files found within the tar archive.
func Decode(file *os.File) ([]os.FileInfo, error) {
	reader := gotar.NewReader(file)
	foundFiles := make([]os.FileInfo, 0)
	for {
		header, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return foundFiles, err
		}
		foundFiles = append(foundFiles, header.FileInfo())
	}
	return foundFiles, nil
}
