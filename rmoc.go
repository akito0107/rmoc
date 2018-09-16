package rmoc

import (
	"io"
	"os"
	"path/filepath"
)


func OverrideFile(src io.Reader, path, filename string) error {
	name := filepath.Join(path, filename)
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		if err := os.Remove(name); err != nil {
			return err
		}
	}
	return writeFile(src, name)
}

//go:generate generr -t fileAlreadyExists -i
type fileAlreadyExists interface {
	FileAlreadyExists() (filename string)
}

func CreateFileWithAbort(src io.Reader, path, filename string) error {
	name := filepath.Join(path, filename)
	if stat, err := os.Stat(name); stat != nil {
		return &FileAlreadyExists{Filename: name}
	} else if !os.IsNotExist(err) {
		return err
	}
	return writeFile(src, name)
}

func writeFile(src io.Reader, filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, src); err != nil {
		return err
	}

	return nil
}
