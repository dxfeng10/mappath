package mappath

import (
	"errors"
	"strconv"
	"strings"
)

func Seek(path string, subject interface{}) (interface{}, error) {
	keys := strings.SplitN(strings.TrimPrefix(path, "."), ".", 2)

	var key, nestedKey string
	key = keys[0]
	if len(keys) > 1 {
		nestedKey = keys[1]
	}

	switch s := subject.(type) {
	case map[string]interface{}:
		if val, ok := s[key]; ok {
			if nestedKey != "" {
				return Seek(nestedKey, val)
			}
			return val, nil
		} else {
			return nil, ErrNotFound
		}
	case []interface{}:
		i, err := strconv.Atoi(key)
		if err != nil {
			return nil, ErrInvalidIndex
		}

		if i < len(s) {
			if nestedKey != "" {
				return Seek(nestedKey, s[i])
			}

			return s[i], nil
		}

		return nil, ErrOutOfBounds
	case interface{}:
		return s, nil
	default:
		return nil, ErrUnseekable
	}

	return nil, ErrInvalidKey
}

var (
	ErrInvalidKey   = errors.New("invalid key")
	ErrUnseekable   = errors.New("unseekable type")
	ErrOutOfBounds  = errors.New("index out of bounds")
	ErrInvalidIndex = errors.New("cannot index slice with non-integer")
	ErrNotFound     = errors.New("value not found")
)
