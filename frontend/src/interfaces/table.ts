export interface TableHeader {
  accessorKey: string;
  header: string;
}

export interface TableFilter {
  level: string;
  levelRegex: string;
  message: string;
  messageRegex: string;
  resourceId: string;
  resourceIdRegex: string;
  timestamp: string;
  timestampRegex: string;
  traceId: string;
  traceIdRegex: string;
  spanId: string;
  spanIdRegex: string;
  commit: string;
  commitRegex: string;
  parentResourceId: string;
  parentResourceIdRegex: string;
  timeStart: string;
  timeEnd: string;
  fullTextSearch: string;
  pageSize: number;
  pageNumber: number;
}
