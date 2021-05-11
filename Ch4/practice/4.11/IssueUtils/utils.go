package IssueUtils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Parameter struct {
	Owner  string
	Repo   string
	Number string
	Token  string
	Issue
}
type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var rootURL = "https://api.github.com/repos/"

func (params Parameter) GetIssueList() ([]Issue, error) {
	url := rootURL + params.Owner + "/" + params.Repo + "/issues"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (params Parameter) GetIssue() (Issue, error) {
	url := rootURL + params.Owner + "/" + params.Repo + "/issues" + "/" + params.Number
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return Issue{}, err
	}
	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return Issue{}, err
	}
	return issue, nil
}

func (params Parameter) CreateIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(params.Issue); err != nil {
		log.Println(err)
		return false
	}
	u := rootURL + params.Owner + "/" + params.Repo + "/issues" +
		"?access_token=" + params.Token
	_, err := http.Post(u, "application/json", &buf)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (params Parameter) EditIssue() bool {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(params.Issue); err != nil {
		return false
	}
	u := rootURL + params.Owner + "/" + params.Repo + "/issues" +
		"/" + params.Number + "?access_token=" + params.Token
	request, err := http.NewRequest(http.MethodPatch, u, &buf)
	if err != nil {
		log.Println(err)
		return false
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
