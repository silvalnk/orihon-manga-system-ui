import { router } from '@inertiajs/vue3'
import type { Manga } from '../domain/manga'
import type { LanguageCommands, ReaderCommands, ShelfCommands } from '../domain/ports'
import { useSession } from './session'

type Query = Record<string, string | number>

export class InertiaCatalog implements ShelfCommands, ReaderCommands, LanguageCommands {
  showShelf(query: string): Promise<void> {
    return this.visit('get', '/', { q: query, offset: 0 })
  }

  moreFolds(query: string, offset: number): Promise<void> {
    return this.visit(
      'get',
      '/',
      { q: query, offset },
      {
        only: ['items', 'total', 'offset', 'favoriteIds', 'error'],
        preserveState: true,
        preserveScroll: true,
      },
    )
  }

  openWork(id: string): Promise<void> {
    return this.visit('get', `/works/${id}`)
  }

  toggleFavorite(manga: Manga, returnTo: string): Promise<void> {
    return this.visit(
      'post',
      `/works/${manga.id}/favorite`,
      {
        title: manga.title,
        cover: manga.cover,
        coverFile: manga.coverFile,
        return: returnTo,
      },
      { preserveScroll: true, preserveState: true },
    )
  }

  openChapter(chapterId: string, mangaId: string): Promise<void> {
    return this.visit('get', `/read/${chapterId}`, { manga: mangaId })
  }

  showSpread(chapterId: string, mangaId: string, index: number, single: boolean): Promise<void> {
    const query: Query = { manga: mangaId, index }
    if (single) query.single = 1
    return this.visit('get', `/read/${chapterId}`, query, { preserveScroll: true })
  }

  backToWork(mangaId: string): Promise<void> {
    return this.visit('get', `/works/${mangaId}`)
  }

  backToShelf(): Promise<void> {
    return this.visit('get', '/')
  }

  switchLanguage(lang: string, returnTo: string): Promise<void> {
    return this.visit('post', '/lang', { lang, return: returnTo })
  }

  private visit(
    method: 'get' | 'post',
    url: string,
    data: Query = {},
    extra: Record<string, unknown> = {},
  ): Promise<void> {
    const session = useSession()
    session.setBusy(true)
    return new Promise((resolve) => {
      const options = {
        ...extra,
        onFinish: () => {
          session.setBusy(false)
          resolve()
        },
      }
      if (method === 'get') router.get(url, data, options)
      else router.post(url, data, options)
    })
  }
}
