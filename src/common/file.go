package common

import (
	"os"
	"path"
)

func CopyFile(src, dst string) error {
	srcData, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(dst), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, srcData, 0755)
}
