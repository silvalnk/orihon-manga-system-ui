package mangadex

import "testing"

func TestAllowed(t *testing.T) {
	ok := []string{
		"https://uploads.mangadex.org/covers/m1/file.256.jpg",
		"https://node.mangadex.network/data-saver/hh/a.jpg",
		"https://uploads.example/data-saver/hh/a.jpg",
	}
	for _, raw := range ok {
		if !Allowed(raw) {
			t.Fatal(raw)
		}
	}
	blocked := []string{
		"http://uploads.mangadex.org/covers/a/b.jpg",
		"https://127.0.0.1/data-saver/a.jpg",
		"https://localhost/data-saver/a.jpg",
		"https://evil.example/secret",
		"https://user:pass@uploads.mangadex.org/covers/a/b.jpg",
	}
	for _, raw := range blocked {
		if Allowed(raw) {
			t.Fatal(raw)
		}
	}
}
