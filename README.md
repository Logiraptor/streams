streams
=======

"Generic" implementation of streams.

This package uses [gen](https://github.com/clipperhouse/gen) to create a
generic pipeline for any go type. See demo for example usage.

Basic Usage
======

Annotate a type with `// +stream` like so:

```Go

// +stream
type T struct {
	A int
	B string
}

```

Then run the command `streamgen` in the package directory to generate a stream type for creating pipelines.