package xkcd

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	comicURL 	   = "https://xkcd.com/%s/info.0.json"
	comicsDirName  = "comics"
)

//go:embed comics
var comicsFs embed.FS

type Comic struct {
	Number		int		`json:"num"`
	Transcript	string	`json:"transcript"`
}

func (c Comic) GetURL() string {
	return fmt.Sprintf(comicURL, fmt.Sprintf("%d", c.Number))
}


func GetComicsByTerm(term string) ([]Comic, error) {
	var result []Comic

	direntries, err := comicsFs.ReadDir(comicsDirName)
	if err != nil {
		return nil, fmt.Errorf("could not get read contents of the comics dir: %s", err)
	}

	for _, entry := range direntries {
		comicPath := fmt.Sprintf("%s/%s", comicsDirName, entry.Name())

		bytes, err := comicsFs.ReadFile(comicPath)
		if err != nil {
			return nil, fmt.Errorf("could not read contents from comic file %s: %s", comicPath, err)
		}

		var c Comic
		if err := json.Unmarshal(bytes, &c); err != nil {
			return nil, fmt.Errorf("something happened while trying to unmarshall contents from comic file %s: %s", comicPath, err)
		}

		if strings.Contains(c.Transcript, term) {
			result = append(result, c)
		}
	}

	return result, nil
}