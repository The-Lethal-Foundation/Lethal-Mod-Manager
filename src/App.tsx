import React, { useEffect } from 'react';
import './App.css';

import { Toaster } from "@/components/ui/sonner";
import { ProfileSelect } from './components/profile-select';
import { BlockUI, useBlockUI } from './components/ui/block-ui';
import useGetProfiles from './lib/getProfiles';
import { FileIcon, GlobeIcon, Package, Package2Icon } from 'lucide-react';
import { Button } from './components/ui/button';
import { Separator } from './components/ui/separator';
import { Input } from './components/ui/input';

const App: React.FC = () => {
  const { isBlocked, theme, block, unblock } = useBlockUI("black", true);
  const { profiles, isLoading: isLoadingProfiles, error } = useGetProfiles();

  useEffect(() => {
    if (!isLoadingProfiles || error) {
      unblock();
    }
  }, [isLoadingProfiles, error, unblock]);

  return (
    <main>
      <div key="1" className="grid min-h-screen w-full lg:grid-cols-[280px_1fr]">
        <div className="hidden border-r border-[#27272a] lg:block">
          <div className="flex h-full max-h-screen flex-col gap-2">
            <div className="flex h-[60px] justify-center border-b border-[#27272a] px-6">
              <div className="flex items-center gap-2 font-semibold">
                <Package2Icon color='white' className="h-6 w-6" />
                <span className="text-white">Lethal Mod Manager</span>
              </div>
            </div>
            <div className="flex-1 overflow-auto py-2">
              <nav className="grid justify-center px-4 text-sm font-medium gap-2">
                <ProfileSelect profiles={profiles} />
                <Separator className='my-2 bg-[#27272a]' />

                <Button variant="link" className='text-white'>
                  <FileIcon className='mr-2 h-4 w-4' />
                  Local mods
                </Button>

                <Button variant="link" className='text-white'>
                  <GlobeIcon className='mr-2 h-4 w-4' />
                  Online mods
                </Button>
              </nav>
            </div>
          </div>
        </div>
        <div className="flex flex-col">
          <header className="flex h-14 lg:h-[60px] items-center gap-4 border-b border-[#27272a] px-4">
            {/* <Link className="lg:hidden" href="#">
              <Package2Icon className="h-6 w-6" />
              <span className="sr-only">Home</span>
            </Link> */}
            <div className="w-full flex-1">
              <form>
                <div className="relative">
                  <Input type="search" placeholder="Search mods..." className="w-3/5 text-white border-[#27272a] focus:border-white" />
                </div>
              </form>
            </div>
          </header>
          <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">
            <div className="flex items-center">
              <h1 className="font-semibold text-lg md:text-2xl text-white">Mods</h1>
            </div>
            <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
              {/* <Card>
                <CardContent className="flex flex-col gap-4">
                  <img
                    alt="Mod Image"
                    className="aspect-square object-cover border border-gray-200 w-full rounded-lg overflow-hidden dark:border-gray-800"
                    height={200}
                    src="/placeholder.svg"
                    width={200}
                  />
                  <h3 className="font-semibold">Mod Name</h3>
                  <p className="text-sm text-gray-500 dark:text-gray-400">
                    This is a description of the mod. It includes details about what the mod does and how it changes the
                    game.
                  </p>
                </CardContent>
              </Card>
              <Card>
                <CardContent className="flex flex-col gap-4">
                  <img
                    alt="Mod Image"
                    className="aspect-square object-cover border border-gray-200 w-full rounded-lg overflow-hidden dark:border-gray-800"
                    height={200}
                    src="/placeholder.svg"
                    width={200}
                  />
                  <h3 className="font-semibold">Mod Name</h3>
                  <p className="text-sm text-gray-500 dark:text-gray-400">
                    This is a description of the mod. It includes details about what the mod does and how it changes the
                    game.
                  </p>
                </CardContent>
              </Card>
              <Card>
                <CardContent className="flex flex-col gap-4">
                  <img
                    alt="Mod Image"
                    className="aspect-square object-cover border border-gray-200 w-full rounded-lg overflow-hidden dark:border-gray-800"
                    height={200}
                    src="/placeholder.svg"
                    width={200}
                  />
                  <h3 className="font-semibold">Mod Name</h3>
                  <p className="text-sm text-gray-500 dark:text-gray-400">
                    This is a description of the mod. It includes details about what the mod does and how it changes the
                    game.
                  </p>
                </CardContent>
              </Card> */}
            </div>
          </main>
        </div>
      </div>
    </main>
  );
}

export default App;
