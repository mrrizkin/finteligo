import * as queries from "@hooks/queries";
import { ChevronLeft } from "lucide-react";
import { Link, useNavigate, useParams } from "react-router-dom";

import { Models } from "@schemas/models";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@components/ui/breadcrumb";
import { Button } from "@components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@components/ui/card";

import Header from "@components/partials/header";
import { Error, Loading } from "@components/partials/utils";

import { ShowModel } from "../show";

export default function EditModelPage() {
  const { id } = useParams();
  const result = queries.useModel(Number(id || 0));

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
              <BreadcrumbPage>{id}</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </Header>
      <div className="flex flex-1 flex-col sm:gap-4 sm:py-4">
        <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8">
          <Conditional result={result} />
        </main>
      </div>
    </div>
  );
}

function Conditional(props: { result: ReturnType<typeof queries.useModel> }) {
  const { isLoading, isError, error, data: response } = props.result;

  if (isLoading) {
    return <Loading />;
  }

  if (isError) {
    return <Error response={error as Response} />;
  }

  if (!response?.data.data) {
    return <div>Model data is empty</div>;
  }

  return <ShowModelPage data={response!.data.data!} />;
}

function ShowModelPage({ data }: { data: Models }) {
  const navigate = useNavigate();

  function goBack() {
    navigate("/dashboard/models");
  }

  return (
    <div className="mx-auto grid w-full max-w-[59rem] flex-1 auto-rows-max gap-4">
      <div className="flex items-center gap-4">
        <Button variant="outline" size="icon" className="h-7 w-7" type="button" onClick={goBack}>
          <ChevronLeft className="h-4 w-4" />
          <span className="sr-only">Back</span>
        </Button>
        <h1 className="flex-1 shrink-0 whitespace-nowrap text-xl font-semibold tracking-tight sm:grow-0">
          Model Detail
        </h1>
        <div className="hidden items-center gap-2 md:ml-auto md:flex">
          <Button size="sm" type="submit" asChild>
            <Link to={`/dashboard/models/${data.id}/edit`}>Edit Model</Link>
          </Button>
        </div>
      </div>
      <ShowModel data={data}>
        <Card x-chunk="dashboard-07-chunk-5">
          <CardHeader>
            <CardTitle>Delete Model</CardTitle>
            <CardDescription>If you delete this model, all associated data will be lost.</CardDescription>
          </CardHeader>
          <CardContent>
            <div></div>
            <Button size="sm" variant="destructive">
              Delete Model
            </Button>
          </CardContent>
        </Card>
      </ShowModel>
      <div className="flex items-center justify-center gap-2 md:hidden">
        <Button size="sm" type="submit" asChild>
          <Link to={`/dashboard/models/${data.id}/edit`}>Edit Model</Link>
        </Button>
      </div>
    </div>
  );
}
