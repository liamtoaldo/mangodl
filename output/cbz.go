package output

import (
	"encoding/xml"
	"archive/zip"
	"io"
	"log"
	"os"
)
type ComicInfo struct  {	
	XMLName xml.Name `xml:"ComicInfo"`
	Title string `xml:"Title"`
	Summary string `xml:"Summary"`
	Series string `xml:"Series"`
}

func ConvertToCBZ(inputImages []string, outputCBZ string, metaData ComicInfo) {
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
	if AddComicInfo(zipStream, metaData) != nil {

	}
}

func AddComicInfo(zipStream *zip.Writer, metaData ComicInfo) error {
	file, err := os.Create("ComicInfo.xml")
	if err != nil {
		return err
	}
	
	// write Metadata to xml file 
	file.WriteString(xml.Header)
	encoder = xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err = encoder.Encode(&metaData)
	if err != nil {
		return err
	}


	// zip stuff ( mimic Adding pages to zip archive)

	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	writer, err := zipStream.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)

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
