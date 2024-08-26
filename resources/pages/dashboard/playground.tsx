import { Slider } from "@/components/ui/slider";
import llamaTokenizer from "llama-tokenizer-js";
import { Bird, CornerDownLeft, Rabbit, Settings, Turtle } from "lucide-react";
import * as React from "react";
import { Link } from "react-router-dom";
import { toast } from "sonner";

import { toastValidation } from "@lib/utils";

import { prompt } from "@services/playground";

import { Badge } from "@components/ui/badge";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@components/ui/breadcrumb";
import { Button } from "@components/ui/button";
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "@components/ui/drawer";
import { Label } from "@components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@components/ui/select";
import { Textarea } from "@components/ui/textarea";

import Header from "@components/partials/header";

export default function PlaygroundPage() {
  const [model, setModel] = React.useState<string | null>(null);
  const [message, setMessage] = React.useState<string | null>(null);
  const [temperature, setTemperature] = React.useState<number | null>(null);
  const [topP, setTopP] = React.useState<number | null>(null);
  const [topK, setTopK] = React.useState<number | null>(null);
  const [role, setRole] = React.useState<string | null>(null);
  const [content, setContent] = React.useState<string | null>(null);

  const [tokenLength, setTokenLength] = React.useState<number>(0);

  React.useEffect(() => {
    if (!message) {
      setTokenLength(0);
    }

    setTokenLength(llamaTokenizer.encode(message || "").length || 0);
  }, [message]);

  function onMessageChange(event: React.ChangeEvent<HTMLTextAreaElement>) {
    setMessage(event.target.value);
  }

  function onTemperatureChange(value: number[]) {
    setTemperature(Number((value[0] || 0) / 100));
  }

  function onTopPChange(value: number[]) {
    setTopP(Number((value[0] || 0) / 100));
  }

  function onTopKChange(value: number[]) {
    setTopK(Number((value[0] || 0) / 100));
  }

  function onContentChange(event: React.ChangeEvent<HTMLTextAreaElement>) {
    setContent(event.target.value);
  }

  function sendMessage() {
    console.log("Sending message...", model, message, temperature, topP, topK, role, content);
    if (!model || !message) {
      toast.error("Please select a model and enter a message.");
      return;
    }

    toastValidation(
      prompt({
        model,
        message,
        temperature,
        topP,
        topK,
        role,
        content,
      }),
      {
        success(data) {
          console.log(data);
        },
      },
    );
  }

  return (
    <div className="flex flex-col">
      <Header>
        <Breadcrumb className="hidden md:flex">
          <BreadcrumbList>
            <BreadcrumbItem>
              <BreadcrumbLink asChild>
                <Link to="/dashboard">Dashboard</Link>
              </BreadcrumbLink>
            </BreadcrumbItem>
            <BreadcrumbSeparator />
            <BreadcrumbItem>
              <BreadcrumbPage>Playground</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
        <Drawer>
          <DrawerTrigger asChild>
            <Button variant="ghost" size="icon" className="md:hidden">
              <Settings className="size-4" />
              <span className="sr-only">Settings</span>
            </Button>
          </DrawerTrigger>
          <DrawerContent className="max-h-[80vh]">
            <DrawerHeader>
              <DrawerTitle>Configuration</DrawerTitle>
              <DrawerDescription>Configure the settings for the model and messages.</DrawerDescription>
            </DrawerHeader>
            <form className="grid w-full items-start gap-6 overflow-auto p-4 pt-0">
              <fieldset className="grid gap-6 rounded-lg border p-4">
                <legend className="-ml-1 px-1 text-sm font-medium">Settings</legend>
                <div className="grid gap-3">
                  <Label htmlFor="model">Model</Label>
                  <Select onValueChange={setModel}>
                    <SelectTrigger id="model" className="items-start [&_[data-description]]:hidden">
                      <SelectValue placeholder="Select a model" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="genesis">
                        <div className="flex items-start gap-3 text-muted-foreground">
                          <Rabbit className="size-5" />
                          <div className="grid gap-0.5">
                            <p>
                              Neural <span className="font-medium text-foreground">Genesis</span>
                            </p>
                            <p className="text-xs" data-description>
                              Our fastest model for general use cases.
                            </p>
                          </div>
                        </div>
                      </SelectItem>
                      <SelectItem value="explorer">
                        <div className="flex items-start gap-3 text-muted-foreground">
                          <Bird className="size-5" />
                          <div className="grid gap-0.5">
                            <p>
                              Neural <span className="font-medium text-foreground">Explorer</span>
                            </p>
                            <p className="text-xs" data-description>
                              Performance and speed for efficiency.
                            </p>
                          </div>
                        </div>
                      </SelectItem>
                      <SelectItem value="quantum">
                        <div className="flex items-start gap-3 text-muted-foreground">
                          <Turtle className="size-5" />
                          <div className="grid gap-0.5">
                            <p>
                              Neural <span className="font-medium text-foreground">Quantum</span>
                            </p>
                            <p className="text-xs" data-description>
                              The most powerful model for complex computations.
                            </p>
                          </div>
                        </div>
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="temperature">Temperature</Label>
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">Factual</span>
                    <span className="text-xs text-muted-foreground">Creative</span>
                  </div>
                  <Slider defaultValue={[40]} max={100} step={5} onValueChange={onTemperatureChange} />
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="top-p">Top P</Label>
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">0.0</span>
                    <span className="text-xs text-muted-foreground">1.0</span>
                  </div>
                  <Slider defaultValue={[40]} max={100} step={10} onValueChange={onTopPChange} />
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="top-k">Top K</Label>
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">0.0</span>
                    <span className="text-xs text-muted-foreground">1.0</span>
                  </div>
                  <Slider defaultValue={[40]} max={100} step={10} onValueChange={onTopKChange} />
                </div>
              </fieldset>
              <fieldset className="grid gap-6 rounded-lg border p-4">
                <legend className="-ml-1 px-1 text-sm font-medium">Messages</legend>
                <div className="grid gap-3">
                  <Label htmlFor="role">Role</Label>
                  <Select defaultValue="system" onValueChange={setRole}>
                    <SelectTrigger>
                      <SelectValue placeholder="Select a role" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="system">System</SelectItem>
                      <SelectItem value="user">User</SelectItem>
                      <SelectItem value="assistant">Assistant</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="content">Content</Label>
                  <Textarea id="content" placeholder="You are a..." onChange={onContentChange} />
                </div>
              </fieldset>
            </form>
          </DrawerContent>
        </Drawer>
      </Header>
      <main className="grid flex-1 gap-4 overflow-auto p-4 md:grid-cols-2 lg:grid-cols-3">
        <div className="relative hidden flex-col items-start gap-8 md:flex" x-chunk="dashboard-03-chunk-0">
          <form className="grid w-full items-start gap-6">
            <fieldset className="grid gap-6 rounded-lg border p-4">
              <legend className="-ml-1 px-1 text-sm font-medium">Settings</legend>
              <div className="grid gap-3">
                <Label htmlFor="model">Model</Label>
                <Select onValueChange={setModel}>
                  <SelectTrigger id="model" className="items-start [&_[data-description]]:hidden">
                    <SelectValue placeholder="Select a model" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="genesis">
                      <div className="flex items-start gap-3 text-muted-foreground">
                        <Rabbit className="size-5" />
                        <div className="grid gap-0.5">
                          <p>
                            Neural <span className="font-medium text-foreground">Genesis</span>
                          </p>
                          <p className="text-xs" data-description>
                            Our fastest model for general use cases.
                          </p>
                        </div>
                      </div>
                    </SelectItem>
                    <SelectItem value="explorer">
                      <div className="flex items-start gap-3 text-muted-foreground">
                        <Bird className="size-5" />
                        <div className="grid gap-0.5">
                          <p>
                            Neural <span className="font-medium text-foreground">Explorer</span>
                          </p>
                          <p className="text-xs" data-description>
                            Performance and speed for efficiency.
                          </p>
                        </div>
                      </div>
                    </SelectItem>
                    <SelectItem value="quantum">
                      <div className="flex items-start gap-3 text-muted-foreground">
                        <Turtle className="size-5" />
                        <div className="grid gap-0.5">
                          <p>
                            Neural <span className="font-medium text-foreground">Quantum</span>
                          </p>
                          <p className="text-xs" data-description>
                            The most powerful model for complex computations.
                          </p>
                        </div>
                      </div>
                    </SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="grid gap-3">
                <Label htmlFor="temperature">Temperature</Label>
                <div className="flex items-center justify-between">
                  <span className="text-xs text-muted-foreground">Factual</span>
                  <span className="text-xs text-muted-foreground">Creative</span>
                </div>
                <Slider defaultValue={[40]} max={100} step={10} onValueChange={onTemperatureChange} />
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div className="grid gap-3">
                  <Label htmlFor="top-p">Top P</Label>
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">0.0</span>
                    <span className="text-xs text-muted-foreground">1.0</span>
                  </div>
                  <Slider defaultValue={[40]} max={100} step={10} onValueChange={onTopPChange} />
                </div>
                <div className="grid gap-3">
                  <Label htmlFor="top-k">Top K</Label>
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-muted-foreground">0.0</span>
                    <span className="text-xs text-muted-foreground">1.0</span>
                  </div>
                  <Slider defaultValue={[40]} max={100} step={10} onValueChange={onTopKChange} />
                </div>
              </div>
            </fieldset>
            <fieldset className="grid gap-6 rounded-lg border p-4">
              <legend className="-ml-1 px-1 text-sm font-medium">Messages</legend>
              <div className="grid gap-3">
                <Label htmlFor="role">Role</Label>
                <Select defaultValue="system" onValueChange={setRole}>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a role" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="system">System</SelectItem>
                    <SelectItem value="user">User</SelectItem>
                    <SelectItem value="assistant">Assistant</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="grid gap-3">
                <Label htmlFor="content">Content</Label>
                <Textarea
                  id="content"
                  placeholder="You are a..."
                  className="min-h-[9.5rem]"
                  onChange={onContentChange}
                />
              </div>
            </fieldset>
          </form>
        </div>
        <div className="relative flex h-full min-h-[50vh] flex-col rounded-xl bg-muted/50 p-4 lg:col-span-2">
          <Badge variant="outline" className="absolute right-3 top-3">
            Output
          </Badge>
          <div className="flex-1" />
          <form
            className="relative overflow-hidden rounded-lg border bg-background focus-within:ring-1 focus-within:ring-ring"
            x-chunk="dashboard-03-chunk-1">
            <Label htmlFor="message" className="sr-only">
              Message
            </Label>
            <Textarea
              id="message"
              placeholder="Type your message here..."
              className="min-h-12 resize-none border-0 p-3 shadow-none focus-visible:ring-0"
              onChange={onMessageChange}
            />
            <div className="flex items-center p-3 pt-0">
              <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
                <span>{tokenLength}</span>
                <span>tokens</span>
              </div>
              {/* <Tooltip>
                <TooltipTrigger asChild>
                  <Button variant="ghost" size="icon">
                    <Paperclip className="size-4" />
                    <span className="sr-only">Attach file</span>
                  </Button>
                </TooltipTrigger>
                <TooltipContent side="top">Attach File</TooltipContent>
              </Tooltip>
              <Tooltip>
                <TooltipTrigger asChild>
                  <Button variant="ghost" size="icon">
                    <Mic className="size-4" />
                    <span className="sr-only">Use Microphone</span>
                  </Button>
                </TooltipTrigger>
                <TooltipContent side="top">Use Microphone</TooltipContent>
              </Tooltip>
              */}
              <Button type="submit" size="sm" className="ml-auto gap-1.5" onClick={sendMessage}>
                Send Message
                <CornerDownLeft className="size-3.5" />
              </Button>
            </div>
          </form>
        </div>
      </main>
    </div>
  );
}
