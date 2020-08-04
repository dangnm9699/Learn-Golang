package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// UpdateIssue queries the GitHub issue tracker.
func UpdateIssue(owner, repo, number string, issue NewIssue) (*Issue, error) {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(&issue); err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", APIURL+fmt.Sprintf("/%s/%s/issues/%s", owner, repo, number), &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(USERNAME, PASSWORD)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("Patch failed: %s", res.Status)
	}
	var result Issue
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
