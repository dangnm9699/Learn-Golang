package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GetMovie is a function get movie information from API
func GetMovie(terms []string) (*Movie, error) {
	result := Movie{}
	// CODE HERE
	t := strings.Join(terms, " ")
	res, err := http.Get(APIURL + "t=" + url.QueryEscape(t))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("Get query by title = %s failed: %s", t, res.Status)
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()
	// END CODE
	return &result, nil
}
