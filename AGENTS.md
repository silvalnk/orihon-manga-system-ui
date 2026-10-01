# AGENTS.md

Briefing **independente de sessão** para o Cursor. O histórico do chat é opcional; estes arquivos não.

## Sempre (sessão nova)

1. Se o usuário for iniciante ou aparecer termo técnico, ler [`docs/GLOSSARY.md`](docs/GLOSSARY.md).
2. Ler [`memory/STATE.md`](memory/STATE.md), [`memory/LIBRARY.md`](memory/LIBRARY.md), [`memory/LAST_SESSION.md`](memory/LAST_SESSION.md).
3. Seguir [`.cursor/skills/orihon-reader/SKILL.md`](.cursor/skills/orihon-reader/SKILL.md).
4. Spec manda: [`.specify/SPEC.md`](.specify/SPEC.md). Estado vivo: [`.specify/CONTEXT.md`](.specify/CONTEXT.md).
5. Nunca integrar Manga Plus, scanlation pirata, ou API não documentada. Fonte = **MangaDex API v5** (pública, grátis).
6. Não baixar nem exibir conteúdo `erotica` / `pornographic` na v1.
7. **CAP-7:** a interface não chama a MangaDex. HTTP só em `internal/infrastructure/mangadex`. `frontend/src/domain` e `frontend/src/application` não importam Vue, Inertia, Pinia nem Zod. `go test . ./internal/...` e `npm test` têm de continuar verdes.
8. Depois de mudar comportamento, atualizar `.specify/CONTEXT.md`, `memory/STATE.md` e o glossário se um termo novo nascer. Se a estante mudar de cara, atualizar `docs/images/estante.jpg` e o README.
9. Quando o usuário pedir **commit**, seguir [`.specify/COMMITS.md`](.specify/COMMITS.md) (`✨ feat:` / `📝 docs:` / …, **mensagens em inglês**).

## Idioma

- Markdown (incluindo este arquivo): **português**
- Código, erros, identificadores, commits: **inglês**

## Produto

**Orihon** é um leitor desktop de mangá. A janela é Wails; o processo é Go e a interface Vue seguem Clean Architecture. A apresentação é MVVM (páginas, viewmodels, componentes) com Tailwind. O Inertia entrega as props. Pinia guarda o idioma e a visita em andamento. A UI imita um orihon (livro-acordeão): estante de dobras, ficha da obra, leitura em *spread* RTL. Pasta local: `orihon_manga_ui/`. Repositório: [silvalnk/orihon-manga-reader-system-ui](https://github.com/silvalnk/orihon-manga-reader-system-ui). Print da estante: [`docs/images/estante.jpg`](docs/images/estante.jpg).

Chrome: cabeçalho (logo, busca, carimbo EN fixo); barra de rolagem na estante/ficha; rodapé só na ficha e na leitura. O catálogo pede só `en`.

## Comandos

```bash
npm install
npm test
go test . ./internal/...
wails dev
```

## Não adicionar na v1

Manga Plus / Shueisha unofficial API, paywall bypass, Tauri, contas de usuário MangaDex, upload, comentários, conteúdo adulto explícito.
