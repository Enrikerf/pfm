import * as React from "react";
import {TableOrder} from "./TableOrder";
import {TableData} from "./TableData";

export interface EnhancedTableProps {
    numSelected: number;
    onRequestSort: (event: React.MouseEvent<unknown>, property: keyof TableData) => void;
    onSelectAllClick: (event: React.ChangeEvent<HTMLInputElement>) => void;
    order: TableOrder;
    orderBy: string;
    rowCount: number;
}