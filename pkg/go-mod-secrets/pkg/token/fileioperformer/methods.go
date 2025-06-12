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

type defaultFileIoPerformer struct{}

func NewDefaultFileIoPerformer() FileIoPerformer {
	return &defaultFileIoPerformer{}
}

func (*defaultFileIoPerformer) OpenFileReader(name string, flag int, perm os.FileMode) (io.Reader, error) {
	return os.OpenFile(name, flag, perm)
}

func (*defaultFileIoPerformer) OpenFileWriter(name string, flag int, perm os.FileMode) (io.WriteCloser, error) {
	return os.OpenFile(name, flag, perm)
}

func (*defaultFileIoPerformer) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// MakeReadCloser will turn an an io.Reader into an io.ReadCloser
// if the underlying object does not already support io.ReadCloser
func MakeReadCloser(reader io.Reader) io.ReadCloser {
	rc, ok := reader.(io.ReadCloser)
	if !ok && reader != nil {
		rc = io.NopCloser(reader)
	}
	return rc
}
