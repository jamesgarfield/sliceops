# sliceops
Slice Operations generator for goast

Provides iteration and sort functionality for abtirary slice types using [goast](http://github.com/jamesgarfield/goast).

## Usage

1. Install goast: `go get github.com/jamesgarfield/goast`
1. Get sliceops: `go get github.com/jamesgarfield/sliceops`
1. Annotate files containing slice types: `//go:generate goast write impl github.com/jamesgarfield/sliceops`
