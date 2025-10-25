package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

func HashFile(path string) (string, error) {
	data, err := ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

func RemoveDirectory(path string) error {
	return os.RemoveAll(path)
}