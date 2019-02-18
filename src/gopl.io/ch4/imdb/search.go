package imdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Search OMDb API
func SearchMovie(titles []string) (*SearchResult, error) {
	queryEscape := url.QueryEscape(strings.Join(titles, " "))
	fmt.Println("Query:%s", queryEscape)

	reqURL := fmt.Sprintf("%s/?apikey=%s&s=%s", ApiURL, ApiKey, queryEscape)
	fmt.Println("ReqURL:", reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query field: %s", resp.Status)
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
