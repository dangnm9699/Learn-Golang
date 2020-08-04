package xkcd

import (
	"strconv"
	"strings"
)

// SearchComics queries the GitHub issue tracker.
func SearchComics(terms []string, list *ComicList) []*Comic {
	result := []*Comic{}
	// CODE HERE
	for _, comic := range list.Comics {
		if hasTerms(comic, terms) {
			result = append(result, comic)
		}
	}
	//
	return result
}

func hasTerms(comic *Comic, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(comic.Month, term) {
			continue
		} else if strings.Contains(comic.Link, term) {
			continue
		} else if strings.Contains(comic.Year, term) {
			continue
		} else if strings.Contains(comic.News, term) {
			continue
		} else if strings.Contains(comic.SafeTitle, term) {
			continue
		} else if strings.Contains(comic.Transcript, term) {
			continue
		} else if strings.Contains(comic.Alt, term) {
			continue
		} else if strings.Contains(comic.Img, term) {
			continue
		} else if strings.Contains(comic.Title, term) {
			continue
		} else if strings.Contains(comic.Day, term) {
			continue
		} else if i, err := strconv.Atoi(term); err == nil && i == comic.Num {
			continue
		} else {
			return false
		}
	}
	return true
}
