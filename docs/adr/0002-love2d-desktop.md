# ADR 0002 — LÖVE como host desktop

## Status

Substituído por [ADR 0004](0004-wails-inertia.md). O texto abaixo é o registro da v1. O binário vigente é `wails dev`. A interface está no [ADR 0005](0005-frontend-layers.md).

## Contexto

O pedido é **somente Lua** e desktop com UI. Sem LÖVE, Lua puro não desenha janela.

## Decisão

LÖVE 11 executa o app. Módulos de domínio (`json`, `http`, `mangadex`, `persist`) não dependem de LÖVE para os testes.

## Consequências

- `lua tests/run.lua` roda sem GPU.
- `love .` é o binário da UI.
- HTTPS via `curl` (já presente no WSL), não via LuaRocks.
- Print de referência da estante: `docs/images/estante.jpg`.
