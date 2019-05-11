package gofpdi

import (
//	"io"
//	"net/http"
//	"fmt"

	realgofpdi "github.com/phpdave11/gofpdi"
	//"github.com/phpdave11/gofpdf"
//	"github.com/davecgh/go-spew/spew"
)

var fpdi = realgofpdi.NewImporter()

// gofpdiPdf is a partial interface that only implements the functions we need
// from the PDF generator to put the HTTP images on the PDF.
type gofpdiPdf interface {
	GetNextObjectID() int
	ImportObjects(objs map[string]string)
	ImportTemplates(tpls map[string]string)
	UseImportedTemplate(tplName string, x float64, y float64, w float64, h float64)
/*
	GetImageInfo(imageStr string) *gofpdf.ImageInfoType
	ImageTypeFromMime(mimeStr string) string
	RegisterImageReader(imgName, tp string, r io.Reader) *gofpdf.ImageInfoType
*/
	SetError(err error)
}

// Register registers a HTTP image. Downloading the image from the provided URL
// and adding it to the PDF but not adding it to the page. Use Image() with the
// same URL to add the image to the page.
func ImportPage(f gofpdiPdf, sourceFile string, pageno int, box string) int {

	// Set source file for fpdi
	fpdi.SetSourceFile(sourceFile)

	// gofpdi needs to know where to start the object id at.
	// By default, it starts at 1, but gofpdf adds a few objects initially.
	startObjId := 3//f.GetNextObjectID()

	// Set gofpdi next object ID to  whatever the value of startObjId is
	fpdi.SetNextObjectID(startObjId)

	// Import page
	tpl := fpdi.ImportPage(pageno, box)

	// Import objects into current pdf document
	tplObjIds := fpdi.PutFormXobjects()

	// Set template names and ids (hashes) in gopdf
	f.ImportTemplates(tplObjIds)

	// Get a map[int]string of the imported objects.
	// The map keys will be the ID of each object.
	imported := fpdi.GetImportedObjects()

	// Import gofpdi objects into gopdf, starting at whatever the value of startObjId is
	f.ImportObjects(imported)

	return tpl
}

func UseImportedTemplate(f gofpdiPdf, tplid int, x float64, y float64, w float64, h float64) {
	// Get values from fpdi
	tplName, scaleX, scaleY, tX, tY := fpdi.UseTemplate(tplid, x, y, w, h)

	f.UseImportedTemplate(tplName, scaleX, scaleY, tX, tY)
}
