import type { GlobalModView } from '@/types/mod'
import { useState, useEffect } from 'react'

const useGetGlobalMods = (pageState: number, orderingState: string) => {
  const [globalMods, setGlobalMods] = useState<GlobalModView[]>([])
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<Error | null>(null)

  useEffect(() => {
    const fetchGlobalMods = async () => {
      setIsLoading(true)
      try {
        const newGlobalMods = await window.getGlobalMods(
          orderingState,
          pageState,
        )
        setGlobalMods(newGlobalMods)
      } catch (err) {
        setError(err as Error)
      } finally {
        setIsLoading(false)
      }
    }

    fetchGlobalMods()
  }, [pageState, orderingState])

  return { globalMods, isLoading, error }
}

export { useGetGlobalMods }
