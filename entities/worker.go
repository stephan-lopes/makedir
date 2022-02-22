package entities

import (
	"io/fs"
	"os"
)

type Worker struct {
	Path []*Directory
}

func CreateList(config *YAMLStruct) []*Directory {
	var list []*Directory
	for _, v := range config.PackageName.Directories {
		list = append(list, &Directory{v.Path, v.Perm})
	}
	return list
}

func (w *Worker) CreateDir() error {

	for _, dir := range w.Path {
		if err := os.MkdirAll(dir.Path, fs.FileMode(dir.Perm)); err != nil {
			return err
		}
	}

	return nil

}
