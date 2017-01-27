# go-expanding

[![Build Status](https://travis-ci.org/schwarmco/go-expanding.svg?branch=master)](https://travis-ci.org/schwarmco/go-expanding)
[![GoDoc](https://godoc.org/github.com/schwarmco/go-expanding?status.svg)](https://godoc.org/github.com/schwarmco/go-expanding)

a package for working with ranges in golang

## Installation

In order to start, `go get` this repository:

```
go get github.com/schwarmco/go-expanding
```

## Usage

```go
import (
    "fmt"
    "github.com/schwarmco/go-expanding"
)

func main() {

	a := expanding.Integer{
		Min:  0,
		Max:  10,
		Step: 2,
	}

	fmt.Println(a.Expand())

	a.Step = 3

	fmt.Println(a.Expand())

	// Output:
	// [0 2 4 6 8 10]
	// [0 3 6 9]
}
```
