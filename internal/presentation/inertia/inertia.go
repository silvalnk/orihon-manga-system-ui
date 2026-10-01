package inertia

import (
	"encoding/json"
	"html"
	"net/http"
	"strings"
)

type Page struct {
	Component string         `json:"component"`
	Props     map[string]any `json:"props"`
	URL       string         `json:"url"`
	Version   string         `json:"version"`
}

type Renderer struct {
	Version  string
	Head     string
	HeadFunc func() string
}

func (r *Renderer) Render(w http.ResponseWriter, req *http.Request, component string, props map[string]any) {
	if props == nil {
		props = map[string]any{}
	}
	page := Page{
		Component: component,
		Props:     partial(req, component, props),
		URL:       req.URL.RequestURI(),
		Version:   r.Version,
	}
	if req.Header.Get("X-Inertia") == "true" {
		if sent := req.Header.Get("X-Inertia-Version"); sent != "" && sent != r.Version {
			w.Header().Set("X-Inertia-Location", req.URL.RequestURI())
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Inertia", "true")
		w.Header().Set("Vary", "X-Inertia")
		_ = json.NewEncoder(w).Encode(page)
		return
	}
	raw, err := json.Marshal(page)
	if err != nil {
		http.Error(w, "page", http.StatusInternalServerError)
		return
	}
	head := r.Head
	if r.HeadFunc != nil {
		head = r.HeadFunc()
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte("<!DOCTYPE html><html lang=\"en\"><head><script>if(location.hostname!=='127.0.0.1')location.replace('http://127.0.0.1:18765'+location.pathname+location.search)</script><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>Orihon</title>" + head + "</head><body><div id=\"app\" data-page=\"" + html.EscapeString(string(raw)) + "\"></div></body></html>"))
}

func partial(req *http.Request, component string, props map[string]any) map[string]any {
	if req.Header.Get("X-Inertia-Partial-Component") != component {
		return props
	}
	keys := req.Header.Get("X-Inertia-Partial-Data")
	if keys == "" {
		return props
	}
	keep := map[string]any{}
	for _, key := range strings.Split(keys, ",") {
		key = strings.TrimSpace(key)
		if value, ok := props[key]; ok {
			keep[key] = value
		}
	}
	return keep
}

func Redirect(w http.ResponseWriter, req *http.Request, loc string) {
	if req.Header.Get("X-Inertia") == "true" {
		w.Header().Set("X-Inertia", "true")
	}
	http.Redirect(w, req, loc, http.StatusSeeOther)
}
