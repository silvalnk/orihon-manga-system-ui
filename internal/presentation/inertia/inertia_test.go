package inertia

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFirstVisitIsHTML(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	(&Renderer{Version: "1", Head: "<script src=\"/assets/app.js\"></script>"}).Render(rec, req, "Shelf", map[string]any{"query": "paper"})
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "data-page=") || !strings.Contains(rec.Body.String(), "Shelf") {
		t.Fatal(rec.Code, rec.Body.String())
	}
}

func TestInertiaJSONAndPartial(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/?offset=32", nil)
	req.Header.Set("X-Inertia", "true")
	req.Header.Set("X-Inertia-Version", "1")
	req.Header.Set("X-Inertia-Partial-Component", "Shelf")
	req.Header.Set("X-Inertia-Partial-Data", "items")
	(&Renderer{Version: "1"}).Render(rec, req, "Shelf", map[string]any{"items": []string{"a"}, "total": 99})
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), `"items"`) || strings.Contains(rec.Body.String(), `"total"`) {
		t.Fatal(rec.Body.String())
	}
}

func TestVersionConflict(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Inertia", "true")
	req.Header.Set("X-Inertia-Version", "old")
	(&Renderer{Version: "1"}).Render(rec, req, "Shelf", nil)
	if rec.Code != http.StatusConflict || rec.Header().Get("X-Inertia-Location") != "/" {
		t.Fatal(rec.Code, rec.Header())
	}
}
