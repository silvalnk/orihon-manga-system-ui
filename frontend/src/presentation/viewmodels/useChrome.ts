import { ref, watch } from 'vue'
import { useEventListener } from '@vueuse/core'
import { searchShelf } from '../../application/browseShelf'
import { catalog, chromeSchema, readProps, useSession } from '../../composition/wire'
import type { ShelfCommands } from '../../domain/ports'

export function useChrome(shelf: ShelfCommands = catalog) {
  const props = readProps(chromeSchema, { lang: 'en' })
  const session = useSession()
  const q = ref('')

  watch(
    () => props.value.lang,
    (lang) => session.syncLang(lang),
    { immediate: true },
  )
  watch(
    () => props.value.query,
    (query) => {
      if (query !== undefined) q.value = query
    },
    { immediate: true },
  )

  function home() {
    q.value = ''
    void searchShelf(shelf, '')
  }

  function submit() {
    void searchShelf(shelf, q.value)
  }

  useEventListener(window, 'keydown', (event: KeyboardEvent) => {
    if (event.key !== '/' || event.target instanceof HTMLInputElement || event.target instanceof HTMLTextAreaElement)
      return
    event.preventDefault()
    document.getElementById('search')?.focus()
  })

  return { q, session, home, submit }
}
