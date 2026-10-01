package mangadex

import (
	"testing"

	"github.com/silvalnk/orihon/internal/domain"
)

func TestSearchURL(t *testing.T) {
	url := SearchURL("one piece", 0, 0, nil)
	if url[:len(API+"/manga?")] != API+"/manga?" {
		t.Fatalf("host: %s", url)
	}
	for _, want := range []string{"contentRating%5B%5D=safe", "contentRating%5B%5D=suggestive", "one%20piece", "limit=32", "offset=0"} {
		if !contains(url, want) {
			t.Fatalf("missing %s in %s", want, url)
		}
	}
	if contains(url, "pornographic") {
		t.Fatalf("adult rating leaked: %s", url)
	}
	if !contains(SearchURL("", 32, 32, nil), "offset=32") {
		t.Fatal("offset page")
	}
	if !contains(SearchURL("", 0, 0, nil), "followedCount") {
		t.Fatal("popular order")
	}
	if SearchLimit != 32 {
		t.Fatal(SearchLimit)
	}
}

func TestFeedAndPages(t *testing.T) {
	if !contains(FeedURL("abc", 0, nil), "limit=100") {
		t.Fatal("feed limit")
	}
	if AtHomeURL("abc") != API+"/at-home/server/abc" {
		t.Fatal(AtHomeURL("abc"))
	}
	if PageURL("https://x", "h", "1.jpg", true) != "https://x/data-saver/h/1.jpg" {
		t.Fatal("page url")
	}
	if !contains(MangaURL("m1"), "/manga/m1") || !contains(MangaURL("m1"), "cover_art") {
		t.Fatal(MangaURL("m1"))
	}
}

func TestParse(t *testing.T) {
	page, err := ParseMangaList([]byte(`{"total":99,"data":[{"id":"m1","attributes":{"title":{"en":"Paper Crane"},"description":{"en":"A quiet story."},"status":"ongoing"},"relationships":[{"type":"cover_art","attributes":{"fileName":"cover.png"}}]}]}`))
	if err != nil {
		t.Fatal(err)
	}
	if page.Items[0].Title != "Paper Crane" {
		t.Fatal(page.Items[0].Title)
	}
	if !contains(page.Items[0].Cover, "cover.png.256.jpg") {
		t.Fatal(page.Items[0].Cover)
	}
	if page.Total != 99 {
		t.Fatal(page.Total)
	}
	merged := MergeUnique([]domain.Manga{{ID: "m1", Title: "A"}}, []domain.Manga{{ID: "m1"}, {ID: "m2", Title: "B"}})
	if len(merged) != 2 || merged[1].ID != "m2" {
		t.Fatal(merged)
	}

	feed, err := ParseFeed([]byte(`{"data":[
	  {"id":"c1","attributes":{"chapter":"2","title":"Later","translatedLanguage":"en","pages":10}},
	  {"id":"c0","attributes":{"chapter":"1","title":"Start","translatedLanguage":"pt-br","pages":8}},
	  {"id":"cx","attributes":{"chapter":"9","externalUrl":"https://example.com","translatedLanguage":"en"}},
	  {"id":"cn","attributes":{"chapter":null,"title":null,"volume":null,"externalUrl":null,"pages":null,"translatedLanguage":"pt-br"}}
	]}`))
	if err != nil {
		t.Fatal(err)
	}
	if len(feed) != 3 || feed[0].ID != "cn" || feed[0].Chapter != "?" || feed[1].ID != "c0" || feed[2].ID != "c1" {
		t.Fatal(feed)
	}

	at, err := ParseAtHome([]byte(`{"baseUrl":"https://uploads.example","chapter":{"hash":"hh","dataSaver":["a.jpg","b.jpg"]}}`))
	if err != nil {
		t.Fatal(err)
	}
	if len(at.Pages) != 2 || !contains(at.Pages[0], "/data-saver/hh/a.jpg") {
		t.Fatal(at.Pages)
	}

	one, err := ParseManga([]byte(`{"data":{"id":"m9","attributes":{"title":{"pt-br":"Tsuru"},"status":"completed"}}}`))
	if err != nil || one.Title != "Tsuru" || one.ID != "m9" {
		t.Fatal(one, err)
	}
}

func contains(s, part string) bool {
	return len(s) >= len(part) && (s == part || len(part) == 0 || indexOf(s, part) >= 0)
}

func indexOf(s, part string) int {
	for i := 0; i+len(part) <= len(s); i++ {
		if s[i:i+len(part)] == part {
			return i
		}
	}
	return -1
}
