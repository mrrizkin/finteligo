import { zodResolver } from "@hookform/resolvers/zod";
import { ChevronLeft } from "lucide-react";
import { useForm } from "react-hook-form";
import { useNavigate, useParams } from "react-router-dom";

import { toastValidation } from "@lib/utils";

import { User } from "@schemas/user";

import * as userService from "@services/user";

import * as queries from "@hooks/queries";

import { Button } from "@components/ui/button";
import { Form } from "@components/ui/form";

import { Error, Loading } from "@components/partials/utils";

import { FormUser, FormUserValues, formSchema } from "../components/form";

export default function HeaderEditUserPage() {
  const { id } = useParams();
  const result = queries.useUser(Number(id || 0));

  return (
    <div className="flex flex-col">
      <div className="flex flex-1 flex-col sm:gap-4 sm:py-4">
        <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
          <Conditional result={result} />
        </main>
      </div>
    </div>
  );
}

function Conditional(props: { result: ReturnType<typeof queries.useUser> }) {
  const { isLoading, isError, error, data: response } = props.result;

  if (isLoading) {
    return <Loading />;
  }

  if (isError) {
    return <Error response={error as Response} />;
  }

  if (!response?.data.data) {
    return <div>User data is empty</div>;
  }

  return <EditUserPage data={response!.data.data!} />;
}

function EditUserPage({ data }: { data: User }) {
  const navigate = useNavigate();
  const form = useForm<FormUserValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: data.name,
      email: data.email,
      username: data.username,
      role_id: data.role_id,
    },
  });

  function onSubmit(values: FormUserValues) {
    toastValidation(userService.update(data.id, values), {
      success() {
        form.reset();
        navigate("/dashboard/settings/advanced/users");
      },
    });
  }

  function onDiscard() {
    form.reset();
    navigate("/dashboard/settings/advanced/users");
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <div className="mx-auto grid w-full max-w-[59rem] flex-1 auto-rows-max gap-4">
          <div className="flex items-center gap-4">
            <Button variant="outline" size="icon" className="h-7 w-7" type="button" onClick={onDiscard}>
              <ChevronLeft className="h-4 w-4" />
              <span className="sr-only">Back</span>
            </Button>
            <h1 className="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0">
              Edit User
            </h1>
            <div className="hidden items-center gap-2 md:ml-auto md:flex">
              <Button variant="outline" size="sm" type="button" onClick={onDiscard}>
                Discard
              </Button>
              <Button size="sm" type="submit">
                Save User
              </Button>
            </div>
          </div>
          <FormUser form={form} />
          <div className="flex items-center justify-center gap-2 md:hidden">
            <Button variant="outline" size="sm" type="button" onClick={onDiscard}>
              Discard
            </Button>
            <Button size="sm" type="submit">
              Save User
            </Button>
          </div>
        </div>
      </form>
    </Form>
  );
}
