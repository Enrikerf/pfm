import * as React from "react";
import {useState} from "react";
import {ResultServiceClient} from "../../protobuf/generated/result_grpc_web_pb";
import {TaskServiceClient} from "../../protobuf/generated/task_grpc_web_pb";
import {useParams} from "react-router-dom";
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from 'recharts';
import GenericTable from "../../Components/MyTable/GenericTable";
import {TableRowData} from "../../Components/MyTable/TableRowData";
import "./RunnerTab.scss"
import Button from "@mui/material/Button";

export default function RunnerTab() {

    let {uuid} = useParams()
    const [batchUuid, setBatchUuid] = useState("")
    const [lastId, setLastId] = useState(0)
    const [rows, setRows] = useState([] as TableRowData[])

    function getResults(): void {
        const resultMessages = require('../../protobuf/generated/result_pb');
        const taskMessages = require('../../protobuf/generated/task_pb');

        let updateTaskRequest = new taskMessages.UpdateTaskRequest()
        let updateTaskParams = new taskMessages.EditableTaskParams()
        updateTaskRequest.setTaskUuid("daa883d8-c9b0-4c49-9f9c-b87cd5603758")
        updateTaskParams.setStatus(3)
        updateTaskRequest.setParams(updateTaskParams)
        let streamResultsRequest = new resultMessages.StreamResultsRequest()
        streamResultsRequest.setBatchUuid("1142dfa2-058d-4e88-8a80-a308a2bdae6f")

        let resultServiceClient = new ResultServiceClient("http://localhost:8080", null, null)
        let taskServiceClient = new TaskServiceClient("http://localhost:8080", null, null)


        let stream = resultServiceClient.streamResults(streamResultsRequest, {})
        // @ts-ignore
        let arrayResult = []
        // @ts-ignore
        stream.on('data', function (response) {
            let responseArray = response.getResultsList()
            responseArray.map(
                // @ts-ignore
                result => arrayResult.push(result.getContent())
            )
            // @ts-ignore
            console.log(arrayResult)
            if (responseArray[responseArray.length - 1].getContent() > 359) {
                console.log("stopping")
                taskServiceClient.updateTask(updateTaskRequest, {}, function (err, response) {
                    if (err) {
                        console.log(err);
                    } else {
                        console.log(response);
                    }
                })
            }
        });
        // @ts-ignore
        stream.on('status', function (status) {
            console.log("status");
            console.log(status.code);
            console.log(status.details);
            console.log(status.metadata);
        });
        // @ts-ignore
        stream.on('end', function (end) {
            console.log("end");
            stream.cancel()
        });

// to close the stream

    }


    const data = [
        {
            name: 'Page A',
            uv: 4000,
            pv: 2400,
            amt: 2400,
        },
        {
            name: 'Page B',
            uv: 3000,
            pv: 1398,
            amt: 2210,
        },
        {
            name: 'Page C',
            uv: 2000,
            pv: 9800,
            amt: 2290,
        },
        {
            name: 'Page D',
            uv: 2780,
            pv: 3908,
            amt: 2000,
        },
        {
            name: 'Page E',
            uv: 1890,
            pv: 4800,
            amt: 2181,
        },
        {
            name: 'Page F',
            uv: 2390,
            pv: 3800,
            amt: 2500,
        },
        {
            name: 'Page G',
            uv: 3490,
            pv: 4300,
            amt: 2100,
        },
    ];

    function handleGoTo() {

    }

    return (
        <div className="runner-container">
            <div className="runner-button-container">
                <Button
                    color={"success"}
                    size={"large"}
                    variant={"contained"}
                    className="runner-button-container-button"
                    onClick={getResults}
                >Run</Button>
                <Button
                    color={"error"}
                    size={"large"}
                    variant={"contained"}
                    className="runner-button-container-button"
                    onClick={getResults}
                >Stop</Button>
            </div>
            <div className={"runner-dashboard"}>
                <div className={"runner-dashboard-data-container"}>
                    <GenericTable rows={rows} handleGoTo={handleGoTo}/>
                </div>
                <div className={"runner-dashboard-graph-container"}>
                    <ResponsiveContainer width='60%' minWidth={"60rem"} aspect={3.0}>
                        <LineChart
                            data={data}
                            margin={{
                                top: 5,
                                right: 30,
                                left: 20,
                                bottom: 5,
                            }}
                        >
                            <CartesianGrid strokeDasharray="3 3"/>
                            <XAxis dataKey="name"/>
                            <YAxis/>
                            <Tooltip/>
                            <Legend/>
                            <Line type="monotone" dataKey="pv" stroke="#8884d8" activeDot={{r: 8}}/>
                            <Line type="monotone" dataKey="uv" stroke="#82ca9d"/>
                        </LineChart>
                    </ResponsiveContainer>
                </div>

            </div>
            {/*            <p>1 create batch on use effect</p>
            <p>2 wait to press run</p>
            <p>3 on run execute getResultsStream</p>*/}


        </div>
    );
}
