import react from "@vitejs/plugin-react-swc";
import path from "path";
import { defineConfig } from "vite";
import backendPlugin from "vite-plugin-backend";
import fullReload from "vite-plugin-full-reload";
import pageRouter from "vite-plugin-pages";

export default defineConfig({
  resolve: {
    alias: {
      "@styles": path.resolve(__dirname, "./resources/assets/styles"),
      "@lib": path.resolve(__dirname, "./resources/assets/scripts/lib"),
      "@schemas": path.resolve(__dirname, "./resources/assets/scripts/schemas"),
      "@hooks": path.resolve(__dirname, "./resources/assets/scripts/hooks"),
      "@services": path.resolve(__dirname, "./resources/assets/scripts/services"),
      "@scripts": path.resolve(__dirname, "./resources/assets/scripts"),
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
      exclude: ["**/components/**/*"],
    }),
  ],
});
