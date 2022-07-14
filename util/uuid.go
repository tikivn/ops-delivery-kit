package util

import (
	"crypto/md5"

	"github.com/google/uuid"
)

func CodeToUUID(code string) uuid.UUID {
	return uuid.NewHash(md5.New(), uuid.Nil, []byte(code), 4)
}

func MustParse(code string) (uuid.UUID, error) {
	val, err := uuid.Parse(code)
	if err != nil {
		return uuid.Nil, err
	}
	return val, nil
}
