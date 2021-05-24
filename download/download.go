package download

import (
	"errors"
	"io"
	"net/http"
	"os"
)

//This function is used to download a file in the given url
func DownloadFile(URL, fileName string) error {
	if URL == "" {
		return errors.New("couldn't download this image, url not existing")
	}
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
