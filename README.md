# Intspec

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/intspec.svg)](https://pkg.go.dev/github.com/akramarenkov/intspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/intspec)](https://goreportcard.com/report/github.com/akramarenkov/intspec)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/intspec/badge.svg)](https://coveralls.io/github/akramarenkov/intspec)

## Purpose

Library that provides to detect the parameters of integer types: minimum, maximum values ​​and bit size

## Usage

Example:

```go
package main


import (
    "fmt"

    "github.com/akramarenkov/intspec"
)

func main() {
    minimum, maximum := intspec.Range[int8]()
    fmt.Println(minimum)
    fmt.Println(maximum)
    // Output:
    // -128
    // 127
}
```
