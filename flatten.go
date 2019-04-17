package mappath

import (
	"strconv"
	"strings"
)

// Flatten a subject, usually a map[string]interface{}, into a map containing
// keys for every keypath. This is probably terribly inperformant for large,
// highly nested maps.
//
// This is most easily understood with an example. Suppose we have the map:
//
//   hello: world
//   goodbye: moon
//   banana:
//     phone: ring
//     ding: ding
//
// Flatten will flatten this into the following:
//
//   .hello: world
//   .goodbye: moon
//   .banana:
//      phone: ring
//      ding: ding
//   .banana.phone: ring
//   .banana.ding: ding
//
// Thus, creating a map where any value can be found via its keypath.
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
