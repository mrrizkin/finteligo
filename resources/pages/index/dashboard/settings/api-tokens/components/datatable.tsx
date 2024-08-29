import {
  ColumnDef,
  PaginationState,
  Updater,
  VisibilityState,
  getCoreRowModel,
  getExpandedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Ban, Clipboard, Ellipsis, Trash } from "lucide-react";
import * as React from "react";

import { ApiToken } from "@schemas/api-token";

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

export const columns: ColumnDef<ApiToken>[] = [
  {
    id: "action",
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="" />;
    },
    enableHiding: false,
    cell: () => {
      return (
        <div className="space-x-2">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="secondary" size="sm">
                <Ellipsis className="size-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>
                <Clipboard className="mr-2 size-4" />
                Copy To Clipboard
              </DropdownMenuItem>
              <DropdownMenuItem>
                <Ban className="mr-2 size-4" />
                Disable API Token
              </DropdownMenuItem>
              <DropdownMenuItem>
                <Trash className="mr-2 size-4" />
                Delete API Token
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      );
    },
  },
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
];
