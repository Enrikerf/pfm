import * as React from "react";
import GenericTable from "../Components/MyTable/GenericTable";
import {Task} from "../Domain/Task";
import {TableRowData} from "../Components/MyTable/TableRowData";

function TasksTab() {
    const rows :TableRowData[] = [
        {
            "values": [
                {"name": "uuid", "value": "daa883d8-c9b0-4c49-9f9c-b87cd5603758"},
                {"name": "host", "value": "2.tcp.ngrok.io"},
                {"name": "port", "value": "14489"},
                {"name": "mode", "value": "Bidirectional"},
                {"name": "status", "value": "Done"},
                {"name": "executionMode", "value": "Manual"},
                {"name": "createdAt", "value": "2022-02-06 12:29:53"},
                {"name": "updatedAt", "value": "2022-02-06 12:29:53"},
            ]
        }
    ]
    return (
        <div>
            <h1>Tasks</h1>
            <GenericTable rows={rows}/>
        </div>

    );

}

export default TasksTab;