import { zodResolver } from "@hookform/resolvers/zod";
import { ChevronLeft } from "lucide-react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";

import * as modelsService from "@services/models";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@components/ui/breadcrumb";
import { Button } from "@components/ui/button";
import { Form } from "@components/ui/form";

import Header from "@components/partials/header";

import { FormModel, FormModelValues, formSchema } from "./components/form";

export default function CreateModelPage() {
  const navigate = useNavigate();
  const form = useForm<FormModelValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      model: "",
      provider: "ollama",
      url: "",
      api_key: "",
      enabled: false,
      token: "",
    },
  });

  function onSubmit(values: FormModelValues) {
    modelsService.create(values).then(() => {
      form.reset();
      navigate("/dashboard/models");
    });
  }

  function onDiscard() {
    form.reset();
    navigate("/dashboard/models");
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
              <BreadcrumbLink asChild>
                <Link to="/dashboard/models">Models</Link>
              </BreadcrumbLink>
            </BreadcrumbItem>
            <BreadcrumbSeparator />
            <BreadcrumbItem>
              <BreadcrumbPage>Create</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </Header>
      <div className="flex flex-col sm:gap-4 sm:py-4">
        <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)}>
              <div className="mx-auto grid w-full max-w-[59rem] flex-1 auto-rows-max gap-4">
                <div className="flex items-center gap-4">
                  <Button variant="outline" size="icon" className="h-7 w-7" type="button" onClick={onDiscard}>
                    <ChevronLeft className="h-4 w-4" />
                    <span className="sr-only">Back</span>
                  </Button>
                  <h1 className="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0">
                    Create Model
                  </h1>
                  <div className="hidden items-center gap-2 md:ml-auto md:flex">
                    <Button variant="outline" size="sm" type="button" onClick={onDiscard}>
                      Discard
                    </Button>
                    <Button size="sm" type="submit">
                      Save Model
                    </Button>
                  </div>
                </div>
                <FormModel form={form} />
                <div className="flex items-center justify-center gap-2 md:hidden">
                  <Button variant="outline" size="sm" type="button" onClick={onDiscard}>
                    Discard
                  </Button>
                  <Button size="sm" type="submit">
                    Save Model
                  </Button>
                </div>
              </div>
            </form>
          </Form>
        </main>
      </div>
    </div>
  );
}
