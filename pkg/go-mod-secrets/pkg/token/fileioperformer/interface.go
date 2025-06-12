//
//
//
//
//
// SPDX-License-Identifier: Apache-2.0'
//

package fileioperformer

import (
	"io"
	"os"
)

type FileIoPerformer interface {
	// OpenFileReader is a function that opens a file and returns an io.Reader (at a minimum)
	OpenFileReader(name string, flag int, perm os.FileMode) (io.Reader, error)
	// OpenFileWriter is a function that opens a file and returns an io.WriteCloser (at a minimum)
	OpenFileWriter(name string, flag int, perm os.FileMode) (io.WriteCloser, error)
	// MkdirAll creates a directory tree (see os.MkdirAll)
	MkdirAll(path string, perm os.FileMode) error
}
