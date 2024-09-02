import {
  ColumnDef,
  PaginationState,
  Updater,
  VisibilityState,
  getCoreRowModel,
  getExpandedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Clipboard, Ellipsis, LockKeyhole, LockKeyholeOpen, Trash } from "lucide-react";
import * as React from "react";
import { toast } from "sonner";

import { copyTextToClipboard, toastValidation } from "@lib/utils";

import { ApiToken } from "@schemas/api-token";

import * as apiTokenService from "@services/api-token";

import { Button } from "@components/ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@components/ui/dropdown-menu";
import { Switch } from "@components/ui/switch";

import { DataTableColumnHeader } from "@components/partials/data-table-column-header";

interface UseTableProps {
  data: ApiToken[];
  columns: ColumnDef<ApiToken>[];
  rowCount: number;
  pagination: PaginationState;
  onPaginationChange: (pagination: Updater<PaginationState>) => void;
}

export function useTable(props: UseTableProps) {
  const [columnVisibility, setColumnVisibility] = React.useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = React.useState({});

  return useReactTable({
    data: props.data,
    columns: props.columns,
    rowCount: props.rowCount,
    manualPagination: true,
    getCoreRowModel: getCoreRowModel(),
    getExpandedRowModel: getExpandedRowModel(),
    getRowCanExpand: () => true,
    onPaginationChange: props.onPaginationChange,
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      pagination: props.pagination,
      columnVisibility,
      rowSelection,
    },
  });
}

function copyTokenToClipboard(id: number) {
  toastValidation(apiTokenService.find(id), {
    success(data) {
      copyTextToClipboard(data.data?.token || "token not found");
    },
  });
}

function disableAPIToken(id: number) {
  toastValidation(apiTokenService.disable(id), {
    success() {
      toast.success("API token sucessfully disabled");
    },
  });
}

function enableAPIToken(id: number) {
  toastValidation(apiTokenService.enable(id), {
    success() {
      toast.success("API token sucessfully enabled");
    },
  });
}

function deleteAPIToken(id: number) {
  toastValidation(apiTokenService.remove(id), {
    success() {
      toast.success("API token sucessfully deleted");
    },
  });
}

export const columns: ColumnDef<ApiToken>[] = [
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Key" />;
    },
    enableHiding: false,
    accessorKey: "key",
    cell: ({ row }) => {
      return (
        <div className="flex flex-col">
          <span>{row.original.key}</span>
          <code className="text-xs text-muted-foreground">{row.original.token}</code>
        </div>
      );
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Created By" />;
    },
    accessorKey: "user.name",
    cell: ({ row }) => {
      return (
        <div className="flex flex-col">
          <span>{row.original.user?.name}</span>
          <code className="text-xs text-muted-foreground">{row.original.user?.email}</code>
        </div>
      );
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Created At" />;
    },
    accessorKey: "created_at",
    cell: ({ row }) => {
      const date = new Date(row.original.created_at).toLocaleString("en-US", {
        day: "numeric",
        month: "short",
        year: "numeric",
      });
      return <span className="text-muted-foreground">{date}</span>;
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Group" />;
    },
    accessorKey: "group",
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Enabled" />;
    },
    accessorKey: "enabled",
    cell: ({ row }) => {
      return <Switch checked={row.original.enabled} />;
    },
  },
  {
    id: "action",
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="" />;
    },
    enableHiding: false,
    cell: ({ row }) => {
      return (
        <div className="space-x-2">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="secondary" size="sm">
                <Ellipsis className="size-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuItem onClick={() => copyTokenToClipboard(row.original.id || 0)}>
                <Clipboard className="mr-2 size-4" />
                Copy To Clipboard
              </DropdownMenuItem>
              {row.original.enabled ? (
                <DropdownMenuItem onClick={() => disableAPIToken(row.original.id || 0)}>
                  <LockKeyhole className="mr-2 size-4" />
                  Disable API Token
                </DropdownMenuItem>
              ) : (
                <DropdownMenuItem onClick={() => enableAPIToken(row.original.id || 0)}>
                  <LockKeyholeOpen className="mr-2 size-4" />
                  Enable API Token
                </DropdownMenuItem>
              )}
              <DropdownMenuItem onClick={() => deleteAPIToken(row.original.id || 0)}>
                <Trash className="mr-2 size-4" />
                Delete API Token
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      );
    },
  },
];
