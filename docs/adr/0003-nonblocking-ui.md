# ADR 0003 — UI nunca espera a rede

## Status

A regra permanece. O mecanismo de threads do LÖVE, registrado no histórico, foi substituído pelos [ADR 0004](0004-wails-inertia.md) e [ADR 0005](0005-frontend-layers.md).

## Contexto

A leitura travava quando a janela esperava o download. A tela precisa continuar respondendo enquanto o catálogo responde.

## Decisão

- HTTP da MangaDex só em `internal/infrastructure/mangadex`, com User-Agent `Orihon/0.1`
- Páginas de capítulo saem por `/media`, com a URL limitada ao at-home
- `frontend/src/domain`, `frontend/src/application` e `frontend/src/presentation` não nomeiam o host da API e não importam o cliente HTTP
- O adaptador Inertia em `frontend/src/infrastructure/inertiaCatalog.ts` é quem pede a próxima página ao Go
- Contrato: `go test . ./internal/...` e `npm test` em `frontend/`

## Histórico (v1, LÖVE)

- `curl` só em `src/download_thread.lua`
- Fila `src/jobs.lua` com `jobs.N == 3`
- Prefetch spread/single `5`/`3` e `layout.images_per_frame == 3`
- Esses arquivos não existem mais no repositório

## Consequências

- A interface pede props ao processo Go. O Go fala com a MangaDex.
- Mudar a fronteira exige manter `go test . ./internal/...` verde.
