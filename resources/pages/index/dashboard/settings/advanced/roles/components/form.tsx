import * as React from "react";
import { UseFormReturn } from "react-hook-form";
import { z } from "zod";

import { cn } from "@lib/utils";

import * as queries from "@hooks/queries";

import { Button } from "@components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";
import { Checkbox } from "@components/ui/checkbox";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@components/ui/form";
import { Input } from "@components/ui/input";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@components/ui/tabs";

export const formSchema = z.object({
  name: z.string().min(3).max(255),
  slug: z.string().min(3).max(255),
  permission_ids: z.array(z.number()),
});

export type FormRoleValues = z.infer<typeof formSchema>;

interface FormRoleProps {
  form: UseFormReturn<FormRoleValues>;
  children?: React.ReactNode;
}

type ShowPermission = {
  [key: string]: {
    id: number;
    name: string;
  }[];
};

export function FormRole({ form, children }: FormRoleProps) {
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
  }, [response]);

  return (
    <div className="flex flex-col gap-4 lg:gap-8">
      <div
        className={cn(
          "grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8",
          !children && "md:grid-cols-1 lg:grid-cols-1",
        )}>
        <div className="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
          <Card x-chunk="dashboard-07-chunk-0">
            <CardHeader>
              <CardTitle>User Detail</CardTitle>
              <CardDescription>Input the details of the user.</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid gap-6">
                <FormField
                  control={form.control}
                  name="name"
                  render={({ field }) => (
                    <div className="grid gap-3">
                      <FormItem>
                        <FormLabel>Name</FormLabel>
                        <FormControl>
                          <Input placeholder="ex: Prompt Engineer" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    </div>
                  )}
                />
                <FormField
                  control={form.control}
                  name="slug"
                  render={({ field }) => (
                    <div className="grid gap-3">
                      <FormItem>
                        <FormLabel>Slug</FormLabel>
                        <FormControl>
                          <Input placeholder="ex: prompt_engineer" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    </div>
                  )}
                />
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
          <Tabs orientation="vertical">
            <TabsList className="relative z-[1] h-full flex-col items-start bg-transparent">
              {(Object.keys(permissions) as Array<keyof ShowPermission>).sort().map((group) => (
                <TabsTrigger value={String(group)} key={`group-${group}`} asChild>
                  <Button size="sm" variant="ghost" className="w-full justify-start py-6 pl-4 pr-10">
                    {group}
                  </Button>
                </TabsTrigger>
              ))}
            </TabsList>
            {/* @ts-expect-error default when there is none selected */}
            <TabsContent className="absolute right-0 top-0 z-[0] w-full pl-[250px]">
              <div className="grid grid-cols-1 gap-2 lg:grid-cols-2">
                <h1 className="text-xl font-semibold">Select permissions</h1>
              </div>
            </TabsContent>
            {(Object.keys(permissions) as Array<keyof ShowPermission>).sort().map((group) => (
              <TabsContent
                value={String(group)}
                key={`tab-group-${group}`}
                className="absolute right-0 top-0 z-[0] w-full pl-[250px]">
                <div className="grid grid-cols-1 gap-2 lg:grid-cols-2">
                  {permissions[group].map((permission) => (
                    <FormItem className="flex flex-row items-start space-x-3 space-y-0" key={permission.id}>
                      <FormControl>
                        <Checkbox
                          id={`permission-${permission.id}`}
                          onCheckedChange={(checked) => {
                            if (checked) {
                              form.setValue("permission_ids", [...form.getValues("permission_ids"), permission.id]);
                            } else {
                              form.setValue(
                                "permission_ids",
                                form.getValues("permission_ids").filter((id) => id !== permission.id),
                              );
                            }
                          }}
                          checked={form.watch("permission_ids").includes(permission.id)}
                        />
                      </FormControl>
                      <FormLabel>{permission.name}</FormLabel>
                      <FormMessage />
                    </FormItem>
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
