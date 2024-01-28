import type { FC } from "react";
import { Input } from "./ui/input";

const Header: FC = () => {
  return (
    <header className="bg-[#09090b] flex h-14 lg:h-[60px] items-center gap-4 border-b border-[#27272a] px-4">
      <div className="w-full flex-1">
        <form>
          <div className="relative">
            <Input
              type="search"
              placeholder="Search mods..."
              className="w-3/5 text-white border-[#27272a] focus:border-white"
            />
          </div>
        </form>
      </div>
    </header>
  );
};

export default Header;
