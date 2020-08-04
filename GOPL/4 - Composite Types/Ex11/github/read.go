package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetIssue queries the GitHub issue tracker.
func GetIssue(owner, repo, number string) (*Issue, error) {
	par := fmt.Sprintf("/%s/%s/issues/%s", owner, repo, number)
	res, err := http.Get(APIURL + par)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}
	var result Issue
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}
	return &result, nil
}
