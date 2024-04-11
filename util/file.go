package util

import (
	"os"
	"path/filepath"
)

func Ls(path string) ([]string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	paths := make([]string, len(dirs))
	for i := range dirs {
		paths[i] = filepath.Join(path, dirs[i].Name())
	}

	return paths, nil
}