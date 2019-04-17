package mappath

import (
	"reflect"
	"testing"
)

func TestSeek(t *testing.T) {

	examples := []struct {
		input  map[string]interface{}
		path   string
		output interface{}
		err    error
	}{
		{
			input: map[string]interface{}{
				"hello":   "world",
				"goodbye": "moon",
			},
			path:   ".hello",
			output: "world",
			err:    nil,
		},
		{
			input: map[string]interface{}{
				"hello":   "world",
				"goodbye": "moon",
			},
			path:   ".moon",
			output: nil,
			err:    ErrNotFound,
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
			err: nil,
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
			err:    nil,
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
			err:    nil,
		},
		{
			input: map[string]interface{}{
				"hello": []interface{}{
					"world",
					"goodbye",
					"moon",
				},
			},
			path:   ".hello.7",
			output: nil,
			err:    ErrOutOfBounds,
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
			err:    nil,
		},
	}

	for _, example := range examples {
		example := example

		r, err := Seek(example.path, example.input)
		if err != example.err {
			t.Errorf("incorrect error; expected %v, got %v", example.err, err)
		}
		if !reflect.DeepEqual(r, example.output) {
			t.Errorf("incorrect result; expected %v, got %v", example.output, r)
		}

	}
}

func BenchmarkSeek(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			Seek(path, input)
		}
	}

}
