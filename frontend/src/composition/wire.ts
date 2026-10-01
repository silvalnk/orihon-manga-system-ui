import { BrowserLocation } from '../infrastructure/browserLocation'
import { InertiaCatalog } from '../infrastructure/inertiaCatalog'
import { readProps } from '../infrastructure/props'
import { useSession } from '../infrastructure/session'

export const catalog = new InertiaCatalog()
export const location = new BrowserLocation()

export {
  chromeSchema,
  emptyReader,
  emptyShelf,
  emptyWork,
  readerSchema,
  shelfSchema,
  workSchema,
} from '../infrastructure/props'
export { readProps, useSession }
