---
name: orihon-reader
description: Guia o leitor desktop Orihon (Wails, Go, Vue, Inertia, MangaDex, estante/ficha/spread RTL). Use quando o usuário falar de mangá, capítulo, API, layout, estante ou leitura.
---

# Orihon reader

Usuário iniciante: cada termo técnico aponta para [docs/GLOSSARY.md](../../../docs/GLOSSARY.md).

## Workflow

1. Ler `memory/STATE.md`, `memory/LIBRARY.md`, `memory/LAST_SESSION.md`.
2. Confirmar a spec: busca, estante, ficha, spread RTL, só MangaDex v5.
3. Rodar `go test . ./internal/...` depois de mudar Go ou a fronteira da UI.
4. UI: paleta washi/vermelhão; dobras de orihon. Chrome: logo + busca + carimbo EN fixo; scrollbar; rodapé só em ficha/leitura. Print: `docs/images/estante.jpg`.
5. HTTP: User-Agent `Orihon/0.1`, só em `internal/infrastructure/mangadex`. Imagens de capítulo só com `baseUrl` do at-home, via `/media`. A interface não chama a API. O adaptador da tela é `frontend/src/infrastructure/inertiaCatalog.ts`.
6. `go test . ./internal/...` e `npm test` precisam continuar cobrindo CAP-7.

## Comandos

```bash
npm install
npm test
go test . ./internal/...
wails dev
```

## Fora

Manga Plus, deviceSecret, paywall, erotica/pornographic, Tauri.
