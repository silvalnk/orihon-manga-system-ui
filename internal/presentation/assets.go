package presentation

import (
	"encoding/json"
	"fmt"
	"html"
	"io/fs"
	"os"
	"strings"
)

const viteDevHead = `<script type="module" src="http://127.0.0.1:5173/@vite/client"></script><script type="module" src="http://127.0.0.1:5173/src/main.ts"></script>`

func Head(fsys fs.FS) string {
	if os.Getenv("frontenddevserverurl") != "" || os.Getenv("devserver") != "" {
		if head := AssetHead(os.DirFS("frontend/dist")); head != viteDevHead {
			return head
		}
	}
	return AssetHead(fsys)
}

func AssetHead(fsys fs.FS) string {
	raw, err := readManifest(fsys)
	if err != nil {
		return viteDevHead
	}
	var manifest map[string]struct {
		File    string   `json:"file"`
		CSS     []string `json:"css"`
		IsEntry bool     `json:"isEntry"`
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		return viteDevHead
	}
	var tags strings.Builder
	for _, entry := range manifest {
		if !entry.IsEntry || entry.File == "" {
			continue
		}
		for _, css := range entry.CSS {
			fmt.Fprintf(&tags, `<link rel="stylesheet" href="/%s" />`, html.EscapeString(strings.TrimPrefix(css, "/")))
		}
		fmt.Fprintf(&tags, `<script type="module" src="/%s"></script>`, html.EscapeString(strings.TrimPrefix(entry.File, "/")))
	}
	if tags.Len() == 0 {
		return viteDevHead
	}
	return tags.String()
}

func readManifest(fsys fs.FS) ([]byte, error) {
	for _, name := range []string{
		"frontend/dist/manifest.json",
		"frontend/dist/.vite/manifest.json",
		"manifest.json",
	} {
		raw, err := fs.ReadFile(fsys, name)
		if err == nil {
			return raw, nil
		}
	}
	return nil, fs.ErrNotExist
}
