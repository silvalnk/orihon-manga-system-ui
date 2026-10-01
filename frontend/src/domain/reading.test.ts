import { modeFor, spreadMode, singleMode } from './reading'
import { describe, expect, it } from 'vitest'

describe('reading strategy', () => {
  it('advances a spread by two pages', () => {
    expect(spreadMode.step(0, 5, true)).toBe(2)
    expect(spreadMode.step(2, 5, false)).toBe(0)
    expect(spreadMode.step(4, 5, true)).toBe(4)
  })

  it('advances a single page by one', () => {
    expect(singleMode.step(1, 5, true)).toBe(2)
    expect(singleMode.single).toBe(true)
  })

  it('picks the mode from the flag', () => {
    expect(modeFor(false)).toBe(spreadMode)
    expect(modeFor(true)).toBe(singleMode)
  })
})
