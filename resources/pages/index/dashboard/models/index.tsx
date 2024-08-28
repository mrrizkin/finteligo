import { PaginationState } from "@tanstack/react-table";
import { PlusCircle } from "lucide-react";
import * as React from "react";
import { Link } from "react-router-dom";

import * as queries from "@hooks/queries";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@components/ui/breadcrumb";
import { Button } from "@components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@components/ui/card";

import Table from "@components/partials/datatable";
import Header from "@components/partials/header";

import { columns, useTable } from "./components/datatable";

export default function ModelsPage() {
  const { data: response } = queries.useModels();
  const [pagination, setPagination] = React.useState<PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  });

  const table = useTable({
    data: response?.data.data || [],
    columns,
    rowCount: response?.data.meta.total || 0,
    pagination,
    onPaginationChange: setPagination,
  });

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
              <BreadcrumbPage>Models</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </Header>
      <div className="flex flex-col sm:gap-4 sm:py-4">
        <main className="grid flex-1 items-start gap-4 p-4 sm:py-0 md:gap-8">
          <div className="flex items-center">
            <div>
              <h2 className="text-lg font-semibold">LLM Models</h2>
              <p className="text-sm text-muted-foreground">A list of LLM your models and their current status.</p>
            </div>
            <div className="ml-auto flex items-center gap-2">
              <Button size="sm" className="h-7 gap-1" asChild>
                <Link to="/dashboard/models/create">
                  <PlusCircle className="h-3.5 w-3.5" />
                  <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">Add Model</span>
                </Link>
              </Button>
            </div>
          </div>
          <Card x-chunk="dashboard-06-chunk-0">
            <CardHeader>
              <CardTitle>Models</CardTitle>
              <CardDescription>Manage your model and view their status.</CardDescription>
            </CardHeader>
            <CardContent>
              <Table.DataTable table={table} />
            </CardContent>
            <CardFooter>
              <Table.Pagination table={table} total={response?.data.meta.total || 0} disableSelectCount />
            </CardFooter>
          </Card>
        </main>
      </div>
    </div>
  );
}
