package application

import (
	"context"
	"testing"

	"github.com/silvalnk/orihon/internal/domain"
)

type catalog struct {
	page domain.CatalogPage
	one  domain.Manga
	feed []domain.Chapter
	urls []string
}

func (c catalog) Search(context.Context, string, int, []string) (domain.CatalogPage, error) {
	return c.page, nil
}
func (c catalog) Get(context.Context, string) (domain.Manga, error) { return c.one, nil }
func (c catalog) Feed(context.Context, string, []string) ([]domain.Chapter, error) {
	return c.feed, nil
}
func (c catalog) Pages(context.Context, string) ([]string, error) { return c.urls, nil }

type shelf struct {
	favorites []domain.Manga
	progress  map[string]int
}

func (s *shelf) Favorites(context.Context) ([]domain.Manga, error) { return s.favorites, nil }
func (s *shelf) ToggleFavorite(_ context.Context, manga domain.Manga) (bool, error) {
	s.favorites = append(s.favorites, manga)
	return true, nil
}
func (s *shelf) Progress(_ context.Context, id string) (int, error) {
	return s.progress[id], nil
}
func (s *shelf) SaveProgress(_ context.Context, id string, index int) error {
	if s.progress == nil {
		s.progress = map[string]int{}
	}
	s.progress[id] = index
	return nil
}

func TestReadRestoresProgress(t *testing.T) {
	lib := &shelf{progress: map[string]int{"c1": 2}}
	svc := New(catalog{urls: []string{"a", "b", "c", "d"}}, lib)
	view := svc.Read(context.Background(), "c1", "m1", 0, false, "en", false)
	if view.Frame.Index != 2 || view.Frame.Right != 2 || !view.Frame.HasLeft {
		t.Fatal(view.Frame)
	}
	view = svc.Read(context.Background(), "c1", "m1", 0, false, "en", true)
	if view.Frame.Index != 0 || lib.progress["c1"] != 0 {
		t.Fatal(view.Frame, lib.progress)
	}
}

func TestShelfFavoriteIDs(t *testing.T) {
	svc := New(catalog{page: domain.CatalogPage{Total: 1, Items: []domain.Manga{{ID: "m1"}}}}, &shelf{favorites: []domain.Manga{{ID: "m1"}}})
	view := svc.Shelf(context.Background(), "", 0, "pt-br")
	if len(view.FavoriteIDs) != 1 || view.FavoriteIDs[0] != "m1" || view.Total != 1 {
		t.Fatal(view)
	}
}
