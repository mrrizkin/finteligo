import * as React from "react";

import { Role } from "@schemas/role";

import * as queries from "@hooks/queries";

import { Button } from "@components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";
import { Checkbox } from "@components/ui/checkbox";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@components/ui/tabs";

interface FormModelProps {
  data: Role;
  children?: React.ReactNode;
}

type ShowPermission = {
  [key: string]: {
    id: number;
    name: string;
  }[];
};

export function ShowRole({ data, children }: FormModelProps) {
  const { data: response } = queries.usePermissions();
  const [permissions, setPermissions] = React.useState<ShowPermission>({});

  React.useEffect(() => {
    if (response?.data.data) {
      const perms: ShowPermission = {};
      for (let i = 0; i < response.data.data.length; i++) {
        const group = response.data.data[i].group;
        const id = response.data.data[i].id;
        const name = response.data.data[i].name;
        if (group && id && name) {
          if (!perms[group]) {
            perms[group] = [];
          }

          perms[group].push({ id, name });
        }
      }

      setPermissions(perms);
    }
  }, [response, data]);

  return (
    <div className="flex flex-col gap-4 lg:gap-8">
      <div className="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
        <div className="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
          <Card x-chunk="dashboard-07-chunk-0">
            <CardHeader>
              <CardTitle>Role Detail</CardTitle>
              <CardDescription>The details of the role.</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid gap-6">
                <div className="grid gap-3">
                  <Label>Name</Label>
                  <Input placeholder="ex: John Doe" value={data.name} disabled />
                </div>
                <div className="grid gap-3">
                  <Label>slug</Label>
                  <Input type="email" placeholder="ex: john.doe@acme.com" value={data.slug} disabled />
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
        <div className="grid auto-rows-max items-start gap-4 lg:gap-8">{children}</div>
      </div>

      <Card className="w-full">
        <CardHeader>
          <CardTitle>Permission Detail</CardTitle>
          <CardDescription>All permissions of the role.</CardDescription>
        </CardHeader>
        <CardContent className="relative">
          <Tabs defaultValue="Role" orientation="vertical">
            <TabsList className="relative z-[1] h-full flex-col items-start bg-transparent">
              {(Object.keys(permissions) as Array<keyof ShowPermission>).sort().map((group) => (
                <TabsTrigger value={String(group)} key={`group-${group}`} asChild>
                  <Button size="sm" variant="ghost" className="w-full justify-start py-6 pl-4 pr-10">
                    {group}
                  </Button>
                </TabsTrigger>
              ))}
            </TabsList>
            {(Object.keys(permissions) as Array<keyof ShowPermission>).sort().map((group) => (
              <TabsContent
                value={String(group)}
                key={`tab-group-${group}`}
                className="absolute right-0 top-0 z-[0] w-full pl-[250px]">
                <div className="grid grid-cols-1 gap-2 lg:grid-cols-2">
                  {permissions[group].map((permission) => (
                    <div className="flex flex-row items-start space-x-3 space-y-0" key={`permission-${permission.id}`}>
                      <Checkbox
                        id={`permission-${permission.id}`}
                        checked={data.role_permissions?.find((perm) => perm.id === permission.id) ? true : false}
                      />
                      <Label id={`permission-${permission.id}`}>{permission.name}</Label>
                    </div>
                  ))}
                </div>
              </TabsContent>
            ))}
          </Tabs>
        </CardContent>
      </Card>
    </div>
  );
}
