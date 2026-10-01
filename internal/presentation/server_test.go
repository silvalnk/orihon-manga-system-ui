package presentation

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/silvalnk/orihon/internal/application"
	"github.com/silvalnk/orihon/internal/domain"
	"github.com/silvalnk/orihon/internal/presentation/inertia"
)

type fakeCatalog struct{}

func (fakeCatalog) Search(context.Context, string, int, []string) (domain.CatalogPage, error) {
	return domain.CatalogPage{Total: 1, Items: []domain.Manga{{ID: "m1", Title: "Paper Crane"}}}, nil
}
func (fakeCatalog) Get(context.Context, string) (domain.Manga, error) {
	return domain.Manga{ID: "m1", Title: "Paper Crane"}, nil
}
func (fakeCatalog) Feed(context.Context, string, []string) ([]domain.Chapter, error) {
	return []domain.Chapter{{ID: "c1", Chapter: "1"}}, nil
}
func (fakeCatalog) Pages(context.Context, string) ([]string, error) {
	return []string{"https://uploads.example/data-saver/hh/a.jpg", "https://uploads.example/data-saver/hh/b.jpg"}, nil
}

type fakeShelf struct{ favorites []domain.Manga }

func (s *fakeShelf) Favorites(context.Context) ([]domain.Manga, error) { return s.favorites, nil }
func (s *fakeShelf) ToggleFavorite(_ context.Context, manga domain.Manga) (bool, error) {
	s.favorites = append(s.favorites, manga)
	return true, nil
}
func (s *fakeShelf) Progress(context.Context, string) (int, error)   { return 0, nil }
func (s *fakeShelf) SaveProgress(context.Context, string, int) error { return nil }

type fakeMedia struct{}

func (fakeMedia) Fetch(_ context.Context, raw string) (string, []byte, error) {
	if strings.Contains(raw, "evil.example") {
		return "", nil, domain.ErrMediaBlocked
	}
	return "image/jpeg", []byte("img"), nil
}

func newHandler() *Handler {
	return New(application.New(fakeCatalog{}, &fakeShelf{}), &inertia.Renderer{Version: "1"}, fakeMedia{})
}

func TestShelfAndReaderStayOffTheNetwork(t *testing.T) {
	h := newHandler()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "Paper Crane") {
		t.Fatal(rec.Body.String())
	}
	rec = httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/read/c1?manga=m1&index=0", nil)
	req.Header.Set("X-Inertia", "true")
	req.Header.Set("X-Inertia-Version", "1")
	h.ServeHTTP(rec, req)
	body := rec.Body.String()
	if !strings.Contains(body, "/media?url=") || strings.Contains(body, "api.mangadex.org") {
		t.Fatal(body)
	}
	if !strings.Contains(body, `"right":0`) || !strings.Contains(body, `"hasLeft":true`) {
		t.Fatal(body)
	}
}

func TestMediaBlocksOpenProxy(t *testing.T) {
	h := newHandler()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/media?url=https://evil.example/secret", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatal(rec.Code)
	}
}

func TestLangCookie(t *testing.T) {
	h := newHandler()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/lang", strings.NewReader("lang=pt-br&return=/"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusSeeOther {
		t.Fatal(rec.Code)
	}
	if cookie := rec.Header().Get("Set-Cookie"); !strings.Contains(cookie, "orihon_lang=en") {
		t.Fatal(cookie)
	}
}

func TestLangStaysEnglish(t *testing.T) {
	h := newHandler()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "orihon_lang", Value: "pt-br"})
	req.Header.Set("X-Inertia", "true")
	req.Header.Set("X-Inertia-Version", "1")
	h.ServeHTTP(rec, req)
	if !strings.Contains(rec.Body.String(), `"lang":"en"`) {
		t.Fatal(rec.Body.String())
	}
}
