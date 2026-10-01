export type Manga = {
  id: string
  title: string
  description: string
  status: string
  year: number
  cover: string
  coverFile: string
}

export type Chapter = {
  id: string
  chapter: string
  title: string
  lang: string
  pages: number
  volume: string
}
