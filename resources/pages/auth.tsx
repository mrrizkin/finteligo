import { Outlet } from "react-router-dom";

export default function AuthPageLayout() {
  return (
    <div className="min-h-[100vh] w-full lg:grid lg:grid-cols-2">
      <div className="flex items-center justify-center py-12">
        <Outlet />
      </div>
      <div className="hidden bg-muted lg:block">
        <img
          src="/financial-advisor-robot.png"
          alt="Image"
          width="1920"
          height="1080"
          className="h-full max-h-[100vh] w-full object-cover dark:brightness-[0.2] dark:grayscale"
        />
      </div>
    </div>
  );
}
