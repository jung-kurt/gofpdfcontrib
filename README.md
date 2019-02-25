# gofpdfcontrib

Packages that extend gofpdf and have non-standard dependencies.

Install ($GOPATH): `go get -u github.com/jung-kurt/gofpdfcontrib/...`

Install (module): `git clone https://github.com/jung-kurt/gofpdfcontrib.git`

Test: `go test -v ./...`

## Quick start

```go
package main

import (
	"fmt"
	"os"

	"github.com/boombuler/barcode/code128"
	"github.com/jung-kurt/gofpdf/v2"
	"github.com/jung-kurt/gofpdfcontrib/barcode"
)

func main() {

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()

	bcode, err := code128.Encode("gofpdf")

	if err == nil {
		key := barcode.Register(bcode)
		var width float64 = 100
		var height float64 = 10.0
		barcode.BarcodeUnscalable(pdf, key, 15, 15, &width, &height, false)
		err = pdf.OutputFileAndClose("barcode.pdf")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
```
