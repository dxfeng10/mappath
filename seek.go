package mappath

import (
	"errors"
	"strconv"
	"strings"
)

// Seek within the subject, generally a map[string]interface{}, for the value
// at the given path. This allows access to deeply nested values without
// creating interim types or continual type-checking -- it is especially
// helpful when dealing with nested JSON responses from APIs.
//
// This is most easily understood with an example. Suppose we have the map:
//
//   hello: world
//   goodbye: moon
//   banana:
//     phone: ring
//     ding: ding
//
// Values can then by found with Seek at various paths:
//
//   .banana.phone: ring
//   .banana:
//      phone: ring
//      ding: ding
//
// Array/Slice keys are also accessible, with an integer component in the
// keypath.
//
// Only the provided path is traversed, with errors being returned if the
// keypath isn't found, if an index it out of bounds, etc.
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
