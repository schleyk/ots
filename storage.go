package main

import (
	"errors"
	"fmt"
	"time"
)

var errSecretNotFound = errors.New("Secret not found")

type storage interface {
	Create(secret string, expireIn time.Duration) (string, error)
	ReadAndDestroy(id string) (string, error)
}

func getStorageByType(t string) (storage, error) {
	switch t {
	case "mem":
		return newStorageMem(), nil
	case "redis":
		return newStorageRedis()
	default:
		return nil, fmt.Errorf("Storage type %q not found", t)
	}
}
