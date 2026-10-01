import type { ReaderCommands } from '../domain/ports'
import type { ReadingMode } from '../domain/reading'

export function openChapter(reader: ReaderCommands, chapterId: string, mangaId: string): Promise<void> {
  return reader.openChapter(chapterId, mangaId)
}

export function turn(
  reader: ReaderCommands,
  mode: ReadingMode,
  chapterId: string,
  mangaId: string,
  index: number,
  count: number,
  forward: boolean,
): Promise<void> | null {
  const next = mode.step(index, count, forward)
  if (next === index) return null
  return reader.showSpread(chapterId, mangaId, next, mode.single)
}

export function flipSingle(
  reader: ReaderCommands,
  chapterId: string,
  mangaId: string,
  index: number,
  single: boolean,
): Promise<void> {
  return reader.showSpread(chapterId, mangaId, index, !single)
}

export function leaveReader(reader: ReaderCommands, mangaId: string): Promise<void> {
  if (mangaId) return reader.backToWork(mangaId)
  return reader.backToShelf()
}

export function leaveWork(reader: ReaderCommands): Promise<void> {
  return reader.backToShelf()
}
