import type { Manga } from '../domain/manga'
import type { ShelfCommands } from '../domain/ports'
import { nextOffset } from '../domain/folds'

export function searchShelf(shelf: ShelfCommands, query: string): Promise<void> {
  return shelf.showShelf(query)
}

export function requestMore(shelf: ShelfCommands, query: string, loaded: number, total: number): Promise<void> | null {
  const offset = nextOffset(loaded, total)
  if (offset === null) return null
  return shelf.moreFolds(query, offset)
}

export function openFold(shelf: ShelfCommands, id: string): Promise<void> {
  return shelf.openWork(id)
}

export function markFavorite(shelf: ShelfCommands, manga: Manga, returnTo: string): Promise<void> {
  return shelf.toggleFavorite(manga, returnTo)
}
