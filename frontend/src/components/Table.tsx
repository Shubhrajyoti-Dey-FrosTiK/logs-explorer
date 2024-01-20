import {
  MantineReactTable,
  useMantineReactTable,
  type MRT_ColumnDef,
} from "mantine-react-table";
import { Log } from "../interfaces/logs";

export default function Table({
  headers,
  data,
}: {
  headers: MRT_ColumnDef<Log>[];
  data: any;
}) {
  //pass table options to useMantineReactTable
  const table = useMantineReactTable({
    columns: headers,
    data,
    enableRowSelection: false,
    enableColumnOrdering: false,
    enableTopToolbar: false,
    enablePagination: false,
    initialState: {
      density: "xs",
    },
  });

  return (
    <div className="mt-10">
      <MantineReactTable table={table} />
    </div>
  );
}
