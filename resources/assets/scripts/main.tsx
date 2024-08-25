import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HashRouter } from "react-router-dom";

import "@styles/index.css";

import { App } from "./app.tsx";
import { Provider } from "./provider.tsx";

const root = document.getElementById("root");

const error = console.error;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
console.error = (...args: any) => {
  if (
    /Warning: %s: Support for defaultProps will be removed from function components in a future major release. Use JavaScript default parameters instead.%s/.test(
      args[0],
    )
  )
    return;
  error(...args);
};

if (!root) {
  throw new Error("Root element not found");
}

const app = createRoot(root);
app.render(
  <StrictMode>
    <Provider>
      <HashRouter>
        <App />
      </HashRouter>
    </Provider>
  </StrictMode>,
);
