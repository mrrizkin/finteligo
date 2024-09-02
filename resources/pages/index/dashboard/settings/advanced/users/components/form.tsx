import { UseFormReturn } from "react-hook-form";
import { z } from "zod";

import * as queries from "@hooks/queries";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "@components/ui/form";
import { Input } from "@components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@components/ui/select";

export const formSchema = z.object({
  name: z.string().min(3).max(255),
  email: z.string().email(),
  username: z.string().min(3).max(255),
  password: z.string().optional(),
  role_id: z.number(),
});

export type FormUserValues = z.infer<typeof formSchema>;

interface FormUserProps {
  form: UseFormReturn<FormUserValues>;
  children?: React.ReactNode;
}

export function FormUser({ form, children }: FormUserProps) {
  const { data: roles } = queries.useRoles({
    per_page: 9999,
  });

  return (
    <div className="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
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
                        <Input placeholder="ex: John Doe" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>Email</FormLabel>
                      <FormControl>
                        <Input type="email" placeholder="ex: john.doe@acme.com" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>Username</FormLabel>
                      <FormControl>
                        <Input type="username" placeholder="ex: john.doe" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl>
                        <Input type="password" {...field} />
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
      <div className="grid auto-rows-max items-start gap-4 lg:gap-8">
        <Card x-chunk="dashboard-07-chunk-3">
          <CardHeader>
            <CardTitle>Role</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid gap-6">
              <FormField
                control={form.control}
                name="role_id"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <Select
                        onValueChange={(val) => {
                          const valInt = parseInt(val);
                          if (isNaN(valInt)) {
                            field.onChange(0);
                            return;
                          }

                          field.onChange(valInt);
                        }}
                        defaultValue={String(field.value)}>
                        <FormControl>
                          <SelectTrigger aria-label="Select provider">
                            <SelectValue placeholder="Select provider" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          {(roles?.data.data || []).map((role) => (
                            <SelectItem key={role.id} value={String(role.id)}>
                              {role.name}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
            </div>
          </CardContent>
        </Card>
        {children}
      </div>
    </div>
  );
}
