import type { LanguageCommands } from '../domain/ports'

const languages = new Set(['en', 'pt-br'])

export function switchLanguage(language: LanguageCommands, lang: string, returnTo: string): Promise<void> | null {
  if (!languages.has(lang)) return null
  return language.switchLanguage(lang, returnTo)
}
