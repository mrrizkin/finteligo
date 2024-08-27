import { Loader } from "lucide-react";
import React from "react";

export function Loading() {
  return (
    <div className="flex h-full items-center justify-center gap-2">
      <Loader className="size-4 animate-spin" />
      <span>Loading...</span>
    </div>
  );
}

export function Error({ response }: { response: Response }) {
  const [error, setError] = React.useState<string | null>(null);

  React.useEffect(() => {
    response.json().then((data) => {
      setError(data.message);
    });
  }, [response]);

  return (
    <div className="flex h-full flex-col items-center justify-center gap-2">
      <div>Error: {error || "An error occurred"}</div>
    </div>
  );
}
