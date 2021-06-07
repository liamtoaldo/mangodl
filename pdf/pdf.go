package pdf

import (
	"bufio"
	"github.com/signintech/gopdf"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//ConvertToPDF takes as input the path of the files and creates a PDF with the images of the pages
func ConvertToPDF(inputImages []string, outputPDF string) {
	pdf := gopdf.GoPdf{}
	pagesWidth, pagesHeight := getSizes(inputImages[len(inputImages)/2])                                       //GetSizes is not precise by default, so dividing the sizes by 1.8 the page should return to the normal size
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: float64(pagesWidth) / 1.8, H: float64(pagesHeight) / 1.8}}) //595.28, 841.89 = A4
	for _, page := range inputImages {
		pdf.AddPage()
		err := pdf.Image(page, 0, 0, nil)
		if err != nil {
			log.Print(err)
		}
	}
	err := pdf.WritePdf(outputPDF)
	if err != nil {
		log.Print(err)
	}
}

func getSizes(file string) (width int, height int) {
	//open the file and create a reader
	inputFile, _ := os.Open(file)
	reader := bufio.NewReader(inputFile)

	// Decode image.
	img, _, _ := image.DecodeConfig(reader)

	return img.Width, img.Height
}

func GetNumberOfPages(path string) int {
	files, _ := ioutil.ReadDir(path)
	length := 0
	for _, file := range files {
		if strings.Contains(file.Name(), "jpg") {
			length++
		}
	}
	return length
}
