package storage

import (
	"os"
	"path/filepath"
)

type SecureStorage struct {
	basePath       string
	masterPassword []byte
}

/* The encrypted files are identified using magic bytes + version. */
var magicBytes = []byte("KELCA")
var version byte = 1

func NewSecureStorage(password string) (*SecureStorage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	basePath := filepath.Join(homeDir, ".kelca")

	dirs := []string{
		filepath.Join(basePath, "ca", "root", "private"),
		filepath.Join(basePath, "ca", "root", "certs"),
		filepath.Join(basePath, "ca", "intermediate", "private"),
		filepath.Join(basePath, "ca", "intermediate", "certs"),
		filepath.Join(basePath, "ca", "crl"),
		filepath.Join(basePath, "ca", "db"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return nil, err
		}
	}

	return &SecureStorage{
		basePath:       basePath,
		masterPassword: []byte(password),
	}, nil
}
