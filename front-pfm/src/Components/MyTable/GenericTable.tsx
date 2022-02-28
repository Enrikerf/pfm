import * as React from 'react';
import {useEffect} from 'react';
import Box from '@mui/material/Box';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import ArrowCircleRightIcon from '@mui/icons-material/ArrowCircleRight';
import IconButton from "@mui/material/IconButton";
import {TableOrder} from "./TableOrder";
import {TableData} from "./TableData";
import {TableRowData} from "./TableRowData";
import TableCell from '@mui/material/TableCell';
import EnhancedTableHead from "./EnhancedTableHead";
import {EnhancedTableToolbar} from "./EnhancedTableToolbar";
import {Checkbox} from "@mui/material";
import {HeadCell} from "./HeadCell";

export default function GenericTable(props: {
    rows: TableRowData[],
    handleGoTo: (event: React.MouseEvent<unknown>, id: number, toGo: TableData) => void
}) {
    const [rows, setRows] = React.useState<TableRowData[]>([]);
    const [heads, setHeads] = React.useState<HeadCell[]>([]);
    const [order, setOrder] = React.useState<TableOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('uuid');
    const [selected, setSelected] = React.useState<readonly number[]>([]);
    const [page, setPage] = React.useState(0);
    const [dense] = React.useState(false);
    const [rowsPerPage, setRowsPerPage] = React.useState(5);

    useEffect(() => {
        function getHeads(currentRows: TableRowData[]): HeadCell[] {
            if (currentRows.length <= 0) {
                return []
            }
            return currentRows[0].values.map((row, index) => {
                if (index === 0) {
                    return {
                        "id": row.name,
                        "numeric": false,
                        "disablePadding": true,
                        "label": row.name
                    }
                } else {
                    return {
                        "id": row.name,
                        "numeric": true,
                        "disablePadding": false,
                        "label": row.name
                    }
                }
            })
        }

        setRows(props.rows)
        setHeads(getHeads(props.rows))
        return () => {
        };
    }, [props.rows])


    const handleRequestSort = (
        event: React.MouseEvent<unknown>,
        property: keyof TableData,
    ) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };

    const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.checked) {
            const newSelecteds = rows.map((n) => n.id);
            setSelected(newSelecteds);
            return;
        }
        setSelected([]);
    };

    const handleClick = (event: React.MouseEvent<unknown>, id: number) => {
        const selectedIndex = selected.indexOf(id);
        let newSelected: readonly number[] = [];

        if (selectedIndex === -1) {
            newSelected = newSelected.concat(selected, id);
        } else if (selectedIndex === 0) {
            newSelected = newSelected.concat(selected.slice(1));
        } else if (selectedIndex === selected.length - 1) {
            newSelected = newSelected.concat(selected.slice(0, -1));
        } else if (selectedIndex > 0) {
            newSelected = newSelected.concat(
                selected.slice(0, selectedIndex),
                selected.slice(selectedIndex + 1),
            );
        }

        setSelected(newSelected);
    };


    const handleChangePage = (event: unknown, newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    const isSelected = (id: number) => selected.indexOf(id) !== -1;

    // Avoid a layout jump when reaching the last page with empty rows.
    const emptyRows =
        page > 0 ? Math.max(0, (1 + page) * rowsPerPage - rows.length) : 0;

    return (
        <Box sx={{width: '100%'}}>
            <Paper sx={{width: '100%', mb: 2}}>
                <EnhancedTableToolbar numSelected={selected.length}/>
                <TableContainer>
                    <Table sx={{minWidth: 750}} aria-labelledby="tableTitle" size={dense ? 'small' : 'medium'}>
                        <EnhancedTableHead
                            heads={heads}
                            numSelected={selected.length}
                            order={order}
                            orderBy={orderBy}
                            onSelectAllClick={handleSelectAllClick}
                            onRequestSort={handleRequestSort}
                            rowCount={rows.length}
                        />
                        <TableBody>
                            {
                                rows.map((row, index) => {
                                    const isItemSelected = isSelected(row.id);
                                    const labelId = `enhanced-table-checkbox-${index}`;

                                    return (
                                        <TableRow
                                            hover
                                            role="checkbox"
                                            aria-checked={isItemSelected}
                                            tabIndex={-1}
                                            key={labelId}
                                            selected={isItemSelected}
                                        >
                                            <TableCell padding="checkbox">
                                                <Checkbox
                                                    onClick={(event) => handleClick(event, row.id)}
                                                    color="primary"
                                                    checked={isItemSelected}
                                                    inputProps={{
                                                        'aria-labelledby': labelId,
                                                    }}
                                                />
                                            </TableCell>
                                            {
                                                row.values.map((value, i) => {
                                                    if (i === 0) {
                                                        if (value.value === "icon") {
                                                            return (
                                                                <TableCell key={i}  component="th"
                                                                           id={labelId}
                                                                           scope="row"
                                                                           padding="none">
                                                                    <IconButton
                                                                        onClick={
                                                                            (event) =>
                                                                                props.handleGoTo(event, row.id, value)
                                                                        }
                                                                    >
                                                                        <ArrowCircleRightIcon/>
                                                                    </IconButton>
                                                                </TableCell>
                                                            )
                                                        } else {
                                                            return (
                                                                <TableCell
                                                                    key={i}
                                                                    component="th"
                                                                    id={labelId}
                                                                    scope="row"
                                                                    padding="none"
                                                                    onClick={
                                                                        (event) =>
                                                                            props.handleGoTo(event, row.id, value)
                                                                    }
                                                                >
                                                                    {value.value}
                                                                </TableCell>)
                                                        }


                                                } else {
                                                        if (value.value === "icon") {
                                                            return (
                                                                <TableCell key={i} align="right">
                                                                    <IconButton
                                                                        onClick={
                                                                            (event) =>
                                                                                props.handleGoTo(event, row.id, value)
                                                                        }
                                                                    >
                                                                        <ArrowCircleRightIcon/>
                                                                    </IconButton>
                                                                </TableCell>
                                                            )
                                                        } else {
                                                            return (
                                                                <TableCell
                                                                    key={i}
                                                                    align="right"
                                                                    onClick={
                                                                        (event) =>
                                                                            props.handleGoTo(event, row.id, value)
                                                                    }
                                                                >
                                                                    {value.value}
                                                                </TableCell>)
                                                        }
                                                    }


                                                })
                                            }
                                        </TableRow>
                                    );
                                })}
                            {emptyRows > 0 && (
                                <TableRow
                                    style={{
                                        height: (dense ? 33 : 53) * emptyRows,
                                    }}
                                >
                                    <TableCell colSpan={6}/>
                                </TableRow>
                            )}
                        </TableBody>
                    </Table>
                </TableContainer>
                <TablePagination
                    rowsPerPageOptions={[5, 10, 25]}
                    component="div"
                    count={rows.length}
                    rowsPerPage={rowsPerPage}
                    page={page}
                    onPageChange={handleChangePage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                />
            </Paper>
        </Box>
    );
}
