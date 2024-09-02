import { PaginationState } from "@tanstack/react-table";
import { PlusCircle } from "lucide-react";
import * as React from "react";
import { Link } from "react-router-dom";

import * as queries from "@hooks/queries";

import { Button } from "@components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@components/ui/card";

import Table from "@components/partials/datatable";

import { columns, useTable } from "./components/datatable";

export default function ManageUserSettingsPage() {
  const [pagination, setPagination] = React.useState<PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  });

  const { data: response } = queries.useUsers({
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

  return (
    <div className="grid gap-6">
      <Card x-chunk="dashboard-04-chunk-1">
        <CardHeader>
          <div className="flex justify-between">
            <div className="space-y-1.5">
              <CardTitle>Manage User</CardTitle>
              <CardDescription>Used to login the dashboard.</CardDescription>
            </div>
            <div className="flex items-center space-x-2">
              <Button size="sm" className="h-7 gap-1" asChild>
                <Link to="/dashboard/settings/advanced/users/create">
                  <PlusCircle className="h-3.5 w-3.5" />
                  <span className="sr-only sm:not-sr-only sm:whitespace-nowrap">Create User</span>
                </Link>
              </Button>
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
