import { MRT_ColumnDef } from "mantine-react-table";
import { TableFilter } from "../interfaces/table";
import { Log } from "../interfaces/logs";

export const TableHeaders: MRT_ColumnDef<Log>[] = [
  {
    accessorKey: "level",
    header: "Level",
  },
  {
    accessorKey: "message",
    header: "Message",
  },
  {
    accessorKey: "resourceId",
    header: "Resource ID",
  },
  {
    accessorKey: "timestamp",
    header: "Timestamp",
  },
  {
    accessorKey: "traceId",
    header: "Trace ID",
  },
  {
    accessorKey: "spanId",
    header: "Span ID",
  },
  {
    accessorKey: "commit",
    header: "Commit",
  },
  {
    accessorKey: "metadata.parentResourceId",
    header: "Parent Resource ID",
  },
];

export const InitialFilters: TableFilter = {
  level: "",
  levelRegex: "",
  message: "",
  messageRegex: "",
  resourceId: "",
  resourceIdRegex: "",
  timestamp: "",
  timestampRegex: "",
  traceId: "",
  traceIdRegex: "",
  spanId: "",
  spanIdRegex: "",
  commit: "",
  commitRegex: "",
  parentResourceId: "",
  parentResourceIdRegex: "",
  timeStart: "",
  timeEnd: "",
  fullTextSearch: "",
  pageNumber: 0,
  pageSize: 60,
};

export const FilterArray: string[] = [
  "level",
  "message",
  "resourceId",
  "timestamp",
  "traceId",
  "spanId",
  "commit",
  "parentResourceId",
];
