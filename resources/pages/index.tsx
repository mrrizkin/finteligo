import { useQuery } from "react-query";
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { toast } from "sonner";

import * as authService from "@services/auth";

import { Error } from "@components/partials/utils";

import Loading from "@components/loading";

export default function IndexPage() {
  const location = useLocation();
  const { isLoading, isError, error } = useQuery("identity", () => authService.identity());

  if (isLoading) {
    return <Loading />;
  }

  if (isError) {
    if (error instanceof Response && error.status === 401) {
      toast.error("You need to login to access this page");
      return <Navigate to="/auth/login" />;
    }
    return <Error response={error as Response} />;
  }

  if (location.pathname === "/") {
    return <Navigate to="/dashboard" />;
  }

  return <Outlet />;
}
