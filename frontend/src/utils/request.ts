export function ObjectToQuery(obj: object): string {
  let query = "";

  Object.keys(obj).map((key) => {
    // @ts-ignore
    if (obj[key] != "") query += `${query == "" ? "" : "&"}${key}=${obj[key]}`;
  });
  return query;
}
