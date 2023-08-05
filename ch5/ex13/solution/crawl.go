package solution

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"net/http"
	urlpkg "net/url"
)

const assetsBasepath = "../solution"

var storableDomain string

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
		return list
	}

	if storableDomain == "" {
		urlp, _ := urlpkg.Parse(url)
		rootdir := fmt.Sprintf("%s/%s", assetsBasepath, urlp.Hostname())
		err := makeDir(rootdir)
		if err != nil {
			log.Print(err)
			return list
		}

		storableDomain = urlp.Hostname()
	}

	filtered := getStorableUrls(list)
	saveUrls(filtered)

	return list
}

func getStorableUrls(urls []string) []string {
	var filtered []string
	for _, url := range urls {
		if !strings.Contains(url, storableDomain) {
			continue
		}

		filtered = append(filtered, url)
	}
	return filtered
}

func makeDir(dirname string) error {
	err := os.Mkdir(dirname, 0750)
	if err != nil {
		return fmt.Errorf("creating dir %s: %s", dirname, err)
	}

	return nil
}

func saveUrls(urls []string) {
	for _, url := range urls {
		saveSingleUrl(url)
	}
}

func saveSingleUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("could not fetch %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("got HTTP %d from %s\n", resp.StatusCode, url)
		return
	}

	bytes, err := io.ReadAll(resp.Body)

	urlp, err := urlpkg.Parse(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	dirname := fmt.Sprintf("%s/%s/%s", assetsBasepath, storableDomain, urlp.Path)
	makeDir(dirname) // omitting errors, since the dir may exist already and can pollut the output

	filepath := fmt.Sprintf("%s/contents.html", dirname)
	os.WriteFile(filepath, bytes, 0750)
}