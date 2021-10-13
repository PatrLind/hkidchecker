# hkidchecker

Checker/validator for Hong Kong IDs

## Description

This Go package validates Hong Kong ID card IDs. Useful for example for
validating form input.

## Getting Started

### Installing

`go get -u github.com/PatrLind/hkidchecker`

### Example

```go
package main

import (
    "fmt"

    "github.com/PatrLind/hkidchecker"
)

func main() {
    hkid := "E364912(5)"
    valid := hkidchecker.CheckHKID(hkid)
    fmt.Printf("HKID '%s' valid: %t\n", hkid, valid)
}
```

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
