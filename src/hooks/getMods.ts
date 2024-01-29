import type { Mod } from '@/types/mod'
import { useState, useEffect } from 'react'

const useGetMods = (profileName: string) => {
  const [mods, setMods] = useState<Mod[]>([])
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<Error | null>(null)

  useEffect(() => {
    const fetchMods = async () => {
      setIsLoading(true)
      try {
        const modsData = await window.getMods(profileName)
        setMods(modsData)
      } catch (err) {
        setError(err as Error)
      } finally {
        setIsLoading(false)
      }
    }

    if (profileName) {
      fetchMods()
    }
  }, [profileName])

  return { mods, isLoading, error }
}

export default useGetMods
