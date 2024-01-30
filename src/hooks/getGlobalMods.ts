import { ordering, section } from '@/types/global'
import type { GlobalModView } from '@/types/mod'
import { useState, useEffect } from 'react'

const useGetGlobalMods = (
  pageState: number,
  orderingState: ordering,
  sectionState: section,
  queryState: string,
) => {
  const [globalMods, setGlobalMods] = useState<GlobalModView[]>([])
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<Error | null>(null)

  useEffect(() => {
    const fetchGlobalMods = async () => {
      setIsLoading(true)
      try {
        const newGlobalMods = await window.getGlobalMods(
          orderingState,
          sectionState,
          queryState,
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
  }, [pageState, orderingState, sectionState, queryState]) // Include queryState in the dependency array

  return { globalMods, isLoading, error }
}

export { useGetGlobalMods }
