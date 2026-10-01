import { computed, type ComputedRef } from 'vue'
import { usePage } from '@inertiajs/vue3'
import { z } from 'zod'
import type { Chapter, Manga } from '../domain/manga'
import type { Frame } from '../domain/reading'

const mangaSchema = z.object({
  id: z.string(),
  title: z.string(),
  description: z.string(),
  status: z.string(),
  year: z.number(),
  cover: z.string(),
  coverFile: z.string(),
})

const chapterSchema = z.object({
  id: z.string(),
  chapter: z.string(),
  title: z.string(),
  lang: z.string(),
  pages: z.number(),
  volume: z.string(),
})

const frameSchema = z.object({
  index: z.number(),
  right: z.number(),
  left: z.number(),
  hasLeft: z.boolean(),
  single: z.boolean(),
})

export const shelfSchema = z.object({
  items: z.array(mangaSchema),
  total: z.number(),
  offset: z.number(),
  query: z.string(),
  lang: z.enum(['en', 'pt-br']),
  favoriteIds: z.array(z.string()),
  favorites: z.array(mangaSchema).default([]),
  error: z.string(),
})

export const workSchema = z.object({
  manga: mangaSchema,
  chapters: z.array(chapterSchema),
  favorited: z.boolean(),
  favorites: z.array(mangaSchema).default([]),
  lang: z.enum(['en', 'pt-br']),
  error: z.string(),
})

export const readerSchema = z.object({
  mangaId: z.string(),
  chapterId: z.string(),
  urls: z.array(z.string()),
  frame: frameSchema,
  lang: z.enum(['en', 'pt-br']),
  error: z.string(),
})

export const chromeSchema = z.object({
  lang: z.enum(['en', 'pt-br']),
  query: z.string().optional(),
})

export type ShelfPage = z.infer<typeof shelfSchema>
export type WorkPage = z.infer<typeof workSchema>
export type ReaderPage = z.infer<typeof readerSchema>
export type ChromePage = z.infer<typeof chromeSchema>

export const emptyShelf: ShelfPage = {
  items: [],
  total: 0,
  offset: 0,
  query: '',
  lang: 'en',
  favoriteIds: [],
  favorites: [],
  error: 'page invalid',
}

export const emptyWork: WorkPage = {
  manga: { id: '', title: '', description: '', status: '', year: 0, cover: '', coverFile: '' },
  chapters: [],
  favorited: false,
  favorites: [],
  lang: 'en',
  error: 'page invalid',
}

export const emptyReader: ReaderPage = {
  mangaId: '',
  chapterId: '',
  urls: [],
  frame: { index: 0, right: 0, left: 0, hasLeft: false, single: false },
  lang: 'en',
  error: 'page invalid',
}

export function parseWith<T>(schema: z.ZodType<T>, input: unknown, fallback: T): T {
  const parsed = schema.safeParse(input)
  return parsed.success ? parsed.data : fallback
}

export function readProps<T>(schema: z.ZodType<T>, fallback: T): ComputedRef<T> {
  const page = usePage()
  return computed(() => parseWith(schema, page.props, fallback))
}

export type { Manga, Chapter, Frame }
