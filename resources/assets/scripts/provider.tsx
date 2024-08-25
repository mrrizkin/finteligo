import { QueryClient, QueryClientProvider } from "react-query";

import { Toaster } from "@components/ui/sonner";
import { TooltipProvider } from "@components/ui/tooltip";

import { ThemeProvider } from "@components/theme-provider";

const queryClient = new QueryClient();

interface ProviderProps {
  children: React.ReactNode;
}

export function Provider(props: ProviderProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider defaultTheme="light" storageKey="finteligo-ui-theme">
        <TooltipProvider>{props.children}</TooltipProvider>
        <Toaster richColors />
      </ThemeProvider>
    </QueryClientProvider>
  );
}
