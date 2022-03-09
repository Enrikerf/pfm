import IconButton from "@mui/material/IconButton";
import ArrowCircleRightIcon from "@mui/icons-material/ArrowCircleRight";
import TableCell from "@mui/material/TableCell";
import * as React from "react";
import {TableData} from "../TableData";


export default function GenericTableCell(
    props: {
        id: number,
        tableData: TableData,
        align: any,
        handleGoTo: (event: React.MouseEvent<unknown>, id: number, toGo: TableData) => void
    }
) {
    return (
        <TableCell
            key={props.id}
            component="th"
            scope="row"
            padding="none"
            align={props.align}
            onClick={
                (event) =>
                    props.handleGoTo(event, props.id, props.tableData.value)
            }
        >
            {props.tableData.value}
        </TableCell>
    )
}