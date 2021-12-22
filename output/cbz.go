package output

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func ConvertToCBZ(inputImages []string, outputCBZ string) {
	cbz, err := os.Create(outputCBZ)
	if err != nil {
		log.Println(err)
	}
	defer cbz.Close()

	zipStream := zip.NewWriter(cbz)
	defer zipStream.Close()
	for _, page := range inputImages {
		if AddPage(zipStream, page) != nil {
			log.Println(err)
		}
	}
}
func AddPage(zipStream *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(fileToZip *os.File) {
		err := fileToZip.Close()
		if err != nil {
			log.Println(err)
		}
	}(fileToZip)

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate

	writer, err := zipStream.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
