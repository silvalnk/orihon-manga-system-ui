package main

import (
	"embed"
	"net"
	"net/http"
	"os"

	"github.com/silvalnk/orihon/internal/composition"
	"github.com/silvalnk/orihon/internal/presentation"
	"github.com/silvalnk/orihon/internal/presentation/inertia"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// WSL's WebKit paints the window but drops clicks and scrolling unless compositing is off.
	_ = os.Setenv("WEBKIT_DISABLE_COMPOSITING_MODE", "1")
	_ = os.Setenv("WEBKIT_DISABLE_DMABUF_RENDERER", "1")

	service, catalog, err := composition.Wire("")
	if err != nil {
		println("Error:", err.Error())
		return
	}
	pages := &inertia.Renderer{Version: "1", HeadFunc: func() string { return presentation.Head(assets) }}
	handler := presentation.New(service, pages, catalog)
	if ln, err := net.Listen("tcp", "127.0.0.1:18765"); err != nil {
		println("Error:", err.Error())
	} else {
		go func() { _ = http.Serve(ln, handler) }()
	}

	app := NewApp()
	err = wails.Run(&options.App{
		Title:  "Orihon — accordion manga reader",
		Width:  1280,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: assetserver.ChainMiddleware(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if presentation.IsAppRoute(r) {
						handler.ServeHTTP(w, r)
						return
					}
					next.ServeHTTP(w, r)
				})
			}),
		},
		BackgroundColour: &options.RGBA{R: 244, G: 239, B: 228, A: 255},
		OnStartup:        app.startup,
		Bind:             []interface{}{app},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
