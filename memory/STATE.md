# Estado

- Tela padrão: estante (busca + populares); print em `docs/images/estante.jpg`
- Repo: [silvalnk/orihon-manga-system-ui](https://github.com/silvalnk/orihon-manga-system-ui) · pasta `orihon_manga_ui/`
- Janela: Wails. Go e a interface em Clean Architecture. Páginas Vue + Tailwind. Inertia roteia `Shelf`, `Work`, `Reader`
- Pinia: idioma (`en`) e visita em andamento
- Cabeçalho: nome **Orihon** + busca + carimbo EN fixo
- Estante: barra de rolagem; sem botões no rodapé
- Ficha: **Back** · Leitura: **Back / Prev / Next / Single|Double** e a página atual
- Clique no logo volta à estante (limpa a busca)
- API: `https://api.mangadex.org`, User-Agent `Orihon/0.1`, só no adaptador Go
- Spread: duas páginas (RTL); `D` = página única. A mesma conta está em `internal/domain/reading.go` e `frontend/src/domain/reading.ts`
- Testes: `npm test` e `go test . ./internal/...` na raiz do repositório
