## go.qrt
[![GoDoc](https://godoc.org/github.com/GeertJohan/go.qrt?status.png)](https://godoc.org/github.com/GeertJohan/go.qrt)
Generate QR code's for in your terminal.

go.qrt generates a QR code based on UTF-8 characters. This allow you to display a QR code in your terminal

### Install
`go get github.com/GeertJohan/go.qrt`

### Usage example
```go
package main

import (
	"fmt"
	"github.com/GeertJohan/go.qrt"
)

func main() {
	str, err := qrt.Generate("some example text")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(str)
}
```

### Documentation
For more information about the options you can use for QR generation, please check the documentation at [godoc.org/github.com/GeertJohan/go.qrt][godoc]

### License
This project is licensed under a Simplified BSD license. Please read the [LICENSE file][license].

### Attribution
This package imports and uses [code.google.com/p/rsc/qr][rsc-qr], a QR generation package by Russ Cox.

### Screenshot
![QR code in terminal](/screenshot.png "QR code in terminal")

Depending on your terminal and font you may need to scan the QR code under an angle to make it appear square for the scanner.

 [godoc]: https://godoc.org/github.com/GeertJohan/go.qrt
 [license]: https://github.com/GeertJohan/go.qrt/blob/master/LICENSE
 [rsc-qr]: https://code.google.com/p/rsc/source/browse/#hg%2Fqr`
