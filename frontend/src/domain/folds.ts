import type { Manga } from './manga'

export function mergeFolds(current: Manga[], incoming: Manga[]): Manga[] {
  const seen = new Set(current.map((item) => item.id))
  const next = [...current]
  for (const item of incoming) {
    if (!seen.has(item.id)) {
      next.push(item)
      seen.add(item.id)
    }
  }
  return next
}

export function nextOffset(loaded: number, total: number): number | null {
  if (loaded >= total) return null
  return loaded
}
