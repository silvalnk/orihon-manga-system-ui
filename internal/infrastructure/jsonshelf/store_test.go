package jsonshelf

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/silvalnk/orihon/internal/domain"
)

func TestToggleFavorite(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "orihon-test-data")
	store, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	kept, err := store.ToggleFavorite(ctx, domain.Manga{ID: "m1", Title: "Paper Crane"})
	if err != nil || !kept {
		t.Fatal(kept, err)
	}
	favorites, err := store.Favorites(ctx)
	if err != nil || !domain.IsFavorite(favorites, "m1") {
		t.Fatal(favorites, err)
	}
	kept, err = store.ToggleFavorite(ctx, domain.Manga{ID: "m1", Title: "Paper Crane"})
	if err != nil || kept {
		t.Fatal(kept, err)
	}
	favorites, err = store.Favorites(ctx)
	if err != nil || domain.IsFavorite(favorites, "m1") {
		t.Fatal(favorites, err)
	}
	if store.Dir != dir {
		t.Fatal(store.Dir)
	}
}

func TestProgress(t *testing.T) {
	store, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if err := store.SaveProgress(ctx, "c0", 4); err != nil {
		t.Fatal(err)
	}
	index, err := store.Progress(ctx, "c0")
	if err != nil || index != 4 {
		t.Fatal(index, err)
	}
}
