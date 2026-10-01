# Contexto da sessão (Orihon)

Atualizado: 2026-09-30

GitHub: [silvalnk/orihon-manga-reader-system-ui](https://github.com/silvalnk/orihon-manga-reader-system-ui) · pasta local: `orihon_manga_ui/`

Print da estante: [`docs/images/estante.jpg`](../docs/images/estante.jpg)

## Estado

- Fonte: MangaDex API v5 (grátis, documentada)
- Janela: Wails. UI: Vue 3 + TypeScript + Inertia + Tailwind. Camadas em `frontend/src/{domain,application,infrastructure,composition,presentation}`
- Sessão global (Pinia): idioma e visita em andamento. O resto mora na página
- ADR vigente da janela: `docs/adr/0004-wails-inertia.md`. Interface: `docs/adr/0005-frontend-layers.md`. O 0002 ficou substituído
- Cabeçalho: selo + nome **Orihon**, busca, carimbo EN fixo
- Estante: grade + barra de rolagem à direita; `limit`/`offset` ao chegar no fim; **sem rodapé**
- Ficha: rodapé só **Back** (Esc também volta; o logo reseta a estante)
- Leitura: rodapé **Back / Prev / Next / Single**. A página da direita é a atual
- Idioma: sempre `en`. O carimbo EN não troca. Um `POST /lang` também grava `en`.
- Conteúdo: safe + suggestive
- Persistência: `library.json` + `progress.json` em `ORIHON_DATA_DIR` ou `~/.local/share/orihon/`
- HTTP: só em `internal/infrastructure/mangadex`. A interface pede `/media` para as páginas do capítulo
- CAP-7: `go test . ./internal/...` falha se `frontend/src` nomear o host da API
- CAP-8: chrome mínimo; rodapé só na ficha e na leitura
- Licença do código: MIT (`LICENSE`); obras da MangaDex continuam dos autores

## Como retomar numa sessão nova do Cursor

1. Abrir a pasta `orihon_manga_ui/`.
2. Ler `docs/GLOSSARY.md` se algum termo não for óbvio.
3. Ler `memory/STATE.md`, `memory/LIBRARY.md`, `memory/LAST_SESSION.md`.
4. Rodar `go test . ./internal/...` e, em `frontend/`, `npm test`.
5. Rodar `wails dev` para a janela.
