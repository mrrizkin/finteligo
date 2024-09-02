import { Models } from "@schemas/models";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";

interface FormModelProps {
  data: Models;
  children?: React.ReactNode;
}

export function ShowModel({ data, children }: FormModelProps) {
  return (
    <div className="grid gap-4 md:grid-cols-[1fr_250px] lg:grid-cols-3 lg:gap-8">
      <div className="grid auto-rows-max items-start gap-4 lg:col-span-2 lg:gap-8">
        <Card x-chunk="dashboard-07-chunk-0">
          <CardHeader>
            <CardTitle>Model Detail</CardTitle>
            <CardDescription>The details of the model.</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid gap-6">
              <div className="grid gap-3">
                <Label>Model</Label>
                <Input placeholder="ex: GPT-3" value={data.model} disabled />
              </div>
              <div className="grid gap-3">
                <Label>Provider</Label>
                <Input placeholder="ex: OpenAI" value={data.provider} disabled />
              </div>

              <div className="grid gap-3">
                <Label>URL</Label>
                <Input placeholder="ex: https://api.openai.com" value={data.url} disabled />
              </div>

              <div className="grid gap-3">
                <Label>API Key</Label>
                <Input placeholder="ex: abc123" value={data.api_key} disabled />
              </div>
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
              <div className="grid gap-3">
                <Label>Enabled</Label>
                <Input placeholder="ex: true" value={data.enabled ? "Yes" : "No"} disabled />
              </div>
            </div>
          </CardContent>
        </Card>
        {children}
      </div>
    </div>
  );
}
