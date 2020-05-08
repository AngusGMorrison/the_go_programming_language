package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// CreateIssue POSTs a draft issue to the GitHub API and returns the created issue
func CreateIssue(owner, repo string, draft *Issue) (*Issue, error) {
	url := fmt.Sprintf(repoURL, owner, repo)
	resp, err := ajax(url, "POST", draft)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("CreateIssue failed: %s", resp.Status)
	}

	issue := Issue{}
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// ReadIssue queries the GitHub API for details of a known issue.
func ReadIssue(owner, repo, issueNum string) (*Issue, error) {
	url := fmt.Sprintf(issueURL, owner, repo, issueNum)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("read issue failed: %v", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &issue, nil
}

// UpdateIssue updates the specified issue via the GitHub API
func UpdateIssue(owner, repo, issueNum string, draft *Issue) (*Issue, error) {
	url := fmt.Sprintf(issueURL, owner, repo, issueNum)
	resp, err := ajax(url, "PATCH", draft)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("UpdateIssue failed: %s", resp.Status)
	}

	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

// CloseIssue closes the specified issue via the GitHub API
func CloseIssue(owner, repo, issueNum string) (*Issue, error) {
	url := fmt.Sprintf(issueURL, owner, repo, issueNum)
	draft := Issue{State: "closed"}

	resp, err := ajax(url, "PATCH", &draft)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("CloseIssue failed: %s", resp.Status)
	}

	issue := Issue{}
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &issue, nil
}

func ajax(url, method string, draft *Issue) (*http.Response, error) {
	payload, err := json.Marshal(draft)
	if err != nil {
		return nil, fmt.Errorf("JSON marshaling failed in ajax: %s", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("Request creation failed in ajax: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GOPL_GITHUB_ACCESS_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
