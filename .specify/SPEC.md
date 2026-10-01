# Spec — Orihon v1

## Objetivo

Leitor desktop de mangá, com UI de orihon (papel, dobras, RTL) e catálogo via API gratuita oficial. A janela é Wails; o processo é Go; a interface é Vue, TypeScript e Inertia. Sessão persistente no Cursor via SDD + BMad.

## Capacidades

- **CAP-1** — Catálogo MangaDex com páginas de **32** (`limit`/`offset`). A primeira página no arranque; as seguintes entram ao **rolar para baixo** até `total`.
- **CAP-2** — Estante persistente (favoritos em JSON local)
- **CAP-3** — Ficha da obra: capa, sinopse, até **100** capítulos no feed
- **CAP-4** — Ler páginas via MangaDex@Home (`/at-home/server/{chapterId}`), qualidade `data-saver`
- **CAP-5** — Spread RTL (duas páginas) com teclado/mouse; progresso por capítulo
- **CAP-6** — `AGENTS.md` + `memory/` + skill Cursor + glossário
- **CAP-7** — A interface não chama a MangaDex. HTTP só no adaptador Go. Contrato em `go test . ./internal/...`.
- **CAP-8** — Chrome da UI, como em `docs/images/estante.jpg`: cabeçalho (selo, **Orihon**, busca, carimbo **EN** fixo); coluna **Favorites**; dobras com capa, título e sinopse; contagem `n / total folds`; barra de rolagem; rodapé só na ficha (**Back**) e na leitura (**Back / Prev / Next / Single|Double** e a página atual).

## Travas

- User-Agent identificável (`Orihon/0.1`)
- Só `contentRating` `safe` e `suggestive`
- Sem credenciais; sem OAuth
- Favoritos e progresso em `ORIHON_DATA_DIR` ou `~/.local/share/orihon/`
- Páginas de capítulo e capas passam pelo processo Go (`/media`). O capítulo fica limitado ao at-home. A capa passa pelo mesmo caminho porque o CDN troca a imagem quando o pedido vem com o agente de um navegador
- Vue, viewmodels, domínio e casos de uso da interface não nomeiam o host da API e não importam o cliente HTTP
- `frontend/src/domain` e `frontend/src/application` não importam Vue, Pinia, Zod nem Inertia

## Fora de escopo

APIs não documentadas, Tauri, APIs pagas, login MangaDex, modo adulto, multiplayer.

## Sucesso

`go test . ./internal/...` verde (CAP-7); `npm test` em `frontend/`; CI `.github/workflows/test.yml`; `wails dev` abre a estante no mesmo espírito de `docs/images/estante.jpg`; sessão nova lê `memory/` sem o chat antigo; iniciante explica *orihon*, *spread* e *MangaDex* pelo glossário.
