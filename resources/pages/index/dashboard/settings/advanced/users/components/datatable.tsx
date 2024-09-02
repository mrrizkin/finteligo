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

import { User } from "@schemas/user";

import { Button } from "@components/ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@components/ui/dropdown-menu";

import { DataTableColumnHeader } from "@components/partials/data-table-column-header";

interface UseTableProps {
  data: User[];
  columns: ColumnDef<User>[];
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

export const columns: ColumnDef<User>[] = [
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Name" />;
    },
    enableHiding: false,
    accessorKey: "username",
    cell: ({ row }) => {
      return (
        <div className="flex flex-col">
          <span>{row.original.name}</span>
          <code className="text-xs text-muted-foreground">{row.original.username}</code>
        </div>
      );
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Email" />;
    },
    accessorKey: "created_at",
    cell: ({ row }) => {
      return <span className="text-muted-foreground">{row.original.email}</span>;
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Role" />;
    },
    accessorKey: "role.name",
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
                <Link to={`/dashboard/settings/advanced/users/${row.original.id}`}>
                  <Eye className="mr-2 size-4" />
                  Show
                </Link>
              </DropdownMenuItem>
              <DropdownMenuItem asChild>
                <Link to={`/dashboard/settings/advanced/users/${row.original.id}/edit`}>
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
