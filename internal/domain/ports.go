package domain

import "context"

type Catalog interface {
	Search(ctx context.Context, query string, offset int, langs []string) (CatalogPage, error)
	Get(ctx context.Context, mangaID string) (Manga, error)
	Feed(ctx context.Context, mangaID string, langs []string) ([]Chapter, error)
	Pages(ctx context.Context, chapterID string) ([]string, error)
}

type Library interface {
	Favorites(ctx context.Context) ([]Manga, error)
	ToggleFavorite(ctx context.Context, manga Manga) (bool, error)
	Progress(ctx context.Context, chapterID string) (int, error)
	SaveProgress(ctx context.Context, chapterID string, index int) error
}
