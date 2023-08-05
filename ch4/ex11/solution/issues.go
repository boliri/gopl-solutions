package issues

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BaseURL 		= "https://api.github.com"
	IssuesURL 		= BaseURL + "/search/issues"
	RepoIssuesURL 	= BaseURL + "/repos/%s/%s/issues"
	RepoIssueURL 	= BaseURL + "/repos/%s/%s/issues/%d"
)

// Types grabbed from gopl.io/ch4/github
type IssuesSearchResult struct {
	TotalCount	int 		`json:"total_count"`
	Items		[]*Issue
}
type Issue struct {
	Number		int
	HTMLURL		string		`json:"html_url"`
	Title		string
	State		string
	User		*User
	CreatedAt	time.Time	`json:"created_at"`
	Body		string		// in Markdown format
}
type User struct {
	Login		string
	HTMLURL		string		`json:"html_url"`
}
/* ------------------------------------ */

type Auth struct {
	Username	string
	Token		string
}

type SearchAllRequest struct {
	Owner		string
	Repo		string
	State		string
	Auth
}
type SearchAllResponse struct {
	Items		[]*Issue
}

type SearchRequest struct {
	SearchAllRequest
	IssueNumber			uint
}
type SearchResponse struct {
	Issue		*Issue
}

type ModifyIssueParameters struct {
	Title		string		`json:"title"`
	Body		string		`json:"body"`
}

type CreateRequest struct {
	Owner					string		
	Repo					string
	CreateRequestParameters
	Auth
}
type CreateRequestParameters struct {
	ModifyIssueParameters
}
type CreateResponse struct {
	Issue		*Issue
}

type EditRequest struct {
	Owner					string		
	Repo					string
	IssueNumber				uint
	EditRequestParameters
	Auth
}
type EditRequestParameters struct {
	ModifyIssueParameters
}
type EditResponse struct {
	Issue		*Issue
}

type CloseRequest struct {
	Owner					string
	Repo					string
	IssueNumber				uint
	CloseRequestParameters
	Auth
}
type CloseRequestParameters struct {
	State			string		`json:"state"`
}
type CloseResponse struct {
	EditResponse
}

func (a Auth) getAuthHeader() (string, string) {
	creds := fmt.Sprintf("%s:%s", a.Username, a.Token)
	return "Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(creds)))
}

func SearchAll(r *SearchAllRequest) (*SearchAllResponse, error) {
	url := fmt.Sprintf(RepoIssuesURL, r.Owner, r.Repo)
	query := fmt.Sprintf("?state=%s", r.State)
	url += query

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(r.getAuthHeader())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
    	body := string(bodyBytes)
		return nil, fmt.Errorf("search query failed:\n\nHTTP %s\n%s", resp.Status, body)
	}

	var result SearchAllResponse
	if err := json.NewDecoder(resp.Body).Decode(&result.Items); err != nil {
		return nil, err
	}

	return &result, nil
}

func Search(r *SearchRequest) (*SearchResponse, error) {
	url := fmt.Sprintf(RepoIssueURL, r.Owner, r.Repo, r.IssueNumber)
	query := fmt.Sprintf("?state=%s", r.State)
	url += query
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(r.getAuthHeader())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
    	body := string(bodyBytes)
		return nil, fmt.Errorf("search query failed:\n\nHTTP %s\n%s", resp.Status, body)
	}

	result := SearchResponse{Issue: &Issue{}}
	if err := json.NewDecoder(resp.Body).Decode(result.Issue); err != nil {
		return nil, err
	}

	return &result, nil
}

func Create(r *CreateRequest) (*CreateResponse, error) {
	url := fmt.Sprintf(RepoIssuesURL, r.Owner, r.Repo)

	reqBody, err := json.Marshal(r.CreateRequestParameters)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set(r.getAuthHeader())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
    	body := string(bodyBytes)
		return nil, fmt.Errorf("creating new issue failed:\n\nHTTP %s\n%s", resp.Status, body)
	}

	var result CreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result.Issue); err != nil {
		return nil, err
	}

	return &result, nil
}

func Edit(r *EditRequest) (*EditResponse, error) {
	url := fmt.Sprintf(RepoIssueURL, r.Owner, r.Repo, r.IssueNumber)

	reqBody, err := json.Marshal(r.EditRequestParameters)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set(r.getAuthHeader())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
    	body := string(bodyBytes)
		return nil, fmt.Errorf("editing existing issue failed:\n\nHTTP %s\n%s", resp.Status, body)
	}

	var result EditResponse
	if err := json.NewDecoder(resp.Body).Decode(&result.Issue); err != nil {
		return nil, err
	}

	return &result, nil
}

func Close(r *CloseRequest) (*CloseResponse, error) {
	url := fmt.Sprintf(RepoIssueURL, r.Owner, r.Repo, r.IssueNumber)

	reqBody, err := json.Marshal(r.CloseRequestParameters)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(reqBody))

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set(r.getAuthHeader())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
    	body := string(bodyBytes)
		return nil, fmt.Errorf("closing existing issue failed:\n\nHTTP %s\n%s", resp.Status, body)
	}

	var result CloseResponse
	if err := json.NewDecoder(resp.Body).Decode(&result.Issue); err != nil {
		return nil, err
	}

	return &result, nil
}
