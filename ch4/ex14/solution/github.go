package github

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	IssuesURL 		= "https://api.github.com/search/issues"
	BugsURL 		= IssuesURL + "?q=label:bug"
	MilestonesURL 	= IssuesURL + "?q=label:milestone"
)

//go:embed tpl/landing.html
var rawTpl string

var landingTpl *template.Template = template.Must(template.New("landing").Parse(rawTpl))

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	Creator   *User `json:"user"`
	Assignee  *User `json:"assignee"`
	Labels	  []*Label
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Label struct {
	Name	string
}

func (iss Issue) GetLabelsStr() string {
	var lnames []string
	for _, l := range iss.Labels {
		lnames = append(lnames, l.Name)
	}

	return strings.Join(lnames, ", ")
}

func get(url string, out *IssuesSearchResult) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
    	body := string(bodyBytes)
		return fmt.Errorf("query failed:\n\nHTTP %s\n%s", resp.Status, body)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return err
	}

	return nil
}

func BuildLanding(out io.Writer) {
	milestones := IssuesSearchResult{}
	err := get(MilestonesURL, &milestones)
	if err != nil {
		log.Fatal(err)
	}

	bugs := IssuesSearchResult{}
	err = get(BugsURL, &bugs)
	if err != nil {
		log.Fatal(err)
	}

	var data struct {
		Bugs, Milestones	[]*Issue
	}
	data.Bugs = bugs.Items
	data.Milestones = milestones.Items
	if err = landingTpl.Execute(out, data); err != nil {
		log.Fatal(err)
	}
}
