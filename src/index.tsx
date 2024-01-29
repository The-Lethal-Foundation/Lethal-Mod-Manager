import React from 'react'

import '@/styles/index.css'
import App from './App'
import * as serviceWorker from './serviceWorker'
import { createRoot } from 'react-dom/client'
import type { Mod } from './types/mod'

// Functions bound from Lorca are available on the window object
declare global {
  interface Window {
    init: () => Promise<string>
    getMods: (profileName: string) => Promise<Mod[]>
    openModDir: (profile: string, modPathName: string) => Promise<string>
    deleteMod: (profile: string, modPathName: string) => Promise<string>

    getProfiles: () => Promise<string[]>
    saveLastUsedProfile: (lastUsedProfile: string) => Promise<string>
    loadLastUsedProfile: () => Promise<string>
  }
}

const root = createRoot(document.getElementById('root') as HTMLElement)
root.render(<App />)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()
