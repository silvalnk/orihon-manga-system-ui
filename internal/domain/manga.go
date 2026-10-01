package domain

import "errors"

var ErrMediaBlocked = errors.New("media blocked")

type Manga struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Year        int    `json:"year"`
	Cover       string `json:"cover"`
	CoverFile   string `json:"coverFile"`
}

type Chapter struct {
	ID      string `json:"id"`
	Chapter string `json:"chapter"`
	Title   string `json:"title"`
	Lang    string `json:"lang"`
	Pages   int    `json:"pages"`
	Volume  string `json:"volume"`
}

type CatalogPage struct {
	Items []Manga `json:"items"`
	Total int     `json:"total"`
}

func IsFavorite(favorites []Manga, id string) bool {
	for _, item := range favorites {
		if item.ID == id {
			return true
		}
	}
	return false
}
