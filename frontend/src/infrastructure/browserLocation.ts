import type { LocationPort } from '../domain/ports'

export class BrowserLocation implements LocationPort {
  here(): string {
    return window.location.pathname + window.location.search
  }
}
