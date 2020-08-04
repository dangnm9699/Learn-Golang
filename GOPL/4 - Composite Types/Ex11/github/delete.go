package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DeleteIssue queries the GitHub issue tracker.
func DeleteIssue(owner, repo, number string, issue NewIssue) error {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(&issue); err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", APIURL+fmt.Sprintf("/%s/%s/issues/%s", owner, repo, number), &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(USERNAME, PASSWORD)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return fmt.Errorf("Delete failed: %s", res.Status)
	}
	return nil
}
