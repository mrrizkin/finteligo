import {
  ColumnDef,
  PaginationState,
  Updater,
  VisibilityState,
  getCoreRowModel,
  getExpandedRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Eye, Pencil } from "lucide-react";
import * as React from "react";
import { Link } from "react-router-dom";

import { Models } from "@schemas/models";

import { Button } from "@components/ui/button";
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
    id: "action",
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Action" />;
    },
    enableHiding: false,
    cell: ({ row }) => {
      return (
        <div className="space-x-2">
          <Button variant="secondary" size="sm" asChild>
            <Link to={`/dashboard/models/${row.original.id}`}>
              <Eye className="size-4" />
            </Link>
          </Button>
          <Button variant="secondary" size="sm" asChild>
            <Link to={`/dashboard/models/${row.original.id}/edit`}>
              <Pencil className="size-4" />
            </Link>
          </Button>
        </div>
      );
    },
  },
  {
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Model" />;
    },
    enableHiding: false,
    accessorKey: "model",
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
