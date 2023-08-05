package poster

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const omdbURL = "https://www.omdbapi.com/?apikey=%s&t=%s&type=movie"

type Movie struct {
	Title		string
	PosterURL	string		`json:"Poster"`
}

type PosterRequest struct {
	ApiKey		string
	Title		string
}

type PosterResponse struct {
	Movie
}

func Download(r *PosterRequest) error {
	url := fmt.Sprintf(omdbURL, r.ApiKey, r.Title)

	// Get OMDB data first
	respOmdb, err := http.Get(url)
	if err != nil {
		return err
	}
	defer respOmdb.Body.Close()

	if respOmdb.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(respOmdb.Body)
    	body := string(bodyBytes)
		return fmt.Errorf("get movie failed:\n\nHTTP %s\n%s", respOmdb.Status, body)
	}

	var result PosterResponse
	if err := json.NewDecoder(respOmdb.Body).Decode(&result); err != nil {
		return err
	}

	// Download the image using the URL returned by OMDB
	respImage, err := http.Get(result.PosterURL)
	if err != nil {
		return fmt.Errorf("download image failed:\n\nHTTP %s\n", respImage.Status)
	}
	defer respImage.Body.Close()

	// Finally, save image bytes to file
	f, err := os.Create(fmt.Sprintf("%s%s", strings.ToLower(result.Title), filepath.Ext(result.PosterURL)))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, respImage.Body)
	if err != nil {
		return err
	}

	return nil
}