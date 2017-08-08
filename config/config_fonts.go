package config

import (
	"path/filepath"
	"os"
	"strings"
)

/**
字体配置
 */

func Child(path string) []string {
	fullPath, _ := filepath.Abs(path)
	listStr := make([]string, 0)
	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".ttf") {
			listStr = append(listStr, path)
		}
		return err
	})
	return listStr
}
