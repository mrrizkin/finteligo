import * as React from "react";
import { useRoutes } from "react-router-dom";
import routes from "~react-pages";

import LoadingComponent from "@components/loading";

function App() {
  return <React.Suspense fallback={<LoadingComponent />}>{useRoutes(routes)}</React.Suspense>;
}

export { App };
