import { Link, NavLink, Outlet, useLocation } from "react-router-dom";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@components/ui/breadcrumb";

import Header from "@components/partials/header";

import { Show } from "@components/show";

export default function SettingsPage() {
  const location = useLocation();

  return (
    <div className="flex flex-col">
      <Header>
        <Breadcrumb className="hidden md:flex">
          <BreadcrumbList>
            <BreadcrumbItem>
              <BreadcrumbLink asChild>
                <Link to="/dashboard">Dashboard</Link>
              </BreadcrumbLink>
            </BreadcrumbItem>
            <BreadcrumbSeparator />
            <BreadcrumbItem>
              <BreadcrumbPage>Settings</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </Header>
      <div className="flex min-h-screen w-full flex-col">
        <main className="flex min-h-[calc(100vh_-_theme(spacing.16))] flex-1 flex-col gap-4 bg-muted/40 p-4 md:gap-8 md:p-10">
          <div className="mx-auto grid w-full gap-2">
            <h1 className="text-3xl font-semibold">Settings</h1>
          </div>
          <div className="mx-auto grid w-full items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]">
            <nav className="grid gap-4 text-sm text-muted-foreground" x-chunk="dashboard-04-chunk-0">
              <NavLink
                end
                to="/dashboard/settings"
                className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                General
              </NavLink>
              <NavLink
                to="/dashboard/settings/api-tokens"
                className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                API Tokens
              </NavLink>
              <NavLink
                to="/dashboard/settings/security"
                className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                Security
              </NavLink>
              <NavLink
                to="/dashboard/settings/integrations"
                className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                Integrations
              </NavLink>
              <NavLink
                to="/dashboard/settings/organizations"
                className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                Organizations
              </NavLink>
              <NavLink
                to="/dashboard/settings/advanced"
                className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                Advanced
              </NavLink>
              <Show when={location.pathname.includes("/dashboard/settings/advanced")}>
                <nav className="grid gap-4 pl-4 text-sm text-muted-foreground" x-chunk="dashboard-04-chunk-0">
                  <NavLink
                    to="/dashboard/settings/advanced/users"
                    className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                    Manage Users
                  </NavLink>
                  <NavLink
                    to="/dashboard/settings/advanced/roles"
                    className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                    Manage Roles
                  </NavLink>
                  <NavLink
                    to="/dashboard/settings/advanced/plans"
                    className={({ isActive }) => (isActive ? "font-semibold text-primary" : "")}>
                    Subscription Plan
                  </NavLink>
                </nav>
              </Show>
            </nav>
            <Outlet />
          </div>
        </main>
      </div>
    </div>
  );
}
