package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func EnsureDirectory(targetDirectory string) error {
	return os.MkdirAll(targetDirectory, 0755)
}

func VerifyChecksum(filePath string, checksum string, algorithm string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	var hash hash.Hash
	if algorithm == "md5" {
		hash = md5.New()
	} else if algorithm == "sha1" {
		hash = sha1.New()
	} else if algorithm == "sha256" {
		hash = sha256.New()
	} else {
		return fmt.Errorf("checksum algorithm not supported [%s]", algorithm)
	}

	bytes, err := io.Copy(hash, file)
	if err != nil {
		return err
	}

	fileChecksum := fmt.Sprintf("%x", hash.Sum(nil))
	if fileChecksum != checksum {
		return fmt.Errorf("checksum from [%s] doesn't match file checksum [%s]", checksum, fileChecksum)
	}

	log.Debug().Msgf("checksums match for file [%s] (size: %d bytes)", filePath, bytes)

	return nil
}

func ListFilesInDirTree(root string) (map[string]string, error) {
	fileMap := make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		return walk(fileMap, root, path, info, err)
	})
	if err != nil {
		return nil, err
	}

	return fileMap, nil
}

func walk(fileMap map[string]string, root string, path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		fileMap[path] = strings.TrimPrefix(strings.Replace(path, root, "", -1), fmt.Sprintf("%c", os.PathSeparator))
	}

	return nil
}

func CheckIfDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}
	return nil
}
