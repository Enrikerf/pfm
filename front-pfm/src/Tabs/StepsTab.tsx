import * as React from "react";
import {useEffect, useState} from "react";
import {useLocation, useParams} from "react-router-dom";
import GenericTable from "../Components/MyTable/GenericTable";
import {TableRowData} from "../Components/MyTable/TableRowData";
import {StepServiceClient} from "../protobuf/generated/step_grpc_web_pb";
import {TableData} from "../Components/MyTable/TableData";

export default function StepsTab() {
    let location = useLocation();
    let {uuid} = useParams()
    const [lastId, setLastId] = useState(0)
    const [rows, setRows] = useState([] as TableRowData[])
    useEffect(() => {
        const messages = require('../protobuf/generated/step_pb');
        let listStepsRequest = new messages.ListStepsRequest()
        let metadata = {};
        let stepService = new StepServiceClient("http://localhost:8080", null, null)
        stepService.listSteps(listStepsRequest, metadata, function (err, response) {
            if (err) {
                console.log(err);
            } else {
                let protoSteps = response.getStepsList()
                let newRows: TableRowData[] = []
                for (let i = 0; i < protoSteps.length; i++) {
                    newRows.push({
                        id: i + 1,
                        values: [
                            {name: "uuid", value: protoSteps[i].getUuid()},
                            {name: "task_uuid", value: protoSteps[i].getTaskUuid()},
                            {name: "sentence", value: protoSteps[i].getSentence()},
                            {name: "created_at", value: ""},
                            {name: "updated_at", value: ""},
                        ]
                    })
                }
                setLastId(protoSteps.length)
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
            <h1>Steps</h1>
            <GenericTable rows={rows} handleGoTo={handleGoTo}/>
        </div>
    );

}
