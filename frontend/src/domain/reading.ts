export type Frame = {
  index: number
  right: number
  left: number
  hasLeft: boolean
  single: boolean
}

export interface ReadingMode {
  readonly single: boolean
  step(index: number, count: number, forward: boolean): number
}

function step(index: number, count: number, single: boolean, forward: boolean): number {
  if (count <= 0) return 0
  let delta = single ? 1 : 2
  if (!forward) delta = -delta
  const next = index + delta
  if (next < 0) return 0
  if (next >= count) return index
  return next
}

export const spreadMode: ReadingMode = {
  single: false,
  step: (index, count, forward) => step(index, count, false, forward),
}

export const singleMode: ReadingMode = {
  single: true,
  step: (index, count, forward) => step(index, count, true, forward),
}

export function modeFor(single: boolean): ReadingMode {
  return single ? singleMode : spreadMode
}
