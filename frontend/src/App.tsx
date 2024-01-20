import { createContext, useEffect, useState } from "react";
import { LatestLogListResponse, Log } from "./interfaces/logs";
import Table from "./components/Table";
import { InitialFilters, TableHeaders } from "./constants/table";
import { TableFilter } from "./interfaces/table";
import { Modal, Tabs } from "@mantine/core/";
import useWebSocket from "react-use-websocket";
import { useSelector } from "react-redux";
import { selectFilter } from "./store/states/filterSlice";
import { useDisclosure } from "@mantine/hooks";
import FilterForm from "./components/FilterForm";
import { IconFilter } from "@tabler/icons-react";
import axios from "axios";
import { ObjectToQuery } from "./utils/request";

export const FilterContext = createContext<TableFilter>(InitialFilters);

function compressedFilter(filter: TableFilter): object {
  let compressedFilter = {};
  Object.keys(filter).forEach((key: string) => {
    // @ts-ignore
    if (filter[key] != "" && filter[key] != "2001-01-01T00:00:00.00Z")
      // @ts-ignore
      compressedFilter[key] = filter[key];
  });

  return compressedFilter;
}

function App() {
  const Filters = useSelector(selectFilter);
  const [logs, setLogs] = useState<Log[]>([]);
  const [fullTextFilterdLogs, setFullTextFilterdLogs] = useState<Log[]>([]);
  const [opened, { open, close }] = useDisclosure(false);
  const [tab, setTab] = useState("rtl");

  const { sendMessage, lastMessage } = useWebSocket(
    `ws://${import.meta.env.VITE_BACKEND}/ws`
  );

  // Fetch all the latest 60 logs first
  const fetchLatestLogs = async () => {
    console.log(compressedFilter(Filters.filter));
    if (tab == "rtl") {
      sendMessage(JSON.stringify(compressedFilter(Filters.filter)));
    } else {
      let query = ObjectToQuery(Filters.filter);

      console.log(`http://${import.meta.env.VITE_BACKEND}/search?${query}`);

      const response: LatestLogListResponse = await axios.get(
        `http://${import.meta.env.VITE_BACKEND}/search?${query}`
      );

      if (!response.data.err && response.data.logs) {
        setFullTextFilterdLogs(response.data.logs);
      } else {
        setFullTextFilterdLogs([]);
      }
    }
  };

  // Trigger whenever any filter is applied
  useEffect(() => {
    if (Filters.filter) {
      fetchLatestLogs();
    }
  }, [Filters]);

  // Triggers when we get any message from the server
  useEffect(() => {
    if (lastMessage !== null) {
      const data = JSON.parse(lastMessage.data);
      if (data) setLogs(data);
      else setLogs([]);
    }
  }, [lastMessage]);

  return (
    <div className="p-10 text-center">
      <Modal opened={opened} onClose={close} title="Filters">
        <FilterForm
          close={close}
          intialState={Filters.filter}
          fullTextSearchEnabled={tab == "fts"}
        />
      </Modal>
      <h1 className="font-bold text-3xl mt-10 mv-">Logs Explorer</h1>
      <h1 className="text-2xl mb-10">Explore all the logs at one place</h1>

      <Tabs
        defaultValue="rtl"
        value={tab}
        onTabChange={(tab: string) => {
          setTab(tab);
        }}
      >
        <Tabs.List>
          <Tabs.Tab value="rtl">Real Time logs</Tabs.Tab>
          <Tabs.Tab value="fts">Full Text search</Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="rtl" pt="xs">
          <div className="flex justify-between flex-wrap items-center w-full">
            <div className="text-left">
              <p>Showing latest {Filters.filter.pageSize} real time logs</p>
              <p style={{ fontSize: "12px", color: "grey" }}>
                Use the dedicated filter option for server side filtering or use
                the column filter to use local filtering
              </p>
            </div>
            <span
              className=" cursor-pointer rounded-md pt-2 pb-2 pl-4 pr-4 text-sm"
              onClick={open}
            >
              <IconFilter />
            </span>
          </div>
          <Table headers={TableHeaders} data={logs} />
        </Tabs.Panel>

        <Tabs.Panel value="fts" pt="xs">
          <div className="flex items-center justify-center md:justify-between flex-wrap w-full">
            <div className="flex justify-between flex-wrap items-center w-full">
              <div className="text-left">
                <p>Showing latest {Filters.filter.pageSize} real time logs</p>
                <p style={{ fontSize: "12px", color: "grey" }}>
                  Use the dedicated filter option for server side filtering or
                  use the column filter to use local filtering
                </p>
              </div>
              <span
                className=" cursor-pointer rounded-md pt-2 pb-2 pl-4 pr-4 text-sm"
                onClick={open}
              >
                <IconFilter />
              </span>
            </div>
          </div>
          <Table headers={TableHeaders} data={fullTextFilterdLogs} />
        </Tabs.Panel>
      </Tabs>
    </div>
  );
}

export default App;
