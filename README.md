# mappath [![GoDoc](https://godoc.org/github.com/ryanfaerman/mappath?status.svg)](https://godoc.org/github.com/ryanfaerman/mappath)

mappath is a go library for traversing generic maps for values by keypath.

This package is most helpful when you need to pull out a view values from a
deeply nested from some sort of map (typically from JSON). Instead of making
numerous interim types, you can extract just the values at a known path.

This pairs well with something like
[mapstructure](https://github.com/mitchellh/mapstructure).

## Installation

Standard `go get`:

```
$ go get github.com/ryanfaerman/mappath
```

## Trivial Example

Consider a deeply nested JSON (much deeper than this example):

```json
{
  "hello": "world",
  "goodbye": "moon",
  "banana": [
    {
      "phone": "ring",
      "ding": "ding"
    }
  ]
}
```

where we need to extract the value at the path `banana.0.phone`. With a
traditional approach, we'd need to create at least one interim struct or dig
through each level in our path with a combination of type-checking and
coercion.

```go
if banana, ok := data["banana"]; ok {
  if bananaVal, ok := banana.([]map[string]interface{}); ok {
    if len(banana) >= 1 {
      if item, ok := banana[0].(map[string]interface{}); ok {
        // ... and so on ...
      }
    }
  }
}
```

If there are multiple keypaths needing extraction, this gets old pretty quick.
Instead, we can read the JSON into a `map[string]interface{}` and using `Seek`,
this is almost trivial.

```go
phone, err := Seek(".banana.0.phone", data)
if err != nil {
  // handle appropriately
}
```

Granted, this will still need a type coercion (with a check) but we're only
doing this once with far less error handling.

This also opens the door to reading that keypath from some other input, like a
flag.

## But Why?
Recently, I've needed to extract only one or two fields from a deeply nested
JSON structure. Each part I needed was at different keypaths. Instead of
creating numerous throw-away structs, or coercing/type-checking each level, I
figured I'd pull out this functionality.

Using something like `mappath.Seek` with a single type-check, reduced the size
and complexity of my codebase. It also made it more readable, which is a Good
Thing.
