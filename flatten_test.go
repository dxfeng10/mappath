package mappath

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {

	examples := []struct {
		input  map[string]interface{}
		path   string
		output interface{}
	}{
		{
			input: map[string]interface{}{
				"hello":   "world",
				"goodbye": "moon",
			},
			path:   ".hello",
			output: "world",
		},
		{
			input: map[string]interface{}{
				"hello": map[string]interface{}{
					"world": map[string]interface{}{
						"goodbye": "moon",
					},
				},
			},
			path: ".hello.world",
			output: map[string]interface{}{
				"goodbye": "moon",
			},
		},
		{
			input: map[string]interface{}{
				"hello": map[string]interface{}{
					"world": map[string]interface{}{
						"goodbye": "moon",
					},
				},
			},
			path:   ".hello.world.goodbye",
			output: "moon",
		},
		{
			input: map[string]interface{}{
				"hello": []interface{}{
					"world",
					"goodbye",
					"moon",
				},
			},
			path:   ".hello.0",
			output: "world",
		},
		{
			input: map[string]interface{}{
				"hello": []interface{}{
					map[string]interface{}{
						"hello": "world",
					},
					map[string]interface{}{
						"goodbye": "moon",
					},
				},
			},
			path:   ".hello.1.goodbye",
			output: "moon",
		},
	}

	for i, example := range examples {
		example := example

		r := Flatten(example.input)
		if !reflect.DeepEqual(r[example.path], example.output) {
			t.Errorf("incorrect result for %d; expected %v, got %v", i, example.output, r[example.path])
		}

	}
}

func BenchmarkFlatten(b *testing.B) {
	input := map[string]interface{}{
		"hello": []interface{}{
			map[string]interface{}{
				"hello": "world",
			},
			map[string]interface{}{
				"goodbye": "moon",
			},
		},
	}

	paths := []string{
		".hello",
		".hello.0",
		".hello.0.hello",
		".hello.1.goodbye",
		".goodbye",
		".hello.7",
		".hello.7.goodbye",
	}

	// var v interface{}
	r := Flatten(input)
	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			_ = r[path]
		}
	}

}
