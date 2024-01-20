import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { RootState } from "../store";

import { TableFilter } from "../../interfaces/table";
import { InitialFilters } from "../../constants/table";

// Define the initial state using that type
const initialState: { filter: TableFilter } = { filter: InitialFilters };

export const filterSlice = createSlice({
  name: "filter",
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
    // Use the PayloadAction type to declare the contents of `action.payload`
    setCurrentFilter: (state: any, action: PayloadAction<TableFilter>) => {
      state.filter = action.payload;
    },

    setFilterField: (
      state: { filter: TableFilter },
      action: PayloadAction<{
        key: string;
        value: string;
      }>
    ) => {
      // @ts-ignore
      state.filter[action.payload.key] = action.payload.value;
    },

    unsetFilterField: (
      state: { filter: TableFilter },
      action: PayloadAction<{
        key: string;
      }>
    ) => {
      // @ts-ignore
      state.filter[action.payload.key] = "";
    },
  },
});

export const { setCurrentFilter, setFilterField, unsetFilterField } =
  filterSlice.actions;

// Other code such as selectors can use the imported `RootState` type
export const selectFilter = (state: RootState) => state.filter;

export default filterSlice;
