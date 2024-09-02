import { zodResolver } from "@hookform/resolvers/zod";
import { ChevronLeft } from "lucide-react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

import { toastValidation } from "@lib/utils";

import * as roleService from "@services/role";

import { Button } from "@components/ui/button";
import { Form } from "@components/ui/form";

import { FormRole, FormRoleValues, formSchema } from "./components/form";

export default function CreateRolePage() {
  const navigate = useNavigate();
  const form = useForm<FormRoleValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      slug: "",
      permission_ids: [],
    },
  });

  function onSubmit(values: FormRoleValues) {
    toastValidation(roleService.create(values), {
      success() {
        form.reset();
        navigate("/dashboard/settings/advanced/roles");
      },
    });
  }

  function onDiscard() {
    form.reset();
    navigate("/dashboard/settings/advanced/roles");
  }

  return (
    <div className="flex flex-col">
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
                    Create Role
                  </h1>
                  <div className="hidden items-center gap-2 md:ml-auto md:flex">
                    <Button variant="outline" size="sm" type="button" onClick={onDiscard}>
                      Discard
                    </Button>
                    <Button size="sm" type="submit">
                      Save Role
                    </Button>
                  </div>
                </div>
                <FormRole form={form} />
                <div className="flex items-center justify-center gap-2 md:hidden">
                  <Button variant="outline" size="sm" type="button" onClick={onDiscard}>
                    Discard
                  </Button>
                  <Button size="sm" type="submit">
                    Save Role
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
