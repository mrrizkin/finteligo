import { CaretSortIcon } from "@radix-ui/react-icons";
import llamaTokenizer from "llama-tokenizer-js";
import { CheckIcon, CornerDownLeft, Settings } from "lucide-react";
import * as React from "react";
import { Link } from "react-router-dom";
import { toast } from "sonner";

import { cn } from "@lib/utils";

import { Models } from "@schemas/models";

import * as playgroundService from "@services/playground";

import * as queries from "@hooks/queries";

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
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@components/ui/command";
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "@components/ui/drawer";
import { Label } from "@components/ui/label";
import { Popover, PopoverContent, PopoverTrigger } from "@components/ui/popover";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@components/ui/select";
import { Slider } from "@components/ui/slider";
import { Textarea } from "@components/ui/textarea";

import Header from "@components/partials/header";

type ChatMessage = {
  role: string;
  content: string[];
};

export default function PlaygroundPage() {
  const [chatHistory, setChatHistory] = React.useState<ChatMessage[]>([]);
  const { data: response } = queries.useModels();

  const [model, setModel] = React.useState<Models | null>(null);
  const [message, setMessage] = React.useState<string | null>(null);
  const [assistantMessage, setAssistantMessage] = React.useState<string[]>([]);
  const [temperature, setTemperature] = React.useState<number | null>(null);
  const [topP, setTopP] = React.useState<number | null>(null);
  const [promptStatus, setPromptStatus] = React.useState<string | null>(null);
  const [topK, setTopK] = React.useState<number | null>(null);
  const [role, setRole] = React.useState<string | null>(null);
  const [content, setContent] = React.useState<string | null>(null);

  const [modelOpen, setModelOpen] = React.useState(false);

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
    if (!model || !message?.trim()) {
      toast.error("Please select a model and enter a message.");
      return;
    }

    const currentChatHistory = [...chatHistory];

    setChatHistory((prevChatHistory) => [
      ...prevChatHistory,
      {
        role: "user",
        content: [message],
      },
    ]);

    setMessage("");
    setPromptStatus("pending");

    playgroundService.prompt({
      payload: {
        model: model.model || "",
        chat_history: currentChatHistory || [],
        message,
        temperature,
        topP,
        topK,
        role,
        content,
        token: model.token,
        stream: true,
      },
      stream(value) {
        setPromptStatus("streaming");
        const message = value.data.slice(1, -1);

        if (message.includes("\\n")) {
          const messages = message.split("\\n");
          setAssistantMessage((prev) => {
            if (prev.length === 0) {
              return messages;
            }

            return [...prev.slice(0, -1), prev[prev.length - 1] + messages[0], ...messages.slice(1)];
          });

          return;
        }

        setAssistantMessage((prev) => {
          if (prev.length === 0) {
            return [message];
          }

          return [...prev.slice(0, -1), prev[prev.length - 1] + message];
        });
      },
      done() {
        setPromptStatus("done");
        setAssistantMessage((finalAssistantMessage) => {
          setChatHistory((prevChatHistory) => [
            ...prevChatHistory,
            {
              role: "assistant",
              content: finalAssistantMessage || "No answer provided.",
            },
          ]);
          return [];
        });
      },
      error(error) {
        setPromptStatus("error");
        toast.error(error);
      },
    });
  }

  function handleCtrlEnterKey(event: React.KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key === "Enter" && event.ctrlKey) {
      event.preventDefault();
      sendMessage();
    }
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
                  <Popover open={modelOpen} onOpenChange={setModelOpen}>
                    <PopoverTrigger asChild>
                      <Button variant="outline" role="combobox" className="justify-between">
                        {model ? (
                          <div className="flex items-start gap-3 text-muted-foreground">
                            <div className="grid gap-0.5">
                              <p className="text-foreground">
                                {model.model}
                                <span className="text-muted-foreground">({model.provider})</span>
                              </p>
                            </div>
                          </div>
                        ) : (
                          "Select model..."
                        )}
                        <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-[200px] p-0">
                      <Command>
                        <CommandInput placeholder="Search model..." className="h-9" />
                        <CommandList>
                          <CommandEmpty>No model found.</CommandEmpty>
                          <CommandGroup>
                            {(response?.data.data || []).map((m) => (
                              <CommandItem
                                key={m.token}
                                value={m.token}
                                onSelect={() => {
                                  setModel(m);
                                  setModelOpen(false);
                                }}>
                                <div className="flex items-start gap-3 text-muted-foreground">
                                  <div className="grid gap-0.5">
                                    <p>{m.model}</p>
                                    <p className="text-xs" data-description>
                                      {m.provider}
                                    </p>
                                  </div>
                                </div>
                                <CheckIcon className={cn("ml-auto h-4 w-4", model ? "opacity-100" : "opacity-0")} />
                              </CommandItem>
                            ))}
                          </CommandGroup>
                        </CommandList>
                      </Command>
                    </PopoverContent>
                  </Popover>
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
                <Popover open={modelOpen} onOpenChange={setModelOpen}>
                  <PopoverTrigger asChild>
                    <Button variant="outline" role="combobox" className="justify-between">
                      {model ? (
                        <div className="flex items-start gap-3 text-muted-foreground">
                          <div className="grid gap-0.5">
                            <p className="text-foreground">
                              {model.model} <span className="text-muted-foreground">({model.provider})</span>
                            </p>
                          </div>
                        </div>
                      ) : (
                        "Select model..."
                      )}
                      <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent className="w-[200px] p-0">
                    <Command>
                      <CommandInput placeholder="Search model..." className="h-9" />
                      <CommandList>
                        <CommandEmpty>No model found.</CommandEmpty>
                        <CommandGroup>
                          {(response?.data.data || []).map((m) => (
                            <CommandItem
                              key={m.token}
                              value={m.token}
                              onSelect={() => {
                                setModel(m);
                                setModelOpen(false);
                              }}>
                              <div className="flex items-start gap-3 text-muted-foreground">
                                <div className="grid gap-0.5">
                                  <p>{m.model}</p>
                                  <p className="text-xs" data-description>
                                    {m.provider}
                                  </p>
                                </div>
                              </div>
                              <CheckIcon className={cn("ml-auto h-4 w-4", model ? "opacity-100" : "opacity-0")} />
                            </CommandItem>
                          ))}
                        </CommandGroup>
                      </CommandList>
                    </Command>
                  </PopoverContent>
                </Popover>
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
          <div className="flex-1 py-8">
            <div className="flex flex-col gap-2">
              {chatHistory.map((chat, index) =>
                chat.role === "user" ? (
                  <div key={index} className="flex items-start justify-end gap-2">
                    <div className="max-w-[80%] space-y-4">
                      {chat.content.map((content, i) => (
                        <p key={i} className="text-sm">
                          {content}
                        </p>
                      ))}
                    </div>
                    <Badge variant="secondary">User</Badge>
                  </div>
                ) : (
                  <div key={index} className="flex items-start justify-start gap-2">
                    <Badge variant="default">Assistant</Badge>
                    <div className="max-w-[80%] space-y-4">
                      {chat.content.map((content, i) => (
                        <p key={i} className="text-sm">
                          {content}
                        </p>
                      ))}
                    </div>
                  </div>
                ),
              )}
              {assistantMessage.length > 0 && (
                <div className="flex items-start justify-start gap-2">
                  <Badge variant="default">Assistant</Badge>
                  <div className="max-w-[80%] space-y-4">
                    {assistantMessage.map((content, i) => (
                      <p key={i} className="text-sm">
                        {content}
                      </p>
                    ))}
                  </div>
                </div>
              )}
            </div>
          </div>
          <form
            className={cn(
              "relative overflow-hidden rounded-lg border focus-within:ring-1 focus-within:ring-ring",
              promptStatus === "streaming" || promptStatus === "pending" ? "bg-muted" : "bg-background",
            )}
            x-chunk="dashboard-03-chunk-1">
            <Label htmlFor="message" className="sr-only">
              Message
            </Label>
            <Textarea
              disabled={promptStatus === "pending" || promptStatus === "streaming"}
              onKeyUp={handleCtrlEnterKey}
              id="message"
              placeholder="Type your message here..."
              className="min-h-12 resize-none border-0 p-3 shadow-none focus-visible:ring-0"
              onChange={onMessageChange}
              value={message || ""}
            />
            <div className="flex items-center p-3 pt-0">
              <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
                <span>{tokenLength}</span>
                <span>tokens</span>
              </div>
              <Button
                type="submit"
                size="sm"
                className="ml-auto gap-1.5"
                onClick={sendMessage}
                disabled={!model || !message || promptStatus === "pending" || promptStatus === "streaming"}>
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
