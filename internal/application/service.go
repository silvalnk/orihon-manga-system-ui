package application

import (
	"context"

	"github.com/silvalnk/orihon/internal/domain"
)

type Service struct {
	Catalog domain.Catalog
	Library domain.Library
}

func New(catalog domain.Catalog, library domain.Library) *Service {
	return &Service{Catalog: catalog, Library: library}
}

type Shelf struct {
	Items       []domain.Manga
	Total       int
	Offset      int
	Query       string
	Lang        string
	FavoriteIDs []string
	Favorites   []domain.Manga
	Error       string
}

type Work struct {
	Manga     domain.Manga
	Chapters  []domain.Chapter
	Favorited bool
	Favorites []domain.Manga
	Lang      string
	Error     string
}

type Reading struct {
	MangaID   string
	ChapterID string
	URLs      []string
	Frame     domain.Frame
	Lang      string
	Error     string
}

func (s *Service) Shelf(ctx context.Context, query string, offset int, lang string) Shelf {
	view := Shelf{
		Items:       []domain.Manga{},
		Offset:      offset,
		Query:       query,
		Lang:        lang,
		FavoriteIDs: []string{},
		Favorites:   []domain.Manga{},
	}
	page, err := s.Catalog.Search(ctx, query, offset, []string{lang})
	if err != nil {
		view.Error = "catalog unavailable"
		return view
	}
	if page.Items == nil {
		page.Items = []domain.Manga{}
	}
	view.Items = page.Items
	view.Total = page.Total
	favorites, err := s.Library.Favorites(ctx)
	if err != nil {
		view.Error = "shelf unavailable"
		return view
	}
	if favorites == nil {
		favorites = []domain.Manga{}
	}
	view.Favorites = favorites
	for _, item := range favorites {
		view.FavoriteIDs = append(view.FavoriteIDs, item.ID)
	}
	return view
}

func (s *Service) Work(ctx context.Context, mangaID, lang string) Work {
	view := Work{Lang: lang, Chapters: []domain.Chapter{}, Favorites: []domain.Manga{}}
	manga, err := s.Catalog.Get(ctx, mangaID)
	if err != nil {
		view.Error = "work unavailable"
		return view
	}
	view.Manga = manga
	chapters, err := s.Catalog.Feed(ctx, mangaID, []string{"en"})
	if err != nil {
		view.Error = "chapters unavailable"
		return view
	}
	if chapters == nil {
		chapters = []domain.Chapter{}
	}
	view.Chapters = chapters
	favorites, err := s.Library.Favorites(ctx)
	if err != nil {
		view.Error = "shelf unavailable"
		return view
	}
	if favorites == nil {
		favorites = []domain.Manga{}
	}
	view.Favorites = favorites
	view.Favorited = domain.IsFavorite(favorites, mangaID)
	return view
}

func (s *Service) Read(ctx context.Context, chapterID, mangaID string, index int, single bool, lang string, remember bool) Reading {
	view := Reading{MangaID: mangaID, ChapterID: chapterID, URLs: []string{}, Lang: lang}
	urls, err := s.Catalog.Pages(ctx, chapterID)
	if err != nil {
		view.Error = "pages unavailable"
		view.Frame = domain.FrameAt(0, 0, single)
		return view
	}
	if urls == nil {
		urls = []string{}
	}
	view.URLs = urls
	if !remember {
		stored, err := s.Library.Progress(ctx, chapterID)
		if err == nil {
			index = stored
		}
	}
	view.Frame = domain.FrameAt(index, len(urls), single)
	if remember {
		_ = s.Library.SaveProgress(ctx, chapterID, view.Frame.Index)
	}
	return view
}

func (s *Service) ToggleFavorite(ctx context.Context, manga domain.Manga) (bool, error) {
	return s.Library.ToggleFavorite(ctx, manga)
}
