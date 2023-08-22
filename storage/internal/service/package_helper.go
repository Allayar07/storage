package service

import (
	"io/fs"
	"os"
	"path/filepath"
)

type CollectedFiles struct {
	Content *os.File
	Size    int64
}

func CollectForUploadFiles(pathRoot string) ([]CollectedFiles, error) {
	var collectFiles []CollectedFiles
	err := filepath.Walk(pathRoot, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			fileInfo, err := file.Stat()
			if err != nil {
				return err
			}

			collectFiles = append(collectFiles, CollectedFiles{
				Content: file,
				Size:    fileInfo.Size(),
			})
		}

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return collectFiles, nil
}
