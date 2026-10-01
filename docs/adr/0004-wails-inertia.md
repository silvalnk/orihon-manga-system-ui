# ADR 0004 — Wails, Go e Inertia

## Status

Aceito. Substitui [ADR 0002](0002-love2d-desktop.md).

## Contexto

O leitor continua o mesmo: estante, ficha, spread da direita para a esquerda, catálogo MangaDex. O pedido passou a ser uma janela desktop em Go, com Vue, TypeScript e Inertia.

## Decisão

A janela é [Wails](https://wails.io/) v2. A interface é Vue 3 + TypeScript e fala com o processo Go pelo protocolo Inertia. No Go, domínio e casos de uso não importam Wails nem a MangaDex; o cliente HTTP e o JSON local são adaptadores; o spread é uma estratégia de índices. Na interface, a mesma separação está no [ADR 0005](0005-frontend-layers.md).

## Consequências

- `go test . ./internal/...` cobre URL, parse, estante, spread e as respostas Inertia.
- Em desenvolvimento a webview abre `http://127.0.0.1:18765` e o Go entrega o pacote já compilado em `frontend/dist`. O WebKitGTK não completa o grafo de módulos do Vite.
- No binário, o mesmo roteador entra como middleware do asset server.
- A tela não chama a MangaDex. A regra do [ADR 0003](0003-nonblocking-ui.md) permanece; o mecanismo de threads do LÖVE não.
- As camadas da interface estão no [ADR 0005](0005-frontend-layers.md).
