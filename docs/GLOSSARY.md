# Glossário Orihon (para iniciante)

> **Única fonte de definições deste projeto.**  
> Não leia tudo de uma vez. Escolha uma seção e 3–5 termos.

| Marca | Significado |
|-------|-------------|
| **Neste projeto** | Aparece no código, na UI ou em `memory/` |
| **Só conceito** | Você pode ouvir no mundo do mangá; a v1 não implementa |

---

## 1. O produto

### Orihon

**Em uma frase:** livro japonês em sanfona, dobrado como um biombo.

**Explicação:** Em vez de páginas soltas grampeadas, o orihon é uma tira contínua dobrada. Este app usa essa metáfora: cada obra é uma tira de dobras.

**Analogia:** Um mapa de estrada que você abre em zigue-zague.

**Erro comum:** Achar que é um leitor web ou o app oficial de uma editora.

**Neste projeto:** o nome do produto; pasta local `orihon_manga_ui/`; repo [silvalnk/orihon-manga-reader-system-ui](https://github.com/silvalnk/orihon-manga-reader-system-ui).

### Estante (`shelf`)

**Em uma frase:** tela inicial com busca e favoritos, como uma prateleira de orihons.

**Neste projeto:** página `Shelf`. Print: `docs/images/estante.jpg`.

### Barra de rolagem

**Em uma frase:** trilha à direita da grade (e da lista de capítulos) para ver o resto das dobras.

**Neste projeto:** a grade rola; ao chegar no fim, o Go pede a página seguinte (`limit`/`offset`).

### Rodapé

**Em uma frase:** faixa de botões só quando a tela precisa de navegação extra.

**Neste projeto:** a estante **não** tem rodapé. A ficha tem **Back**. A leitura tem **Back / Prev / Next / Single**. `layout.shows_footer`.

### Carimbo EN

**Em uma frase:** marca de inglês no canto direito do cabeçalho.

**Neste projeto:** o catálogo e os capítulos pedem só `en`. O carimbo fica aceso e não troca.

### Selo

**Em uma frase:** círculo vermelhão à esquerda do nome Orihon, com a letra M.

**Neste projeto:** o logo. Não guarda favorito.

### Estrela

**Em uma frase:** marca na ficha que guarda a obra na coluna Favorites.

**Neste projeto:** botão circular vermelhão à direita do título. Preenchida, a obra está na estante.

### Ficha / dobra (`fold`)

**Em uma frase:** a obra aberta no meio: capa, texto e capítulos empilhados como dobras.

**Neste projeto:** página `Work`.

### Spread

**Em uma frase:** duas páginas lado a lado, no sentido de leitura do mangá (direita → esquerda).

**Analogia:** Abrir um caderno no meio e ver a página da direita primeiro.

**Neste projeto:** tela `spread`; tecla `D` força uma página só.

### RTL

**Em uma frase:** right-to-left — lê-se da direita para a esquerda.

**Neste projeto:** clique na metade **direita** avança; na **esquerda**, volta.

---

## 2. Fonte dos mangás

### MangaDex

**Em uma frase:** site/catálogo com API pública e gratuita para listar obras e capítulos.

**Explicação:** Não é a Shueisha nem o MANGA Plus. Editores e scanlators oficiais/parceiros publicam lá com regras próprias.

**Neste projeto:** única API da v1. Docs: https://api.mangadex.org/docs/

### MangaDex@Home

**Em uma frase:** rede que entrega as **imagens** das páginas. A URL muda e vale pouco tempo.

**Neste projeto:** `GET /at-home/server/{chapterId}` no adaptador Go. A página usa o `baseUrl` devolvido. A interface recebe a imagem por `/media`.

### data-saver

**Em uma frase:** versão comprimida da página, mais leve.

**Neste projeto:** qualidade padrão da leitura (banda de WSL/residencial).

### contentRating

**Em uma frase:** classificação da obra (`safe`, `suggestive`, `erotica`, `pornographic`).

**Neste projeto:** só `safe` e `suggestive`.

### feed

**Em uma frase:** lista de capítulos de um mangá.

**Neste projeto:** `GET /manga/{id}/feed`.

---

## 3. Ferramentas

### LÖVE (Love2D)

**Em uma frase:** motor para fazer app/jogo em Lua, com janela, desenho e input.

**Neste projeto:** a v1 usava LÖVE. A janela atual é Wails.

### Lua

**Em uma frase:** linguagem da v1. O app atual é Go e TypeScript.

### Wails

**Em uma frase:** ferramenta que abre uma janela desktop e mostra a interface web lá dentro.

**Neste projeto:** `wails dev` abre o Orihon.

### Inertia

**Em uma frase:** a janela pede uma página ao Go e o Vue troca o miolo, sem a interface falar com a API do catálogo.

**Neste projeto:** a primeira resposta é HTML; as seguintes, com o cabeçalho `X-Inertia`, são JSON.

### MVVM

**Em uma frase:** a vista desenha, o viewmodel reage aos cliques, e os dados vêm de fora.

**Neste projeto:** `frontend/src/presentation/pages` desenha; `frontend/src/presentation/viewmodels` reage; as props validadas são os dados. O adaptador Inertia fica em `frontend/src/infrastructure`.

### Clean Architecture

**Em uma frase:** o miolo não conhece a janela nem a rede; quem fala com o mundo fica na borda.

**Neste projeto:** a mesma ordem no Go (`internal/`) e na interface (`frontend/src/`): domínio, casos de uso, adaptadores, apresentação. Ver [ADR 0005](adr/0005-frontend-layers.md).

### Tailwind

**Em uma frase:** o visual escrito em classes no próprio componente.

**Neste projeto:** washi, tinta e vermelhão em `frontend/src/styles/app.css` e nas páginas.

### Pinia

**Em uma frase:** um lugar só para o que todas as telas precisam lembrar.

**Neste projeto:** idioma e se uma visita ao Go está em andamento. A lista da estante fica na página, não aqui.

### Zod

**Em uma frase:** confere se os dados que chegaram têm o formato combinado.

**Neste projeto:** as props do Inertia são lidas em `frontend/src/infrastructure/props.ts`.

### SDD

**Em uma frase:** Spec-Driven Development — a spec diz o que o código deve fazer.

**Neste projeto:** `.specify/SPEC.md` manda se divergir do código.

### BMad

**Em uma frase:** método para o contexto não morrer entre chats (clarify → plan → build → learn).

**Neste projeto:** `docs/bmad/PROCESS.md`.

### Rede fora da tela

**Em uma frase:** a janela não espera o catálogo. Quem fala com a MangaDex é o processo Go.

**Neste projeto:** o adaptador Go fala com a MangaDex. A interface só pede páginas ao processo local. `go test . ./internal/...` cobre esse contrato.

---

## 4. Só conceito (fora da v1)

### MANGA Plus / deviceSecret

App oficial da Shueisha; secret do celular para sessão paga. **Orihon não usa.**
