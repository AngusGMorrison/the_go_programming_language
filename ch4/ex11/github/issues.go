package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ReadIssue queries the Github API for details of a known issue.
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

// CloseIssue closes the specified issue via the GitHub API
func CloseIssue(owner, repo, issueNum string) (*Issue, error) {
	url := fmt.Sprintf(issueURL, owner, repo, issueNum)
	issue := Issue{State: "closed"}

	resp, err := patch(url, &issue)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("close issue failed: %v", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &issue, nil
}

func patch(url string, issue *Issue) (*http.Response, error) {
	payload, err := json.MarshalIndent(*issue, "", "    ")
	fmt.Printf("%s", payload)
	if err != nil {
		return nil, fmt.Errorf("JSON marshaling failed in CloseIssue: %s", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("Request creation failed in CloseIssue: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GOPL_GITHUB_ACCESS_TOKEN"))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
