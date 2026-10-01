import { computed } from 'vue'
import { useEventListener } from '@vueuse/core'
import { flipSingle, leaveReader, turn } from '../../application/readChapter'
import { catalog, emptyReader, readProps, readerSchema } from '../../composition/wire'
import { modeFor } from '../../domain/reading'
import type { ReaderCommands } from '../../domain/ports'

export function useReader(reader: ReaderCommands = catalog) {
  const props = readProps(readerSchema, emptyReader)

  function step(forward: boolean) {
    const frame = props.value.frame
    void turn(
      reader,
      modeFor(frame.single),
      props.value.chapterId,
      props.value.mangaId,
      frame.index,
      props.value.urls.length,
      forward,
    )
  }

  function click(event: MouseEvent) {
    const el = event.currentTarget as HTMLElement
    const x = event.clientX - el.getBoundingClientRect().left
    step(x > el.clientWidth / 2)
  }

  function single() {
    const frame = props.value.frame
    void flipSingle(reader, props.value.chapterId, props.value.mangaId, frame.index, frame.single)
  }

  function back() {
    void leaveReader(reader, props.value.mangaId)
  }

  useEventListener(window, 'keydown', (event: KeyboardEvent) => {
    if (event.target instanceof HTMLInputElement || event.target instanceof HTMLTextAreaElement) return
    if (event.key === 'ArrowRight') {
      event.preventDefault()
      step(true)
    }
    if (event.key === 'ArrowLeft') {
      event.preventDefault()
      step(false)
    }
    if (event.key === 'd' || event.key === 'D') single()
    if (event.key === 'Escape') back()
  })

  const modeLabel = computed(() => (props.value.frame.single ? 'Double' : 'Single'))
  const status = computed(() => {
    if (props.value.error) return props.value.error
    const count = props.value.urls.length
    if (count === 0) return 'Requesting pages…'
    return `${props.value.frame.index + 1} / ${count}`
  })

  return { props, step, click, single, back, modeLabel, status }
}
