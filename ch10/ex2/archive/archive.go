// Package archive provides a generic archive-reading function by maintaining a table of known
// archive formats.
package archive

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
)

type decoder func(file *os.File) ([]os.FileInfo, error)

// Format represents a registered archive file format
type Format struct {
	ext    string
	sig    []byte
	offset uint16
	decode decoder
}

var (
	formatsMu sync.Mutex // guards formats
	formats   = make([]*Format, 0, 4)
)

// RegisterFormat adds an archive format to the formats table, making it possible to detect and
// decode input files with that format.
func RegisterFormat(ext, sig string, offset uint16, decode decoder) {
	formatsMu.Lock()
	formats = append(formats, &Format{ext, []byte(sig), offset, decode})
	formatsMu.Unlock()
}

// Decode decodes an input archive file whose format has been registered and returns information
// about the files in contains.
func Decode(path string) ([]os.FileInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileFmt, err := detectFileFormat(file)
	if err != nil {
		return nil, err
	}
	return fileFmt.decode(file)
}

// detectFileFormat compares the input file's signature with the signatures of registered formats
func detectFileFormat(file *os.File) (*Format, error) {
	defer file.Seek(0, 0) // Rewind the file once done

	reader := bufio.NewReader(file)
	for _, fileFmt := range formats {
		sigWithOffset, err := reader.Peek(int(fileFmt.offset) + len(fileFmt.sig))
		if err != nil {
			return nil, fmt.Errorf("checking signature of %s: %v", file.Name(), err)
		}

		if bytes.Equal(sigWithOffset[fileFmt.offset:], fileFmt.sig) {
			return fileFmt, nil
		}
	}
	return nil, fmt.Errorf("no matching decoder for %s", file.Name())
}
