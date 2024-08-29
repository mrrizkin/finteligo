import {
  ColumnDef,
  PaginationState,
  Updater,
  VisibilityState,
  getCoreRowModel,
  getExpandedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Ellipsis, Eye, Pencil } from "lucide-react";
import * as React from "react";
import { Link } from "react-router-dom";

import { Models } from "@schemas/models";

import { Badge } from "@components/ui/badge";
import { Button } from "@components/ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@components/ui/dropdown-menu";
import { Switch } from "@components/ui/switch";

import { DataTableColumnHeader } from "@components/partials/data-table-column-header";

interface UseTableProps {
  data: Models[];
  columns: ColumnDef<Models>[];
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

export const columns: ColumnDef<Models>[] = [
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Model" />;
    },
    enableHiding: false,
    accessorKey: "model",
    cell: ({ row }) => {
      return (
        <div className="flex flex-col">
          <span>{row.original.model}</span>
          <code className="text-xs text-muted-foreground">{row.original.token}</code>
        </div>
      );
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Provider" />;
    },
    accessorKey: "provider",
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Status" />;
    },
    accessorKey: "status",
    cell: ({ row }) => {
      let variant: "default" | "secondary" | "destructive" | "outline" = "default";
      switch (row.original.status) {
        case "ok":
          variant = "default";
          break;
        case "error":
          variant = "destructive";
          break;
        case "pending":
          variant = "secondary";
          break;
      }

      return <Badge variant={variant}>{row.original.status}</Badge>;
    },
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
              <DropdownMenuItem asChild>
                <Link to={`/dashboard/models/${row.original.id}`}>
                  <Eye className="mr-2 size-4" />
                  Show
                </Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link to={`/dashboard/models/${row.original.id}/edit`}>
                  <Pencil className="mr-2 size-4" />
                  Edit
                </Link>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      );
    },
  },
];
