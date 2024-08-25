import { Toaster as Sonner } from "sonner";

import { useTheme } from "@components/theme-provider";

type ToasterProps = React.ComponentProps<typeof Sonner>;

const Toaster = ({ ...props }: ToasterProps) => {
  const { theme } = useTheme();

  return (
    <Sonner
      theme={theme as ToasterProps["theme"]}
      className="toaster group"
      toastOptions={{
        classNames: {
          toast:
            "group toast group-[.toaster]:bg-background group-[.toaster]:text-foreground group-[.toaster]:border-border group-[.toaster]:border group-[.toaster]:shadow-lg",
          description: "group-[.toast]:text-muted-foreground",
          actionButton: "group-[.toast]:bg-primary group-[.toast]:text-primary-foreground",
          cancelButton: "group-[.toast]:bg-muted group-[.toast]:text-muted-foreground",
          error:
            "group-[.toaster]:bg-red-100 group-[.toaster]:text-red-500 group-[.toaster]:border-red-200 group-[.toaster]:border",
          success:
            "group-[.toaster]:bg-green-100 group-[.toaster]:text-green-500 group-[.toaster]:border-green-200 group-[.toaster]:border",
          warning:
            "group-[.toaster]:bg-yellow-100 group-[.toaster]:text-yellow-500 group-[.toaster]:border-yellow-200 group-[.toaster]:border",
          info: "group-[.toaster]:bg-blue-100 group-[.toaster]:text-blue-500 group-[.toaster]:border-blue-200 group-[.toaster]:border",
        },
      }}
      {...props}
    />
  );
};

export { Toaster };
