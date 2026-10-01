import { useEventListener } from '@vueuse/core'
import { leaveWork, openChapter } from '../../application/readChapter'
import { markFavorite, openFold } from '../../application/browseShelf'
import { catalog, emptyWork, location, readProps, workSchema } from '../../composition/wire'
import type { Chapter, Manga } from '../../domain/manga'
import type { LocationPort, ReaderCommands, ShelfCommands } from '../../domain/ports'

export function useWork(
  reader: ReaderCommands = catalog,
  shelf: ShelfCommands = catalog,
  where: LocationPort = location,
) {
  const props = readProps(workSchema, emptyWork)

  function back() {
    void leaveWork(reader)
  }

  function open(chapter: Chapter) {
    void openChapter(reader, chapter.id, props.value.manga.id)
  }

  function toggle() {
    const manga = props.value.manga
    if (!manga.id) return
    void markFavorite(shelf, manga, where.here())
  }

  function openFavorite(manga: Manga) {
    void openFold(shelf, manga.id)
  }

  useEventListener(window, 'keydown', (event: KeyboardEvent) => {
    if (event.key === 'Escape') back()
  })

  return { props, back, open, toggle, openFavorite }
}
