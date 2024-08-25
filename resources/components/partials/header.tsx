import React from "react";

import { ThemeSelector } from "@components/theme-selector";

interface HeaderProps {
  children?: React.ReactNode;
  actions?: React.ReactNode;
}

export default function Header(props: HeaderProps) {
  return (
    <header className="sticky top-0 z-10 flex h-[53px] items-center gap-1 border-b bg-background px-4">
      {props.children}
      <div className="ml-auto gap-1.5 text-sm">
        {props.actions}
        <ThemeSelector />
      </div>
    </header>
  );
}
