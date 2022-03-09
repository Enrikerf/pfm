import * as React from "react";
import {useEffect, useState} from "react";
import {useLocation, useParams} from "react-router-dom";
import GenericTable from "../Components/MyTable/GenericTable";
import {TableRowData} from "../Components/MyTable/TableRowData";
import {ResultServiceClient} from "../protobuf/generated/result_grpc_web_pb";
import {TableData} from "../Components/MyTable/TableData";
import messages from "../protobuf/generated/batch_pb";

export default function ResultsTab() {
    let location = useLocation();
    let {uuid} = useParams()
    const [lastId, setLastId] = useState(0)
    const [rows, setRows] = useState([] as TableRowData[])
    useEffect(() => {
        const messages = require('../protobuf/generated/result_pb');
        let listTaskRequest = new messages.ListResultRequest()
        let google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
        if (uuid) {
            let uuidStringValue = new google_protobuf_wrappers_pb.StringValue([uuid])
            let filters = new messages.Filters()
            filters.setBatchUuid(uuidStringValue)
            listTaskRequest.setFilters(filters)
        }
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
        <div style={{flex: " 1 1 auto", display: "flex", flexDirection: "column"}}>
            <h1>Results</h1>
            <GenericTable rows={rows} handleGoTo={handleGoTo}/>
        </div>
    );

}
