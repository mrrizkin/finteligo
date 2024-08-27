import { ChevronLeftIcon, ChevronRightIcon } from "@radix-ui/react-icons";
import { Table } from "@tanstack/react-table";

import { Button } from "@components/ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@components/ui/select";

interface DataTablePaginationProps<TData> {
  table: Table<TData>;
  total?: number;
  disableSelectCount?: boolean;
}

export default function Pagination<TData>(props: DataTablePaginationProps<TData>) {
  return (
    <div className="flex w-full items-center justify-between px-2">
      <div className="flex-1 text-sm text-muted-foreground">
        {props.disableSelectCount ? (
          <div className="text-sm text-muted-foreground">
            {props.total || props.table.getFilteredRowModel().rows.length} row(s) found.
          </div>
        ) : (
          <div className="text-xs text-muted-foreground">
            {props.table.getFilteredSelectedRowModel().rows.length} of{" "}
            {props.total || props.table.getFilteredRowModel().rows.length} row(s) selected.
          </div>
        )}
      </div>
      <div className="flex items-center space-x-6 lg:space-x-8">
        <div className="flex items-center space-x-2">
          <p className="text-sm font-medium">Show</p>
          <Select
            value={`${props.table.getState().pagination.pageSize}`}
            onValueChange={(value: string | undefined) => {
              props.table.setPageSize(Number(value));
            }}>
            <SelectTrigger className="h-8 w-[70px]">
              <SelectValue placeholder={props.table.getState().pagination.pageSize} />
            </SelectTrigger>
            <SelectContent side="top">
              {[5, 10, 20, 30, 40, 50, 100].map((pageSize) => (
                <SelectItem key={pageSize} value={`${pageSize}`}>
                  {pageSize}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>
        <div className="flex w-[100px] items-center justify-center text-sm font-medium">
          Page {props.table.getState().pagination.pageIndex + 1} of {props.table.getPageCount()}
        </div>
        <div className="flex items-center space-x-2">
          <Button
            variant="outline"
            className="h-8 w-8 p-0"
            onClick={() => props.table.previousPage()}
            disabled={!props.table.getCanPreviousPage()}>
            <span className="sr-only">Go to previous page</span>
            <ChevronLeftIcon className="h-4 w-4" />
          </Button>
          <Button
            variant="outline"
            className="h-8 w-8 p-0"
            onClick={() => props.table.nextPage()}
            disabled={!props.table.getCanNextPage()}>
            <span className="sr-only">Go to next page</span>
            <ChevronRightIcon className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
