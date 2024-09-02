import { PaginationState } from "@tanstack/react-table";
import * as React from "react";

import { toastValidation } from "@lib/utils";

import * as apiTokenService from "@services/api-token";

import * as queries from "@hooks/queries";

import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@components/ui/card";

import Table from "@components/partials/datatable";

import { columns, useTable } from "./components/datatable";
import { FormApiToken, FormApiTokenValues } from "./components/form";

export default function ApiTokenSettingsPage() {
  const [pagination, setPagination] = React.useState<PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  });

  const { data: response, refetch } = queries.useApiTokens({
    page: pagination.pageIndex + 1,
    per_page: pagination.pageSize,
  });

  const table = useTable({
    data: response?.data.data || [],
    columns,
    rowCount: response?.data.meta.total || 0,
    pagination,
    onPaginationChange: setPagination,
  });

  function onSubmit(values: FormApiTokenValues) {
    toastValidation(apiTokenService.create(values), {
      success() {
        refetch();
      },
    });
  }

  return (
    <div className="grid gap-6">
      <Card x-chunk="dashboard-04-chunk-1">
        <CardHeader>
          <div className="flex justify-between">
            <div className="space-y-1.5">
              <CardTitle>API Tokens</CardTitle>
              <CardDescription>Used to authenticate requests to the API.</CardDescription>
            </div>
            <div className="flex items-center space-x-2">
              <FormApiToken onSubmit={onSubmit} />
            </div>
          </div>
        </CardHeader>
        <CardContent>
          <Table.DataTable table={table} />
        </CardContent>
        <CardFooter>
          <Table.Pagination table={table} total={response?.data.meta.total || 0} disableSelectCount />
        </CardFooter>
      </Card>
    </div>
  );
}
