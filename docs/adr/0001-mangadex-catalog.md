# ADR 0001 — MangaDex como catálogo

## Contexto

Orihon precisa de um catálogo gratuito, documentado e sem credenciais. APIs de apps oficiais de editora costumam exigir handshake de dispositivo e tokens de sessão.

## Decisão

Orihon usa só a [MangaDex API v5](https://api.mangadex.org/docs/03-manga/search/), pública e gratuita, com User-Agent identificável.

## Consequências

- Catálogo = obras publicadas na MangaDex.
- Sem paywall, sem extração de secret de celular.
- A metáfora de orihon permanece. A janela é Wails ([ADR 0004](0004-wails-inertia.md)); a interface em camadas está no [ADR 0005](0005-frontend-layers.md).
