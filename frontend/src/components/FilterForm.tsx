import { useState } from "react";
import { FilterArray } from "../constants/table";
import { Input, NumberInput, Switch, TextInput } from "@mantine/core";
import { TableFilter } from "../interfaces/table";
import { useDispatch } from "react-redux";
import { setCurrentFilter } from "../store/states/filterSlice";
import { DateTimePicker } from "@mantine/dates";

function FilterForm({
  close,
  intialState,
  fullTextSearchEnabled,
}: {
  close: () => void;
  intialState: TableFilter;
  fullTextSearchEnabled?: boolean;
}) {
  const dispatch = useDispatch();
  const [regexSearchStatus, setRegexSearchStatus] = useState<{
    [key: string]: boolean;
  }>({
    level: false,
    message: false,
    resourceId: false,
    timestamp: false,
    traceId: false,
    spanId: false,
    commit: false,
    parentResourceId: false,
  });

  const [filterData, setFilter] = useState<TableFilter>(intialState);

  const applyFilter = () => {
    dispatch(setCurrentFilter(filterData));
    close();
  };

  return (
    <div className="p-5">
      {FilterArray.map((filter: string, filterIndex: number) => {
        return (
          <div key={`Filter_${filterIndex}`} className="mb-2">
            <div className="flex justify-between">
              <Input.Label className="font-bold">
                {filter.toUpperCase()}
              </Input.Label>
              <Switch
                checked={regexSearchStatus[filter]}
                onChange={() => {
                  // @ts-ignore
                  const input = filterData[filter];

                  if (regexSearchStatus[filter])
                    setFilter({
                      ...filterData,
                      [`${filter}`]: input,
                      [`${filter}Regex`]: "",
                    });
                  else
                    setFilter({
                      ...filterData,
                      [`${filter}`]: "",
                      [`${filter}Regex`]: input,
                    });

                  setRegexSearchStatus({
                    ...regexSearchStatus,
                    [filter]: !regexSearchStatus[filter],
                  });
                }}
                labelPosition="left"
                label="Regex Search"
              />
            </div>
            <TextInput
              placeholder={filter}
              key={filterIndex}
              className="w-full"
              value={
                regexSearchStatus[filter]
                  ? // @ts-ignore
                    filterData[`${filter}Regex`]
                  : // @ts-ignore
                    filterData[filter]?.toString()
              }
              onChange={(e: any) => {
                if (regexSearchStatus[filter]) {
                  setFilter({
                    ...filterData,
                    [`${filter}Regex`]: e.target.value,
                  });
                } else {
                  setFilter({
                    ...filterData,
                    [filter]: e.target.value,
                  });
                }
              }}
            />
          </div>
        );
      })}
      <div className="mb-2">
        <DateTimePicker
          valueFormat="YYYY-MM-DDThh:mm:ssZ"
          label="TIMESTAMP"
          clearable
          placeholder="Pick date and time"
          value={
            filterData.timestamp == ""
              ? undefined
              : new Date(filterData.timestamp)
          }
          onChange={(value) => {
            setFilter({ ...filterData, timestamp: value?.toISOString() || "" });
          }}
        />
      </div>
      <div className="mb-2">
        <DateTimePicker
          valueFormat="YYYY-MM-DDThh:mm:ssZ"
          label="LOGS AFTER"
          clearable
          placeholder="Pick date and time"
          value={
            filterData.timeStart == ""
              ? undefined
              : new Date(filterData.timeStart)
          }
          onChange={(value) => {
            setFilter({
              ...filterData,
              timeStart: value ? value?.toISOString() : "",
            });
          }}
        />
      </div>
      <div className="mb-2">
        <DateTimePicker
          clearable
          valueFormat="YYYY-MM-DDThh:mm:ssZ"
          label="LOGS BEFORE"
          placeholder="Pick date and time"
          value={
            filterData.timeEnd == "" ? undefined : new Date(filterData.timeEnd)
          }
          onChange={(value) => {
            console.log(value);
            setFilter({ ...filterData, timeEnd: value?.toISOString() || "" });
          }}
        />
      </div>
      <div className="mb-2">
        <NumberInput
          label="LOG COUNT"
          placeholder="No of logs to display at max"
          value={filterData.pageSize}
          onChange={(value: number) => {
            setFilter({ ...filterData, pageSize: value });
          }}
        />
      </div>
      {fullTextSearchEnabled && (
        <div>
          <Input.Label className="font-bold">FULL TEXT SEARCH</Input.Label>
          <TextInput
            placeholder={"Full text search"}
            key={"fulltextSearch"}
            className="w-full"
            value={filterData.fullTextSearch}
            onChange={(e: any) => {
              setFilter({
                ...filterData,
                fullTextSearch: e.target.value,
              });
            }}
          />
        </div>
      )}
      <div className="text-center mt-10">
        <span
          onClick={applyFilter}
          className="bg-blue-500 cursor-pointer rounded-md m-auto pt-2 pb-2 pl-7 pr-7 text-white"
        >
          APPLY
        </span>
      </div>
    </div>
  );
}

export default FilterForm;
