import type { FC } from "react";
import { useEffect } from "react";

import useGetProfiles from "./lib/getProfiles";
import { MainLayout } from "./layouts/MainLayout";
import Sidebar from "./components/sidebar-main";
import Header from "./components/header-main";
import { useBlockUI } from "./components/ui/block-ui";

const App: FC = () => {
  const { isBlocked, theme, unblock } = useBlockUI("black", true);
  const p = useGetProfiles();

  useEffect(() => {
    if (!p.isLoading || p.error) {
      unblock();
    } else if (p.error) {
      alert(p.error);
    }
  }, [p.isLoading, p.error, unblock]);

  return (
    <MainLayout
      sidebar={<Sidebar profiles={p.profiles} />}
      header={<Header />}
      blocking={{ isBlocked, theme }}
    >
      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3 text-white">
        <h1 className="font-semibold text-lg md:text-2xl text-white">Mods</h1>
      </div>
    </MainLayout>
  );
};

export default App;
