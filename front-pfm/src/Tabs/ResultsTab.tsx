import * as React from "react";
import {useEffect, useState} from "react";
import {useLocation, useParams} from "react-router-dom";
import GenericTable from "../Components/MyTable/GenericTable";
import {TableRowData} from "../Components/MyTable/TableRowData";
import {ResultServiceClient} from "../protobuf/generated/result_grpc_web_pb";
import {TableData} from "../Components/MyTable/TableData";

export default function ResultsTab() {
    let location = useLocation();
    let {uuid} = useParams()
    const [lastId, setLastId] = useState(0)
    const [rows, setRows] = useState([] as TableRowData[])
    useEffect(() => {
        const messages = require('../protobuf/generated/result_pb');
        let listTaskRequest = new messages.ListResultRequest()
        let metadata = {};
        let taskService = new ResultServiceClient("http://localhost:8080", null, null)
        taskService.listResult(listTaskRequest, metadata, function (err, response) {
            if (err) {
                console.log(err);
            } else {
                let protoResults = response.getResultsList()
                let newRows: TableRowData[] = []
                for (let i = 0; i < protoResults.length; i++) {
                    newRows.push({
                        id: i + 1,
                        values: [
                            {name: "uuid", value: protoResults[i].getUuid()},
                            {name: "content", value: protoResults[i].getContent()},
                            {name: "batch_uuid", value: protoResults[i].getBatchUuid()},
                            {name: "created_at", value: protoResults[i].getCreatedAt()},
                            {name: "updated_at", value: protoResults[i].getUpdatedAt()},
                        ]
                    })
                }
                setLastId(protoResults.length)
                setRows(newRows)
            }
        })
        return () => {
        };
    }, [])

    async function handleGoTo(event: React.MouseEvent<unknown>, id: number, toGo: TableData) {
    }

    return (
        <div>
            <h1>from task: {uuid}</h1>
            <h1>Results</h1>
            <GenericTable rows={rows} handleGoTo={handleGoTo}/>
        </div>
    );

}
