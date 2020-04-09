// +build !windows

package flatfs

import (
	"io/ioutil"
	"os"
)

func tempFile(dir, pattern string) (f *os.File, err error) {
	return ioutil.TempFile(dir, pattern)
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}