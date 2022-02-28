import * as React from "react";
import {useEffect, useState} from "react";
import {BatchServiceClient} from "../protobuf/generated/batch_grpc_web_pb";
import GenericTable from "../Components/MyTable/GenericTable";
import {TableRowData} from "../Components/MyTable/TableRowData";
import {TableData} from "../Components/MyTable/TableData";
import {useNavigate} from "react-router-dom";


function BatchesTab() {
    let navigate = useNavigate();

    const [lastId, setLastId] = useState(0)
    const [rows, setRows] = useState([] as TableRowData[])
    useEffect(() => {
        const messages = require('../protobuf/generated/batch_pb');
        let listTaskRequest = new messages.ListBatchesRequest()
        let metadata = {};
        let taskService = new BatchServiceClient("http://localhost:8080", null, null)
        taskService.listBatches(listTaskRequest, metadata, function (err, response) {
            if (err) {
                console.log(err);
            } else {
                let protoTasks = response.getBatchesList()
                let newRows: TableRowData[] = []
                for (let i = 0; i < protoTasks.length; i++) {
                    newRows.push({
                        id: i + 1,
                        values: [
                            {name: "uuid", value: protoTasks[i].getUuid()},
                            {name: "task_uuid", value: protoTasks[i].getTaskUuid()},
                            {name: "created_at", value: protoTasks[i].getCreatedAt()},
                            {name: "updated_at", value: protoTasks[i].getUpdatedAt()},
                        ]
                    })
                }
                setLastId(protoTasks.length)
                setRows(newRows)
            }
        })
        return () => {
        };
    }, [])


    async function handleGoTo(event: React.MouseEvent<unknown>, id: number, toGo: TableData) {
        if (toGo.value === "icon") {
            let uuid = rows.find(row => row.id === id)?.values.find(value => value.name === "uuid")
            if (uuid) {
                navigate(uuid.value + "/" + toGo.name, {});
            }
        } else {
            switch (toGo.name) {
                case "taskUuid":
                    navigate("/tasks", {replace: true});
                    break;
            }
        }

    }

    return (
        <div>
            <h1>Batches</h1>
            <GenericTable rows={rows} handleGoTo={handleGoTo}/>
        </div>

    );

}

export default BatchesTab