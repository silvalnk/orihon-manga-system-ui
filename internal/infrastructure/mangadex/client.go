package mangadex

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/silvalnk/orihon/internal/domain"
)

const (
	API         = "https://api.mangadex.org"
	Covers      = "https://uploads.mangadex.org/covers"
	SearchLimit = 32
	FeedLimit   = 100
	UserAgent   = "Orihon/0.1"
)

type AtHome struct {
	BaseURL string   `json:"baseUrl"`
	Hash    string   `json:"hash"`
	Pages   []string `json:"pages"`
}

func SearchURL(query string, limit, offset int, langs []string) string {
	if limit <= 0 {
		limit = SearchLimit
	}
	if offset < 0 {
		offset = 0
	}
	q := url.Values{}
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(offset))
	q.Add("includes[]", "cover_art")
	q.Add("contentRating[]", "safe")
	q.Add("contentRating[]", "suggestive")
	for _, lang := range normalizeLangs(langs) {
		q.Add("availableTranslatedLanguage[]", lang)
	}
	if strings.TrimSpace(query) != "" {
		q.Set("title", query)
		q.Set("order[relevance]", "desc")
	} else {
		q.Set("order[followedCount]", "desc")
	}
	return API + "/manga?" + encode(q)
}

func MangaURL(mangaID string) string {
	q := url.Values{}
	q.Add("includes[]", "cover_art")
	return API + "/manga/" + mangaID + "?" + encode(q)
}

func FeedURL(mangaID string, limit int, langs []string) string {
	if limit <= 0 {
		limit = FeedLimit
	}
	q := url.Values{}
	q.Set("limit", strconv.Itoa(limit))
	for _, lang := range normalizeLangs(langs) {
		q.Add("translatedLanguage[]", lang)
	}
	q.Set("order[chapter]", "asc")
	q.Set("includeEmptyPages", "0")
	q.Set("includeFuturePublishAt", "0")
	q.Set("includeExternalUrl", "0")
	return API + "/manga/" + mangaID + "/feed?" + encode(q)
}

func AtHomeURL(chapterID string) string {
	return API + "/at-home/server/" + chapterID
}

func CoverURL(mangaID, fileName string) string {
	if mangaID == "" || fileName == "" {
		return ""
	}
	return Covers + "/" + mangaID + "/" + fileName + ".256.jpg"
}

func PageURL(baseURL, hash, filename string, saver bool) string {
	quality := "data"
	if saver {
		quality = "data-saver"
	}
	return fmt.Sprintf("%s/%s/%s/%s", strings.TrimRight(baseURL, "/"), quality, hash, filename)
}

func normalizeLangs(langs []string) []string {
	out := make([]string, 0, len(langs))
	for _, lang := range langs {
		lang = strings.TrimSpace(lang)
		if lang != "" {
			out = append(out, lang)
		}
	}
	if len(out) == 0 {
		return []string{"en", "pt-br"}
	}
	return out
}

func encode(q url.Values) string {
	return strings.ReplaceAll(q.Encode(), "+", "%20")
}

type listed struct {
	ID            string `json:"id"`
	Attributes    attrs  `json:"attributes"`
	Relationships []struct {
		Type       string `json:"type"`
		Attributes struct {
			FileName string `json:"fileName"`
		} `json:"attributes"`
	} `json:"relationships"`
}

func ParseMangaList(payload []byte) (domain.CatalogPage, error) {
	var raw struct {
		Total int      `json:"total"`
		Data  []listed `json:"data"`
	}
	if err := json.Unmarshal(payload, &raw); err != nil {
		return domain.CatalogPage{}, err
	}
	page := domain.CatalogPage{Total: raw.Total, Items: []domain.Manga{}}
	for _, item := range raw.Data {
		page.Items = append(page.Items, mangaFrom(item))
	}
	if raw.Total == 0 {
		page.Total = len(page.Items)
	}
	return page, nil
}

func ParseManga(payload []byte) (domain.Manga, error) {
	var raw struct {
		Data listed `json:"data"`
	}
	if err := json.Unmarshal(payload, &raw); err != nil {
		return domain.Manga{}, err
	}
	if raw.Data.ID == "" {
		return domain.Manga{}, fmt.Errorf("mangadex manga missing")
	}
	return mangaFrom(raw.Data), nil
}

func mangaFrom(item listed) domain.Manga {
	coverFile := ""
	for _, rel := range item.Relationships {
		if rel.Type == "cover_art" {
			coverFile = rel.Attributes.FileName
		}
	}
	year := 0
	if item.Attributes.Year != nil {
		year = *item.Attributes.Year
	}
	return domain.Manga{
		ID:          item.ID,
		Title:       firstLocalized(item.Attributes.Title, "Untitled"),
		Description: firstLocalized(item.Attributes.Description, ""),
		Status:      item.Attributes.Status,
		Year:        year,
		Cover:       CoverURL(item.ID, coverFile),
		CoverFile:   coverFile,
	}
}

func MergeUnique(dst, src []domain.Manga) []domain.Manga {
	seen := map[string]bool{}
	for _, item := range dst {
		if item.ID != "" {
			seen[item.ID] = true
		}
	}
	for _, item := range src {
		if item.ID != "" && !seen[item.ID] {
			dst = append(dst, item)
			seen[item.ID] = true
		}
	}
	return dst
}

// looseString accepts the nulls MangaDex sends for title, volume and externalUrl.
type looseString string

func (s *looseString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = ""
		return nil
	}
	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return err
	}
	*s = looseString(text)
	return nil
}

func ParseFeed(payload []byte) ([]domain.Chapter, error) {
	var raw struct {
		Data []struct {
			ID         string `json:"id"`
			Attributes struct {
				Chapter            looseString `json:"chapter"`
				Title              looseString `json:"title"`
				TranslatedLanguage looseString `json:"translatedLanguage"`
				Pages              *int        `json:"pages"`
				Volume             looseString `json:"volume"`
				ExternalURL        looseString `json:"externalUrl"`
			} `json:"attributes"`
		} `json:"data"`
	}
	if err := json.Unmarshal(payload, &raw); err != nil {
		return nil, err
	}
	list := []domain.Chapter{}
	for _, item := range raw.Data {
		if item.Attributes.ExternalURL != "" {
			continue
		}
		chapter := string(item.Attributes.Chapter)
		if chapter == "" {
			chapter = "?"
		}
		pages := 0
		if item.Attributes.Pages != nil {
			pages = *item.Attributes.Pages
		}
		list = append(list, domain.Chapter{
			ID:      item.ID,
			Chapter: chapter,
			Title:   string(item.Attributes.Title),
			Lang:    string(item.Attributes.TranslatedLanguage),
			Pages:   pages,
			Volume:  string(item.Attributes.Volume),
		})
	}
	sort.SliceStable(list, func(i, j int) bool {
		return chapterNumber(list[i].Chapter) < chapterNumber(list[j].Chapter)
	})
	return list, nil
}

func ParseAtHome(payload []byte) (AtHome, error) {
	var raw struct {
		BaseURL string `json:"baseUrl"`
		Chapter struct {
			Hash      string   `json:"hash"`
			Data      []string `json:"data"`
			DataSaver []string `json:"dataSaver"`
		} `json:"chapter"`
	}
	if err := json.Unmarshal(payload, &raw); err != nil {
		return AtHome{}, err
	}
	files := raw.Chapter.DataSaver
	saver := true
	if len(files) == 0 {
		files = raw.Chapter.Data
		saver = false
	}
	pages := make([]string, 0, len(files))
	for _, name := range files {
		pages = append(pages, PageURL(raw.BaseURL, raw.Chapter.Hash, name, saver))
	}
	return AtHome{BaseURL: raw.BaseURL, Hash: raw.Chapter.Hash, Pages: pages}, nil
}

type attrs struct {
	Title       map[string]string `json:"title"`
	Description map[string]string `json:"description"`
	Status      string            `json:"status"`
	Year        *int              `json:"year"`
}

func firstLocalized(values map[string]string, fallback string) string {
	if len(values) == 0 {
		return fallback
	}
	for _, key := range []string{"en", "pt-br", "ja"} {
		if values[key] != "" {
			return values[key]
		}
	}
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return fallback
}

func chapterNumber(chapter string) float64 {
	n, err := strconv.ParseFloat(chapter, 64)
	if err != nil {
		return 0
	}
	return n
}

type Client struct {
	HTTP *http.Client
}

func NewClient() *Client {
	return &Client{HTTP: &http.Client{Timeout: 20 * time.Second}}
}

func (c *Client) get(ctx context.Context, rawURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("mangadex %d", res.StatusCode)
	}
	return body, nil
}

func (c *Client) Search(ctx context.Context, query string, offset int, langs []string) (domain.CatalogPage, error) {
	body, err := c.get(ctx, SearchURL(query, SearchLimit, offset, langs))
	if err != nil {
		return domain.CatalogPage{}, err
	}
	return ParseMangaList(body)
}

func (c *Client) Get(ctx context.Context, mangaID string) (domain.Manga, error) {
	body, err := c.get(ctx, MangaURL(mangaID))
	if err != nil {
		return domain.Manga{}, err
	}
	return ParseManga(body)
}

func (c *Client) Feed(ctx context.Context, mangaID string, langs []string) ([]domain.Chapter, error) {
	body, err := c.get(ctx, FeedURL(mangaID, FeedLimit, langs))
	if err != nil {
		return nil, err
	}
	return ParseFeed(body)
}

func (c *Client) Pages(ctx context.Context, chapterID string) ([]string, error) {
	at, err := c.AtHome(ctx, chapterID)
	if err != nil {
		return nil, err
	}
	return at.Pages, nil
}

func (c *Client) AtHome(ctx context.Context, chapterID string) (AtHome, error) {
	body, err := c.get(ctx, AtHomeURL(chapterID))
	if err != nil {
		return AtHome{}, err
	}
	return ParseAtHome(body)
}

func (c *Client) Fetch(ctx context.Context, rawURL string) (string, []byte, error) {
	if !Allowed(rawURL) {
		return "", nil, domain.ErrMediaBlocked
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(io.LimitReader(res.Body, 20<<20))
	if err != nil {
		return "", nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", nil, fmt.Errorf("mangadex %d", res.StatusCode)
	}
	kind := res.Header.Get("Content-Type")
	if kind == "" {
		kind = "application/octet-stream"
	}
	return kind, body, nil
}
