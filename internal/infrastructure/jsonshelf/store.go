package jsonshelf

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/silvalnk/orihon/internal/domain"
)

type Store struct {
	Dir string
	mu  sync.Mutex
}

func Dir() string {
	if env := os.Getenv("ORIHON_DATA_DIR"); env != "" {
		return env
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ".orihon"
	}
	return filepath.Join(home, ".local", "share", "orihon")
}

func Open(dir string) (*Store, error) {
	if dir == "" {
		dir = Dir()
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return &Store{Dir: dir}, nil
}

func (s *Store) Favorites(context.Context) ([]domain.Manga, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.loadFavorites()
}

func (s *Store) ToggleFavorite(_ context.Context, manga domain.Manga) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	favorites, err := s.loadFavorites()
	if err != nil {
		return false, err
	}
	for i, item := range favorites {
		if item.ID == manga.ID {
			favorites = append(favorites[:i], favorites[i+1:]...)
			return false, s.saveFavorites(favorites)
		}
	}
	favorites = append(favorites, domain.Manga{
		ID: manga.ID, Title: manga.Title, Cover: manga.Cover, CoverFile: manga.CoverFile,
	})
	return true, s.saveFavorites(favorites)
}

func (s *Store) Progress(_ context.Context, chapterID string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	progress, err := s.loadProgress()
	if err != nil {
		return 0, err
	}
	return progress[chapterID], nil
}

func (s *Store) SaveProgress(_ context.Context, chapterID string, index int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	progress, err := s.loadProgress()
	if err != nil {
		return err
	}
	progress[chapterID] = index
	return s.write("progress.json", progress)
}

func (s *Store) loadFavorites() ([]domain.Manga, error) {
	var data struct {
		Favorites []domain.Manga `json:"favorites"`
	}
	if err := s.read("library.json", &data); err != nil {
		return nil, err
	}
	if data.Favorites == nil {
		data.Favorites = []domain.Manga{}
	}
	return data.Favorites, nil
}

func (s *Store) saveFavorites(favorites []domain.Manga) error {
	return s.write("library.json", map[string]any{"favorites": favorites})
}

func (s *Store) loadProgress() (map[string]int, error) {
	progress := map[string]int{}
	if err := s.read("progress.json", &progress); err != nil {
		return nil, err
	}
	if progress == nil {
		progress = map[string]int{}
	}
	return progress, nil
}

func (s *Store) read(name string, dest any) error {
	raw, err := os.ReadFile(filepath.Join(s.Dir, name))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(raw) == 0 {
		return nil
	}
	return json.Unmarshal(raw, dest)
}

func (s *Store) write(name string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(s.Dir, name), raw, 0o644)
}
