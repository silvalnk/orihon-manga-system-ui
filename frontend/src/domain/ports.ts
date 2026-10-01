import type { Manga } from './manga'

export interface ShelfCommands {
  showShelf(query: string): Promise<void>
  moreFolds(query: string, offset: number): Promise<void>
  openWork(id: string): Promise<void>
  toggleFavorite(manga: Manga, returnTo: string): Promise<void>
}

export interface ReaderCommands {
  openChapter(chapterId: string, mangaId: string): Promise<void>
  showSpread(chapterId: string, mangaId: string, index: number, single: boolean): Promise<void>
  backToWork(mangaId: string): Promise<void>
  backToShelf(): Promise<void>
}

export interface LanguageCommands {
  switchLanguage(lang: string, returnTo: string): Promise<void>
}

export interface LocationPort {
  here(): string
}
