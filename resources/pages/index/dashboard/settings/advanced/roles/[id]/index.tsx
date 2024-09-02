import { ChevronLeft } from "lucide-react";
import { Link, useNavigate, useParams } from "react-router-dom";

import { toastValidation } from "@lib/utils";

import { Role } from "@schemas/role";

import * as roleService from "@services/role";

import * as queries from "@hooks/queries";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@components/ui/alert-dialog";
import { Button } from "@components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";

import { Error, Loading } from "@components/partials/utils";

import { ShowRole } from "../components/show";

export default function HeaderShowRolePage() {
  const { id } = useParams();
  const result = queries.useRole(Number(id || 0));

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

function Conditional(props: { result: ReturnType<typeof queries.useRole> }) {
  const { isLoading, isError, error, data: response } = props.result;

  if (isLoading) {
    return <Loading />;
  }

  if (isError) {
    return <Error response={error as Response} />;
  }

  if (!response?.data.data) {
    return <div>Role data is empty</div>;
  }

  return <ShowRolePage data={response!.data.data!} />;
}

function ShowRolePage({ data }: { data: Role }) {
  const navigate = useNavigate();

  function goBack() {
    navigate("/dashboard/settings/advanced/roles");
  }

  function deleteRole() {
    toastValidation(roleService.remove(data.id), {
      success() {
        goBack();
      },
    });
  }

  return (
    <div className="mx-auto grid w-full max-w-[59rem] flex-1 auto-rows-max gap-4">
      <div className="flex items-center gap-4">
        <Button variant="outline" size="icon" className="h-7 w-7" type="button" onClick={goBack}>
          <ChevronLeft className="h-4 w-4" />
          <span className="sr-only">Back</span>
        </Button>
        <h1 className="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0">
          Role Detail
        </h1>
        <div className="hidden items-center gap-2 md:ml-auto md:flex">
          <Button size="sm" type="submit" asChild>
            <Link to={`/dashboard/settings/advanced/roles/${data.id}/edit`}>Edit Role</Link>
          </Button>
        </div>
      </div>
      <ShowRole data={data}>
        <Card x-chunk="dashboard-07-chunk-5">
          <CardHeader>
            <CardTitle>Delete Role</CardTitle>
            <CardDescription>If you delete this role, all associated data will be lost.</CardDescription>
          </CardHeader>
          <CardContent>
            <div></div>

            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button size="sm" variant="destructive">
                  Delete Role
                </Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                  <AlertDialogDescription>
                    This action cannot be undone. This will permanently delete the role and remove it from our servers.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction onClick={deleteRole}>Continue</AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          </CardContent>
        </Card>
      </ShowRole>
      <div className="flex items-center justify-center gap-2 md:hidden">
        <Button size="sm" type="submit" asChild>
          <Link to={`/dashboard/settings/advanced/roles/${data.id}/edit`}>Edit Role</Link>
        </Button>
      </div>
    </div>
  );
}
