import { useState, useEffect } from 'react';

export interface Mod {
  mod_name: string;
  mod_author: string;
  mod_version: string;
  mod_description: string;
  mod_path_name: string;
}

const useGetMods = (profileName: string) => {
  const [mods, setMods] = useState<Mod[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    const fetchMods = async () => {
      setIsLoading(true);
      try {
        const modsData = await window.getMods(profileName);
        console.log('got moooooods', modsData);
        setMods(modsData);
      } catch (err) {
        setError(err as Error);
      } finally {
        setIsLoading(false);
      }
    };

    if (profileName) {
      fetchMods();
    }
  }, [profileName]);

  return { mods, isLoading, error };
};

export default useGetMods;
