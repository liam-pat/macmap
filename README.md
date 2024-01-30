# Macmap

## Description

Find the vendor information by mac address

Office lookup website : https://regauth.standards.ieee.org/standards-ra-web/pub/view.html

## Table Bits
| MAL     | MAM     | MAS     |
|---------|---------|---------|
| 36 bits | 28 bits | 24 bits |

## Installation

```bash
bash > go get -u github.com/YaoMiss/macmap
```

## Usage

```go
import (
    macmap "github.com/YaoMiss/macmap@v1.0.6"
)

m1 := macmap.Search("18:65:90:dc:c0:cb")
m2 := manuf.Search("00:ec:0a:ff:b7:27")

fmt.Println(m1)
fmt.Println(m2)
```