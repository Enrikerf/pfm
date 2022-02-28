import * as React from "react";
import {useEffect, useState} from "react";
import GenericTable from "../Components/MyTable/GenericTable";
import {TableRowData} from "../Components/MyTable/TableRowData";
import {useNavigate} from "react-router-dom";
import {TableData} from "../Components/MyTable/TableData";
import {TaskServiceClient} from "../protobuf/generated/task_grpc_web_pb";

function TasksList() {

    let navigate = useNavigate();
    const [lastId, setLastId] = useState(0)
    const [rows, setRows] = useState([] as TableRowData[])
    useEffect(() => {
        const messages = require('../protobuf/generated/task_pb');
        let listTaskRequest = new messages.ListTasksRequest()
        let metadata = {};
        let taskService = new TaskServiceClient("http://localhost:8080", null, null)
        taskService.listTasks(listTaskRequest, metadata, function (err, response) {
            if (err) {
                console.log(err);
            } else {
                let protoTasks = response.getTasksList()
                let newRows: TableRowData[] = []
                for (let i = 0; i < protoTasks.length; i++) {
                    newRows.push({
                        id: i + 1,
                        values: [
                            {name: "run", value: "icon"},
                            {name: "uuid", value: protoTasks[i].getUuid()},
                            {name: "host", value: protoTasks[i].getHost()},
                            {name: "port", value: protoTasks[i].getPort()},
                            {name: "mode", value: protoTasks[i].getMode()},
                            {name: "status", value: protoTasks[i].getStatus()},
                            {name: "executionMode", value: protoTasks[i].getExecutionmode()},
                            {name: "batches", value: "icon"},
                            {name: "results", value: "icon"},
                            {name: "createdAt", value: ""},
                            {name: "updatedAt", value: ""},
                        ]
                    })
                }
                setLastId(protoTasks.length)
                setRows(newRows)
            }
        })
        return () => {};
    }, [])

    async function handleGoTo(event: React.MouseEvent<unknown>, id: number, toGo: TableData) {
        let uuid = rows.find(row => row.id === id)?.values.find(value => value.name === "uuid")
        if (uuid) {
            navigate(uuid.value + "/" + toGo.name, {});
        }
    }

    return (
        <div>
            <h1>Tasks</h1>
            <GenericTable rows={rows} handleGoTo={handleGoTo}/>
        </div>

    );

}

export default TasksList;