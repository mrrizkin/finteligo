import react from "@vitejs/plugin-react-swc";
import path from "path";
import { defineConfig } from "vite";
import backendPlugin from "vite-plugin-backend";
import fullReload from "vite-plugin-full-reload";
import pageRouter from "vite-plugin-pages";

export default defineConfig({
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./resources"),
      "@styles": path.resolve(__dirname, "./resources/assets/styles"),
      "@scripts": path.resolve(__dirname, "./resources/assets/scripts"),
      "@lib": path.resolve(__dirname, "./resources/assets/scripts/lib"),
      "@schemas": path.resolve(__dirname, "./resources/assets/scripts/schemas"),
      "@services": path.resolve(__dirname, "./resources/assets/scripts/services"),
      "@components": path.resolve(__dirname, "./resources/components"),
      "@pages": path.resolve(__dirname, "./resources/pages"),
    },
  },
  plugins: [
    backendPlugin({
      input: ["resources/assets/scripts/main.tsx"],
    }),
    fullReload(["tmp/main"]),
    react(),
    pageRouter({
      pagesDir: "resources/pages",
    }),
  ],
});
