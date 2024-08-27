import { Table } from "@tanstack/react-table";
import { Eraser, Trash } from "lucide-react";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@components/ui/alert-dialog";
import { Button } from "@components/ui/button";

import { Show } from "@components/show";

interface SelectionControlProps<T> {
  table: Table<T>;
  onDeleteBatchSubmit: (values: T[]) => void;
}

export default function SelectionControl<T>(props: SelectionControlProps<T>) {
  function clearSelection() {
    props.table.setRowSelection({});
  }

  function deleteSelected() {
    const selectedRows = props.table.getFilteredSelectedRowModel().rows;
    props.onDeleteBatchSubmit(selectedRows.map((value) => value.original));
  }

  return (
    <Show when={props.table.getIsSomeRowsSelected() || props.table.getIsAllRowsSelected()}>
      <AlertDialog>
        <AlertDialogTrigger asChild>
          <Button variant="destructive">
            <Trash className="mr-2 h-4 w-4" />
            Delete selected
          </Button>
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Apakah Anda benar-benar yakin?</AlertDialogTitle>
            <AlertDialogDescription>
              Tindakan ini tidak bisa dibatalkan. Tindakan ini akan menghapus{" "}
              {props.table.getFilteredSelectedRowModel().rows.length} data secara permanen.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Batal</AlertDialogCancel>
            <AlertDialogAction asChild onClick={deleteSelected}>
              <Button variant="destructive">Lanjutkan</Button>
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
      <Button variant="outline" onClick={clearSelection}>
        <Eraser className="mr-2 h-4 w-4" />
        Clear selection
      </Button>
    </Show>
  );
}
