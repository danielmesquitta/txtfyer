package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ToTxtConfig struct {
	IgnorePaths []string
	Separator   string
}

func ToTtx(folderPath string, outputFilePath string, config ToTxtConfig) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if config.Separator == "" {
		config.Separator = "===== %s =====\n"
	}

	var processFile = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || shouldIgnore(path, config.IgnorePaths) {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		path = strings.TrimPrefix(path, folderPath)
		separator := fmt.Sprintf(config.Separator, path)
		if _, err := outputFile.WriteString(separator + string(content) + "\n"); err != nil {
			return err
		}
		return nil
	}

	err = filepath.Walk(folderPath, processFile)
	if err != nil {
		return err
	}

	return nil
}

func shouldIgnore(path string, ignore []string) bool {
	for _, ignorePath := range ignore {
		if strings.Contains(path, ignorePath) {
			return true
		}
	}
	return false
}
