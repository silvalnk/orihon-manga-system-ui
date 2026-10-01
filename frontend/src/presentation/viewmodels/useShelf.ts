import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { openFold, requestMore } from '../../application/browseShelf'
import { catalog, emptyShelf, readProps, shelfSchema, useSession } from '../../composition/wire'
import { mergeFolds } from '../../domain/folds'
import type { Manga } from '../../domain/manga'
import type { ShelfCommands } from '../../domain/ports'

export function useShelf(shelf: ShelfCommands = catalog) {
  const props = readProps(shelfSchema, emptyShelf)
  const session = useSession()
  const folds = ref<Manga[]>([])
  const tail = ref<HTMLElement | null>(null)
  let observer: IntersectionObserver | null = null

  watch(
    () => [props.value.offset, props.value.items] as const,
    () => {
      const items = props.value.items
      folds.value = props.value.offset === 0 ? [...items] : mergeFolds(folds.value, items)
    },
    { immediate: true },
  )

  function more() {
    if (session.busy) return
    void requestMore(shelf, props.value.query, folds.value.length, props.value.total)
  }

  function onScroll(event: Event) {
    const el = event.target as HTMLElement
    if (el.scrollTop + el.clientHeight >= el.scrollHeight - 320) more()
  }

  onMounted(() => {
    observer = new IntersectionObserver(
      (entries) => {
        if (entries.some((entry) => entry.isIntersecting)) more()
      },
      { rootMargin: '320px' },
    )
    if (tail.value) observer.observe(tail.value)
  })

  onBeforeUnmount(() => observer?.disconnect())

  function open(manga: Manga) {
    void openFold(shelf, manga.id)
  }

  const status = computed(() => {
    if (props.value.error) return props.value.error
    const loaded = folds.value.length
    const total = props.value.total
    if (session.busy && loaded === 0) return props.value.query ? 'Searching…' : 'Fetching popular folds…'
    if (session.busy) return `${loaded} / ${total} folds · loading more…`
    if (loaded === 0) return 'No folds match.'
    if (loaded < total) return `${loaded} / ${total} folds`
    return `${loaded} folds on the shelf.`
  })

  return { props, folds, status, tail, onScroll, open }
}
