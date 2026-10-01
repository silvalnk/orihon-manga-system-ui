import { describe, expect, it } from 'vitest'
import { spreadMode } from '../domain/reading'
import { turn } from './readChapter'
import type { ReaderCommands } from '../domain/ports'

function fakeReader(): ReaderCommands & { calls: unknown[] } {
  const calls: unknown[] = []
  return {
    calls,
    openChapter: (chapterId, mangaId) => {
      calls.push(['open', chapterId, mangaId])
      return Promise.resolve()
    },
    showSpread: (chapterId, mangaId, index, single) => {
      calls.push(['spread', chapterId, mangaId, index, single])
      return Promise.resolve()
    },
    backToWork: (mangaId) => {
      calls.push(['work', mangaId])
      return Promise.resolve()
    },
    backToShelf: () => {
      calls.push(['shelf'])
      return Promise.resolve()
    },
  }
}

describe('turn', () => {
  it('asks for the next spread index', async () => {
    const reader = fakeReader()
    await turn(reader, spreadMode, 'c1', 'm1', 0, 5, true)
    expect(reader.calls).toEqual([['spread', 'c1', 'm1', 2, false]])
  })

  it('stays put at the end', () => {
    const reader = fakeReader()
    expect(turn(reader, spreadMode, 'c1', 'm1', 4, 5, true)).toBeNull()
    expect(reader.calls).toEqual([])
  })
})
