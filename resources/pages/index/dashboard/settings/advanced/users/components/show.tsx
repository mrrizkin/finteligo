import { User } from "@schemas/user";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";

interface FormModelProps {
  data: User;
  children?: React.ReactNode;
}

export function ShowUser({ data, children }: FormModelProps) {
  return (
    <div className="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
      <div className="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
        <Card x-chunk="dashboard-07-chunk-0">
          <CardHeader>
            <CardTitle>User Detail</CardTitle>
            <CardDescription>The details of the user.</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-6">
              <div className="grid gap-3">
                <Label>Name</Label>
                <Input placeholder="ex: John Doe" value={data.name} disabled />
              </div>
              <div className="grid gap-3">
                <Label>Email</Label>
                <Input type="email" placeholder="ex: john.doe@acme.com" value={data.email} disabled />
              </div>
              <div className="grid gap-3">
                <Label>Username</Label>
                <Input type="username" placeholder="ex: john.doe" value={data.username} disabled />
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
      <div className="grid auto-rows-max items-start gap-4 lg:gap-8">
        <Card x-chunk="dashboard-07-chunk-3">
          <CardHeader>
            <CardTitle>Role</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid gap-6">
              <div className="grid gap-3">
                <Input placeholder="ex: true" value={data.role?.name} disabled />
              </div>
            </div>
          </CardContent>
        </Card>
        {children}
      </div>
    </div>
  );
}
