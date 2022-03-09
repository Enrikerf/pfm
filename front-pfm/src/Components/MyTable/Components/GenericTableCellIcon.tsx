import IconButton from "@mui/material/IconButton";
import ArrowCircleRightIcon from "@mui/icons-material/ArrowCircleRight";
import TableCell from "@mui/material/TableCell";
import * as React from "react";
import {TableData} from "../TableData";


export default function GenericTableCellIcon(
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
        >
            <IconButton
                onClick={
                    (event) =>
                        props.handleGoTo(event,props.id,props.tableData)
                }
            >
                <ArrowCircleRightIcon/>
            </IconButton>
        </TableCell>
    )
}