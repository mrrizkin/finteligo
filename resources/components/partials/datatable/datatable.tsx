import { Row, Table as TanstackTable, flexRender } from "@tanstack/react-table";
import { Fragment, useMemo } from "react";

import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@components/ui/table";

import { Show } from "@components/show";

interface ChildrenProps<T> {
  row: Row<T>;
}

interface DataTableProps<T> {
  table: TanstackTable<T>;
  children?: (props: ChildrenProps<T>) => React.ReactNode;
  onRowClick?: (row: Row<T>) => void;
}

export default function DataTable<T>(props: DataTableProps<T>) {
  const columnsLength = useMemo(() => {
    let length = 0;

    for (const headerGroup of props.table.getHeaderGroups()) {
      if (headerGroup.headers.length > length) {
        length = headerGroup.headers.length;
      }
    }

    return length;
  }, [props.table]);

  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          {props.table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead
                    key={header.id}
                    className={
                      header.column.columnDef.id === "expander" ||
                      header.column.columnDef.id === "select" ||
                      header.column.columnDef.id === "action"
                        ? "w-[1%] whitespace-nowrap"
                        : ""
                    }>
                    <Show when={!header.isPlaceholder}>
                      {flexRender(header.column.columnDef.header, header.getContext())}
                    </Show>
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {props.table.getRowModel().rows?.length ? (
            props.table.getRowModel().rows.map((row) => (
              <Fragment key={row.id}>
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                  onClick={() => {
                    if (props.onRowClick) {
                      props.onRowClick(row);
                    }
                  }}>
                  {row.getVisibleCells().map((cell) => (
                    <TableCell
                      key={cell.id}
                      className={
                        (cell.column.columnDef.id === "expander" && row.getCanExpand()) ||
                        (cell.column.columnDef.id === "select" && row.getCanSelect()) ||
                        cell.column.columnDef.id === "action"
                          ? "w-[1%] whitespace-nowrap"
                          : ""
                      }>
                      {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </TableCell>
                  ))}
                </TableRow>
                <Show when={row.getIsExpanded() && !!props.children}>
                  <TableRow>
                    <TableCell colSpan={row.getVisibleCells().length} className="p-0">
                      {props.children?.({ row })}
                    </TableCell>
                  </TableRow>
                </Show>
              </Fragment>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columnsLength} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
