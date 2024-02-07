import { GlobalModView, Mod } from './mod'

export type ordering =
  | 'last-updated'
  | 'newest'
  | 'most-downloaded'
  | 'top-rated'
export type section = 'mods' | 'asset-replacements' | 'libraries' | 'modpacks'

declare global {
  interface Window {
    init: () => Promise<string>
    getAddr: () => Promise<string>
    runGame: (profile: string) => Promise<string>

    getMods: (profileName: string) => Promise<Mod[]>
    openModDir: (profile: string, modPathName: string) => Promise<string>
    deleteMod: (profile: string, modPathName: string) => Promise<string>
    installModFromUrl: (profile: string, url: string) => Promise<string>

    getGlobalMods: (
      ordering: ordering,
      section: section,
      query: string,
      page: number,
    ) => Promise<GlobalModView[]>
    installMod: (
      profile: string,
      modAuthor: string,
      modName: string,
    ) => Promise<string>

    getProfiles: () => Promise<string[]>
    saveLastUsedProfile: (lastUsedProfile: string) => Promise<string>
    loadLastUsedProfile: () => Promise<string>
    renameProfile: (
      oldProfileName: string,
      newProfileName: string,
    ) => Promise<string>
    deleteProfile: (profileName: string) => Promise<string>
    createProfile: (profileName: string) => Promise<string>
  }
}
