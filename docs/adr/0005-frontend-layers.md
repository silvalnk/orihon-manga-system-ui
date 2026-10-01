# ADR 0005 — Interface em camadas

## Status

Aceito. Complementa [ADR 0004](0004-wails-inertia.md).

## Contexto

A janela já é Wails e o processo Go já separa domínio, casos de uso e adaptadores. A interface Vue ainda precisa da mesma disciplina: a vista não conhece o protocolo Inertia, e o domínio da tela não conhece Vue.

## Decisão

`frontend/src` segue a mesma ordem do Go:

- `domain` — obras, estratégia do spread (duas páginas ou uma) e a junção da estante
- `application` — buscar, abrir, virar a página, favoritar. O idioma da tela é inglês fixo.
- `infrastructure` — adaptador Inertia, leitura do endereço, sessão e validação das props com Zod
- `composition` — liga o adaptador (`wire.ts`)
- `presentation` — páginas, viewmodels e componentes (MVVM)

Padrões em uso: Strategy no spread, Adapter no Inertia, Factory no `wire.ts`. Pinia guarda só o que atravessa as páginas: idioma e visita em andamento. VueUse escuta o teclado. Tailwind desenha o washi e o vermelhão. Vitest cobre a estratégia e os casos de uso. ESLint e Prettier formatam `frontend/src`.

Ficam de fora, de propósito:

- Vue Router — o Inertia já escolhe `Shelf`, `Work` e `Reader`
- Axios, ofetch e TanStack Query — a tela não chama a MangaDex; o HTTP está no adaptador Go
- PrimeVue — o chrome é o orihon, não um kit de widgets
- Playwright — a janela é o Wails; o contrato HTTP continua em `go test . ./internal/...`

## Consequências

- `npm test` e `npm run build` em `frontend/`
- `go test . ./internal/...` falha se `domain` ou `application` importarem Vue, Pinia, Zod ou Inertia, e se `presentation` importar Inertia ou a pasta `infrastructure`
