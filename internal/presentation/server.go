package presentation

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/silvalnk/orihon/internal/application"
	"github.com/silvalnk/orihon/internal/domain"
	"github.com/silvalnk/orihon/internal/presentation/inertia"
)

type MediaSource interface {
	Fetch(ctx context.Context, rawURL string) (string, []byte, error)
}

type Handler struct {
	App   *application.Service
	Page  *inertia.Renderer
	Media MediaSource
}

func New(app *application.Service, page *inertia.Renderer, media MediaSource) *Handler {
	return &Handler{App: app, Page: page, Media: media}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch {
	case path == "/" && r.Method == http.MethodGet:
		h.shelf(w, r)
	case path == "/lang" && r.Method == http.MethodPost:
		h.lang(w, r)
	case path == "/media" && r.Method == http.MethodGet:
		h.media(w, r)
	case strings.HasPrefix(path, "/works/") && strings.HasSuffix(path, "/favorite") && r.Method == http.MethodPost:
		h.favorite(w, r)
	case strings.HasPrefix(path, "/works/") && r.Method == http.MethodGet:
		h.work(w, r)
	case strings.HasPrefix(path, "/read/") && strings.HasSuffix(path, "/progress") && r.Method == http.MethodPost:
		h.progress(w, r)
	case strings.HasPrefix(path, "/read/") && r.Method == http.MethodGet:
		h.read(w, r)
	case strings.HasPrefix(path, "/assets/"):
		http.FileServer(http.Dir("frontend/dist")).ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

func IsAppRoute(r *http.Request) bool {
	if r.Header.Get("X-Inertia") == "true" {
		return true
	}
	path := r.URL.Path
	if path == "/" || path == "/lang" || path == "/media" {
		return true
	}
	return strings.HasPrefix(path, "/works/") || strings.HasPrefix(path, "/read/")
}

func (h *Handler) shelf(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	offset := atoi(query.Get("offset"))
	lang := langOf(r)
	view := h.App.Shelf(r.Context(), query.Get("q"), offset, lang)
	h.Page.Render(w, r, "Shelf", map[string]any{
		"items":       presentMangaList(view.Items),
		"total":       view.Total,
		"offset":      view.Offset,
		"query":       view.Query,
		"lang":        view.Lang,
		"favoriteIds": view.FavoriteIDs,
		"favorites":   presentMangaList(view.Favorites),
		"error":       view.Error,
	})
}

func (h *Handler) work(w http.ResponseWriter, r *http.Request) {
	id := segment(r.URL.Path, "/works/")
	view := h.App.Work(r.Context(), id, langOf(r))
	h.Page.Render(w, r, "Work", map[string]any{
		"manga":     presentManga(view.Manga),
		"chapters":  view.Chapters,
		"favorited": view.Favorited,
		"favorites": presentMangaList(view.Favorites),
		"lang":      view.Lang,
		"error":     view.Error,
	})
}

func (h *Handler) read(w http.ResponseWriter, r *http.Request) {
	id := segment(r.URL.Path, "/read/")
	query := r.URL.Query()
	_, remember := query["index"]
	index := atoi(query.Get("index"))
	single := query.Get("single") == "1"
	view := h.App.Read(r.Context(), id, query.Get("manga"), index, single, langOf(r), remember)
	h.Page.Render(w, r, "Reader", map[string]any{
		"mangaId":   view.MangaID,
		"chapterId": view.ChapterID,
		"urls":      mediaURLs(view.URLs),
		"frame":     view.Frame,
		"lang":      view.Lang,
		"error":     view.Error,
	})
}

func (h *Handler) favorite(w http.ResponseWriter, r *http.Request) {
	form, err := fields(r)
	if err != nil {
		http.Error(w, "form", http.StatusBadRequest)
		return
	}
	id := strings.TrimSuffix(segment(r.URL.Path, "/works/"), "/favorite")
	_, err = h.App.ToggleFavorite(r.Context(), domain.Manga{
		ID:        id,
		Title:     form["title"],
		Cover:     form["cover"],
		CoverFile: form["coverFile"],
	})
	if err != nil {
		http.Error(w, "shelf", http.StatusInternalServerError)
		return
	}
	inertia.Redirect(w, r, back(r, form, "/"))
}

func (h *Handler) progress(w http.ResponseWriter, r *http.Request) {
	form, err := fields(r)
	if err != nil {
		http.Error(w, "form", http.StatusBadRequest)
		return
	}
	id := strings.TrimSuffix(segment(r.URL.Path, "/read/"), "/progress")
	q := url.Values{}
	q.Set("manga", form["manga"])
	q.Set("index", form["index"])
	if form["single"] == "1" {
		q.Set("single", "1")
	}
	inertia.Redirect(w, r, "/read/"+id+"?"+q.Encode())
}

func (h *Handler) lang(w http.ResponseWriter, r *http.Request) {
	form, err := fields(r)
	if err != nil {
		http.Error(w, "form", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "orihon_lang",
		Value:    "en",
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 365,
		SameSite: http.SameSiteLaxMode,
	})
	inertia.Redirect(w, r, back(r, form, "/"))
}

func (h *Handler) media(w http.ResponseWriter, r *http.Request) {
	raw := r.URL.Query().Get("url")
	if h.Media == nil {
		http.NotFound(w, r)
		return
	}
	kind, body, err := h.Media.Fetch(r.Context(), raw)
	if errors.Is(err, domain.ErrMediaBlocked) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "media", http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", kind)
	w.Header().Set("Cache-Control", "private, max-age=300")
	_, _ = w.Write(body)
}

func langOf(*http.Request) string {
	return "en"
}

func atoi(raw string) int {
	n, err := strconv.Atoi(raw)
	if err != nil || n < 0 {
		return 0
	}
	return n
}

func segment(path, prefix string) string {
	rest := strings.TrimPrefix(path, prefix)
	rest = strings.Trim(rest, "/")
	if i := strings.IndexByte(rest, '/'); i >= 0 {
		return rest[:i]
	}
	return rest
}

func back(r *http.Request, form map[string]string, fallback string) string {
	if loc := safePath(form["return"]); loc != "" {
		return loc
	}
	ref := r.Referer()
	if ref == "" {
		return fallback
	}
	parsed, err := url.Parse(ref)
	if err != nil || parsed.Host != r.Host {
		return fallback
	}
	if parsed.Path == "" {
		return fallback
	}
	return parsed.RequestURI()
}

func safePath(loc string) string {
	if !strings.HasPrefix(loc, "/") || strings.HasPrefix(loc, "//") || strings.Contains(loc, "://") {
		return ""
	}
	return loc
}

func fields(r *http.Request) (map[string]string, error) {
	out := map[string]string{}
	if strings.Contains(r.Header.Get("Content-Type"), "json") {
		var raw map[string]any
		if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
			return nil, err
		}
		for key, value := range raw {
			switch typed := value.(type) {
			case string:
				out[key] = typed
			case float64:
				out[key] = strconv.FormatInt(int64(typed), 10)
			case bool:
				if typed {
					out[key] = "1"
				}
			}
		}
		return out, nil
	}
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	for key, values := range r.Form {
		if len(values) > 0 {
			out[key] = values[0]
		}
	}
	return out, nil
}

func mediaURLs(urls []string) []string {
	out := make([]string, 0, len(urls))
	for _, raw := range urls {
		out = append(out, mediaURL(raw))
	}
	return out
}

func presentManga(item domain.Manga) domain.Manga {
	item.Cover = mediaURL(item.Cover)
	return item
}

func presentMangaList(items []domain.Manga) []domain.Manga {
	out := make([]domain.Manga, 0, len(items))
	for _, item := range items {
		out = append(out, presentManga(item))
	}
	return out
}

func mediaURL(raw string) string {
	if raw == "" || strings.HasPrefix(raw, "/media?") {
		return raw
	}
	return "/media?url=" + url.QueryEscape(raw)
}
