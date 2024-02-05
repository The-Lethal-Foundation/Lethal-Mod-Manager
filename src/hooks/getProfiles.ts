import { useState, useEffect } from 'react'

const useGetProfiles = () => {
  const [profiles, setProfiles] = useState<{ value: string; label: string }[]>(
    [],
  )
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState<Error | null>(null)

  const fetchProfiles = async () => {
    setIsLoading(true)
    try {
      const result = await window.getProfiles()
      setProfiles(
        result.map((profile: string) => ({ value: profile, label: profile })),
      )
    } catch (err) {
      setError(err as Error)
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    fetchProfiles()
  }, [])

  return { profiles, isLoading, error, fetchProfiles }
}

export default useGetProfiles
