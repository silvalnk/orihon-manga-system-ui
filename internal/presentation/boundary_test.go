package presentation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFrontendDependsInward(t *testing.T) {
	root := filepath.Join("..", "..", "frontend", "src")
	banned := map[string][]string{
		"domain":        {"@inertiajs", "from 'vue'", "from \"vue\"", "pinia", "zod"},
		"application":   {"@inertiajs", "from 'vue'", "from \"vue\"", "pinia", "zod"},
		"presentation":  {"@inertiajs", "from 'pinia'", "from 'zod'", "infrastructure/"},
	}
	for dir, words := range banned {
		err := filepath.WalkDir(filepath.Join(root, dir), func(path string, entry os.DirEntry, err error) error {
			if err != nil || entry.IsDir() || !strings.HasSuffix(path, ".ts") && !strings.HasSuffix(path, ".vue") {
				return err
			}
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			text := string(raw)
			for _, word := range words {
				if strings.Contains(text, word) {
					t.Fatalf("%s mentions %s", path, word)
				}
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestVueDoesNotCallTheCatalog(t *testing.T) {
	root := filepath.Join("..", "..", "frontend", "src")
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(raw), "mangadex.org") {
			t.Fatalf("%s names the catalog host", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
