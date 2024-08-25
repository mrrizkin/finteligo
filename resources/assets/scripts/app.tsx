import * as React from "react";
import { useNavigate, useRoutes } from "react-router-dom";
import routes from "~react-pages";

import LoadingComponent from "@components/loading";

function App() {
  const navigate = useNavigate();

  React.useEffect(() => {
    let isLoggedIn = localStorage.getItem("isLoggedIn");
    if (!isLoggedIn) {
      localStorage.setItem("isLoggedIn", "false");
      isLoggedIn = "false";
    }

    if (isLoggedIn !== "true") {
      navigate("/auth/login");
    }
  }, [navigate]);

  return <React.Suspense fallback={<LoadingComponent />}>{useRoutes(routes)}</React.Suspense>;
}

export { App };
