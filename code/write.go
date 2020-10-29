package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	err := WriteFileAtomic("./ds/faabc.list", []byte("1"), 0755)
	fmt.Println(err)
}

// Write file to temp and atomically move when everything else succeeds.
func WriteFileAtomic(filename string, data []byte, perm os.FileMode) error {
	dir, name := path.Split(filename)
	fmt.Println(dir, "kkk", name)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm)
		}
	}
	f, err := ioutil.TempFile(dir, name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = f.Write(data)
	if err == nil {
		err = f.Sync()
	}
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	if permErr := os.Chmod(f.Name(), perm); err == nil {
		err = permErr
	}
	if err == nil {
		err = os.Rename(f.Name(), filename)
	}
	// Any err should result in full cleanup.
	if err != nil {
		os.Remove(f.Name())
	}
	return err
}
