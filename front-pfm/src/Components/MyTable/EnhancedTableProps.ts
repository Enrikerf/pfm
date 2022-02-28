import * as React from "react";
import {TableOrder} from "./TableOrder";
import {TableData} from "./TableData";
import {HeadCell} from "./HeadCell";

export interface EnhancedTableProps {
    heads: HeadCell[];
    numSelected: number;
    onRequestSort: (event: React.MouseEvent<unknown>, property: keyof TableData) => void;
    onSelectAllClick: (event: React.ChangeEvent<HTMLInputElement>) => void;
    order: TableOrder;
    orderBy: string;
    rowCount: number;
}