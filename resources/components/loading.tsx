import { Loader } from "lucide-react";

export default function Loading() {
  return (
    <div className="flex min-h-[100vh] w-full items-center justify-center">
      <div className="flex items-center">
        <Loader className="mr-2 size-4 animate-spin" />
        <span>Loading...</span>
      </div>
    </div>
  );
}
