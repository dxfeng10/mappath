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

## But Why?
Recently, I've needed to extract only one or two fields from a deeply nested
JSON structure. Each part I needed was at different keypaths. Instead of
creating numerous throw-away structs, or coercing/type-checking each level, I
figured I'd pull out this functionality.

Using something like `mappath.Seek` with a single type-check, reduced the size
and complexity of my codebase. It also made it more readable, which is a Good
Thing.
