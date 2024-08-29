import { zodResolver } from "@hookform/resolvers/zod";
import { PlusCircle } from "lucide-react";
import * as React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "@components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@components/ui/dialog";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@components/ui/form";
import { Input } from "@components/ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@components/ui/select";

export const formSchema = z.object({
  key: z.string().min(3).max(255),
  group: z.enum(["all"]),
});

export type FormApiTokenValues = z.infer<typeof formSchema>;

interface FormApiTokenProps {
  title?: string;
  onSubmit: (values: FormApiTokenValues) => void;
}

export function FormApiToken(props: FormApiTokenProps) {
  const [isOpen, setIsOpen] = React.useState(false);
  const form = useForm<FormApiTokenValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      key: "",
      group: "all",
    },
  });

  React.useEffect(() => {
    if (isOpen) {
      form.reset();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isOpen]);

  function onSubmit(values: FormApiTokenValues) {
    form.reset();
    setIsOpen(false);
    props.onSubmit(values);
  }

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button>
          <PlusCircle className="mr-2 size-4" />
          <span>{props.title || "Create API Token"}</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-4">
            <DialogHeader>
              <DialogTitle>{props.title || "Create API Token"}</DialogTitle>
              <DialogDescription>Create a new API token to authenticate requests to the API.</DialogDescription>
            </DialogHeader>
            <FormField
              control={form.control}
              name="key"
              render={({ field }) => (
                <div className="grid gap-3">
                  <FormItem>
                    <FormLabel>Key</FormLabel>
                    <FormControl>
                      <Input placeholder="ex: super-duper-secret-key" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                </div>
              )}
            />
            <FormField
              control={form.control}
              name="group"
              render={({ field }) => (
                <div className="grid gap-3">
                  <FormItem>
                    <FormLabel>Group</FormLabel>
                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                      <FormControl>
                        <SelectTrigger aria-label="Select group">
                          <SelectValue placeholder="Select group" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value="all">All</SelectItem>
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                </div>
              )}
            />
            <DialogFooter>
              <Button type="submit">Save changes</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
