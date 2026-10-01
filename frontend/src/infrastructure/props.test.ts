import { describe, expect, it } from 'vitest'
import { emptyShelf, parseWith, shelfSchema } from './props'

describe('shelf props', () => {
  it('accepts a catalog page', () => {
    const page = parseWith(
      shelfSchema,
      {
        items: [
          { id: 'm1', title: 'Paper Crane', description: '', status: 'ongoing', year: 0, cover: '', coverFile: '' },
        ],
        total: 1,
        offset: 0,
        query: '',
        lang: 'pt-br',
        favoriteIds: ['m1'],
        error: '',
      },
      emptyShelf,
    )
    expect(page.items[0]?.title).toBe('Paper Crane')
    expect(page.lang).toBe('pt-br')
  })

  it('falls back when the payload drifts', () => {
    expect(parseWith(shelfSchema, { items: 'nope' }, emptyShelf).error).toBe('page invalid')
  })
})
