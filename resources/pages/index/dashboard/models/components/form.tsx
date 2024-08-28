import { UseFormReturn } from "react-hook-form";
import { z } from "zod";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";
import { FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@components/ui/form";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@components/ui/select";

export const formSchema = z.object({
  model: z.string().min(3).max(255),
  provider: z.enum(["ollama", "openai", "anthropic"]),
  url: z.string().optional(),
  api_key: z.string().optional(),
  enabled: z.boolean().optional(),
  token: z.string().optional(),
});

export type FormModelValues = z.infer<typeof formSchema>;

interface FormModelProps {
  form: UseFormReturn<FormModelValues>;
  children?: React.ReactNode;
}

export function FormModel({ form, children }: FormModelProps) {
  return (
    <div className="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
      <div className="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
        <Card x-chunk="dashboard-07-chunk-0">
          <CardHeader>
            <CardTitle>Model Detail</CardTitle>
            <CardDescription>Input the details of the model.</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-6">
              <FormField
                control={form.control}
                name="model"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>Model</FormLabel>
                      <FormControl>
                        <Input placeholder="ex: GPT-3" {...field} />
                      </FormControl>
                      <FormDescription>Select the model you want to use.</FormDescription>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="provider"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>Provider</FormLabel>
                      <Select onValueChange={field.onChange} defaultValue={field.value}>
                        <FormControl>
                          <SelectTrigger aria-label="Select provider">
                            <SelectValue placeholder="Select provider" />
                          </SelectTrigger>
                        </FormControl>
                        <SelectContent>
                          <SelectItem value="ollama">Ollama</SelectItem>
                          <SelectItem value="openai">OpenAI</SelectItem>
                          <SelectItem value="anthropic">Anthropic</SelectItem>
                        </SelectContent>
                      </Select>
                      <FormDescription>Select the provider of the model.</FormDescription>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="url"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>URL</FormLabel>
                      <FormControl>
                        <Input placeholder="ex: https://api.openai.com" {...field} />
                      </FormControl>
                      <FormDescription>Enter the URL of the model.</FormDescription>
                      <FormMessage />
                    </FormItem>
                  </div>
                )}
              />
              <FormField
                control={form.control}
                name="api_key"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <FormItem>
                      <FormLabel>API Key</FormLabel>
                      <FormControl>
                        <Input placeholder="ex: abc123" {...field} />
                      </FormControl>
                      <FormDescription>Enter the API key of the model.</FormDescription>
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
            <CardTitle>Enable Model</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid gap-6">
              <FormField
                control={form.control}
                name="enabled"
                render={({ field }) => (
                  <div className="grid gap-3">
                    <Label htmlFor="enabled">Enabled</Label>
                    <Select
                      onValueChange={(value) => {
                        if (value === "yes") {
                          field.onChange(true);
                        }
                        if (value === "no") {
                          field.onChange(false);
                        }
                      }}
                      defaultValue={field.value ? "yes" : "no"}>
                      <SelectTrigger id="enabled" aria-label="Select enabled">
                        <SelectValue placeholder="Select enabled" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="yes">Yes</SelectItem>
                        <SelectItem value="no">No</SelectItem>
                      </SelectContent>
                    </Select>
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
