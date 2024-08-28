import {
  Book,
  Bot,
  Code2,
  LifeBuoy,
  LogOut,
  ReceiptText,
  Settings2,
  SquareTerminal,
  SquareUser,
  Triangle,
  User,
} from "lucide-react";
import { Link, NavLink, Outlet, useNavigate } from "react-router-dom";

import { toastValidation } from "@lib/utils";

import * as authService from "@services/auth";

import { Button } from "@components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@components/ui/dropdown-menu";
import { Tooltip, TooltipContent, TooltipTrigger } from "@components/ui/tooltip";

export default function DashboardPage() {
  const navigate = useNavigate();

  function logout() {
    toastValidation(authService.logout(), {
      success() {
        navigate("/auth/login");
      },
    });
  }

  return (
    <div className="grid h-screen w-full pl-[53px]">
      <aside className="inset-y fixed left-0 z-20 flex h-full flex-col border-r">
        <div className="border-b p-2">
          <Button variant="outline" size="icon" aria-label="Home" asChild>
            <Link to="/dashboard">
              <Triangle className="size-5 fill-foreground" />
            </Link>
          </Button>
        </div>
        <nav className="grid gap-1 p-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="nav-button-link rounded-lg"
                aria-label="Playground"
                asChild>
                <NavLink to="/dashboard/playground">
                  <SquareTerminal className="size-5" />
                </NavLink>
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Playground
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button variant="ghost" size="icon" className="nav-button-link rounded-lg" aria-label="Models" asChild>
                <NavLink to="/dashboard/models">
                  <Bot className="size-5" />
                </NavLink>
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Models
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button variant="ghost" size="icon" className="nav-button-link rounded-lg" aria-label="API" asChild>
                <NavLink to="/dashboard/api-docs">
                  <Code2 className="size-5" />
                </NavLink>
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              API
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button variant="ghost" size="icon" className="nav-button-link rounded-lg" aria-label="Documents" asChild>
                <NavLink to="/dashboard/documents">
                  <Book className="size-5" />
                </NavLink>
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Documents
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button variant="ghost" size="icon" className="nav-button-link rounded-lg" aria-label="Settings" asChild>
                <NavLink to="/dashboard/settings">
                  <Settings2 className="size-5" />
                </NavLink>
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Settings
            </TooltipContent>
          </Tooltip>
        </nav>
        <nav className="mt-auto grid gap-1 p-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="nav-button-link mt-auto rounded-lg"
                aria-label="Help"
                asChild>
                <NavLink to="/dashboard/help">
                  <LifeBuoy className="size-5" />
                </NavLink>
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Help
            </TooltipContent>
          </Tooltip>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon" className="nav-button-link mt-auto rounded-lg" aria-label="Account">
                <SquareUser className="size-5" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent className="w-56" side="right">
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuGroup>
                <DropdownMenuItem>
                  <User className="mr-2 size-4" />
                  Profile
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <ReceiptText className="mr-2 size-4" />
                  Billing
                </DropdownMenuItem>
                <DropdownMenuItem onSelect={() => navigate("/dashboard/settings")}>
                  <Settings2 className="mr-2 size-4" />
                  Settings
                </DropdownMenuItem>
              </DropdownMenuGroup>
              <DropdownMenuSeparator />
              <DropdownMenuItem onSelect={logout}>
                <LogOut className="mr-2 size-4" />
                Log out
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </nav>
      </aside>
      <Outlet />
    </div>
  );
}
