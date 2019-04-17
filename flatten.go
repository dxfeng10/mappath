package mappath

import (
	"strconv"
	"strings"
)

// hello: world, goodbye: moon
// .hello: world, .goodbye: moon
func Flatten(subject interface{}) map[string]interface{} {
	out := map[string]interface{}{}

	var accum func(subject interface{}, keys ...string)

	accum = func(subject interface{}, keys ...string) {

		switch s := subject.(type) {
		case map[string]interface{}:
			for key, val := range s {
				out["."+strings.Join(append(keys, key), ".")] = val
				accum(val, append(keys, key)...)
			}
		case []interface{}:
			for i, val := range s {
				out["."+strings.Join(append(keys, strconv.Itoa(i)), ".")] = val
				accum(val, append(keys, strconv.Itoa(i))...)
			}
		default:
			out["."+strings.Join(keys, ".")] = s
		}
	}

	accum(subject)

	return out
}
