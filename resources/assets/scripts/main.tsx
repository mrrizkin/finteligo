import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HashRouter } from "react-router-dom";
import { registerSW } from "virtual:pwa-register";

import "@styles/index.css";

import { App } from "./app.tsx";
import { Provider } from "./provider.tsx";

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

const intervalMS = 60 * 60 * 1000;
registerSW({
  onRegistered(r) {
    if (r) {
      setInterval(() => {
        r.update();
      }, intervalMS);
    }
  },
});

const root = document.getElementById("root");

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
