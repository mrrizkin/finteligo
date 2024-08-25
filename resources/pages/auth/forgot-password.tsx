import { Link } from "react-router-dom";

import { Button } from "@components/ui/button";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";

export default function ForgotPasswordPage() {
  return (
    <div className="mx-auto grid w-[350px] gap-6">
      <div className="grid gap-2 text-center">
        <h1 className="text-3xl font-bold">Forgot Password</h1>
        <p className="text-balance text-muted-foreground">Enter your email below to reset your password</p>
      </div>
      <div className="grid gap-4">
        <div className="grid gap-2">
          <Label htmlFor="email">Email</Label>
          <Input id="email" type="email" placeholder="claire49@acme.com" required />
        </div>
        <Button type="submit" className="w-full">
          Reset Password
        </Button>
      </div>
      <div className="mt-4 text-center text-sm">
        Remember your password?{" "}
        <Link to="/auth/login" className="underline">
          Login
        </Link>
      </div>
    </div>
  );
}
