package common

import "github.com/google/uuid"

func UUIDv4() string {
	return uuid.New().String()
}
