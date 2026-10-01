# Plano

1. Domínio Go: obra, capítulo, portas, estratégia do spread RTL
2. Adaptador MangaDex (URL, parse, at-home, data-saver) e estante JSON
3. Casos de uso Go: estante, ficha, leitura, favorito, progresso
4. Inertia no processo Go (HTML na primeira visita, JSON com `X-Inertia`)
5. Interface nas mesmas camadas: domínio, casos de uso, adaptador Inertia, viewmodels e páginas Tailwind
6. Janela Wails; em dev a webview aponta para o servidor Go

Stack em uso: Go, Wails v2, Vue 3, TypeScript, Inertia, Vite, Tailwind CSS, Pinia, Zod, VueUse, Vitest, ESLint, Prettier, MangaDex API v5.

Fora da interface: Vue Router (o Inertia roteia), Axios/ofetch/TanStack Query (o Go faz o HTTP), PrimeVue (o chrome é próprio), Playwright (o contrato fica em `go test . ./internal/...` e `npm test`).
