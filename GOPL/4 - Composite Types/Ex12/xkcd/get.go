package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetComic is GET method
func GetComic(number string) (*Comic, error) {
	par := fmt.Sprintf("%s/info.0.json", number)
	res, err := http.Get(APIURL + par)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", res.Status)
	}
	var result Comic
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
