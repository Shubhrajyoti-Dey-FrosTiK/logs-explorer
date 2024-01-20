export interface Log {
  _id: string;
  level: string;
  message: string;
  resourceId: string;
  timestamp: string;
  traceId: string;
  spanId: string;
  commit: string;
  metadata: {
    parentResourceId: string;
  };
}

export interface LatestLogListResponse {
  data: {
    err: string;
    logs: Log[];
  };
}
