import ColumnVisibilty from "./column-visiblity";
import DataTable from "./datatable";
import Pagination from "./pagination";
import SelectionControl from "./selection-control";

function Root(props: { children: React.ReactNode }) {
  return <div>{props.children}</div>;
}

Root.DataTable = DataTable;
Root.Pagination = Pagination;
Root.SelectionControl = SelectionControl;
Root.ColumnVisibility = ColumnVisibilty;

export default Root;
