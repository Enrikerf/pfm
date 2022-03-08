import * as React from "react";
import {useState} from "react";
import {ResultServiceClient} from "../../protobuf/generated/result_grpc_web_pb";
import {TaskServiceClient} from "../../protobuf/generated/task_grpc_web_pb";
import {BatchServiceClient} from "../../protobuf/generated/batch_grpc_web_pb";
import {useParams} from "react-router-dom";
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from 'recharts';
import "./RunnerTab.scss"
import Button from "@mui/material/Button";


interface GraphPoint {
    name: string,
    point: number
}


export default function RunnerTab() {

    let {uuid} = useParams()
    const [batchUuid, setBatchUuid] = useState("")
    const [rows, setRows] = useState([] as GraphPoint[])
    const [yMax, setYMax] = useState(0)
    const [yMin, setYMin] = useState(0)

    async function run(): Promise<void> {
        const batchMessages = require('../../protobuf/generated/batch_pb')
        let createBatchRequest = new batchMessages.CreateBatchRequest()
        let createBatchParams = new batchMessages.CreateBatchParams()
        const resultMessages = require('../../protobuf/generated/result_pb');
        let streamResultsRequest = new resultMessages.StreamResultsRequest()
        let resultServiceClient = new ResultServiceClient("http://localhost:8080", null, null)
        createBatchParams.setTaskUuid(uuid)
        createBatchRequest.setBatchparams(createBatchParams)
        let batchServiceClient = new BatchServiceClient("http://localhost:8080", null, null)
        let batchUuid: string
        batchServiceClient.createBatch(createBatchRequest, {}, function (err, response) {
            if (err) {
                console.log(err);
            } else {
                setBatchUuid(response.getBatch().getUuid())
                streamResultsRequest.setBatchUuid(response.getBatch().getUuid())
                let stream = resultServiceClient.streamResults(streamResultsRequest, {})

                let newYmax: number = 0
                let newYmin: number = 0
                let newRows: GraphPoint[] = []
                let slice: GraphPoint[]
                // @ts-ignore
                stream.on('data', function (response) {
                    let protoResults = response.getResultsList()
                    for (let i = 0; i < protoResults.length; i++) {

                        if (~~(protoResults[i].getContent()) > yMax) {
                            newYmax = ~~(protoResults[i].getContent())
                        }
                        if (protoResults[i].getContent() < yMin) {
                            newYmin = protoResults[i].getContent()
                        }
                        let date = new Date(protoResults[i].getCreatedAt())


                        newRows.push({
                            name: date?.getHours() + ":" + date?.getMinutes() + ":" + date?.getSeconds(),
                            point: protoResults[i].getContent()
                        })
                    }
                    if (newYmax > yMax) {
                        setYMax(newYmax)
                    }
                    if (newYmin > yMin) {
                        setYMin(newYmin)
                    }
                    if (rows.length + newRows.length > 1000) {
                        console.log("sliced")
                        slice = rows.slice(0, 1000 - (rows.length + newRows.length))
                    } else {
                        console.log("not ")
                        slice = rows
                    }
                    let newTotalRows = [...slice, ...newRows]
                    console.log(newTotalRows)
                    setRows(newTotalRows)
                });
                // @ts-ignore
                stream.on('status', function (status) {
                    console.log("status");
                });
                // @ts-ignore
                stream.on('end', function (end) {
                    console.log("end");
                    stream.cancel()
                });
            }
        })
    }

    function clear() {
        setRows([])
        setYMax(0)
        setYMin(0)
    }

    async function stop() {
        const taskMessages = require('../../protobuf/generated/task_pb');

        let updateTaskRequest = new taskMessages.UpdateTaskRequest()
        let updateTaskParams = new taskMessages.EditableTaskParams()
        updateTaskRequest.setTaskUuid(uuid)
        updateTaskParams.setStatus(3)
        updateTaskRequest.setParams(updateTaskParams)
        let taskServiceClient = new TaskServiceClient("http://localhost:8080", null, null)

        console.log("stopping")
        taskServiceClient.updateTask(updateTaskRequest, {}, function (err, response) {
            if (err) {
                console.log(err);
            } else {
                console.log("finish run")
            }
        })

    }

    return (
        <div className="runner-container">
            <div className="runner-button-container">
                <Button
                    color={"success"}
                    size={"large"}
                    variant={"contained"}
                    className="runner-button-container-button"
                    onClick={run}
                >Run</Button>
                <Button
                    color={"error"}
                    size={"large"}
                    variant={"contained"}
                    className="runner-button-container-button"
                    onClick={stop}
                >Stop</Button>
                <Button
                    color={"warning"}
                    size={"large"}
                    variant={"contained"}
                    className="runner-button-container-button"
                    onClick={clear}
                >Clear</Button>
            </div>
            <h3>task: {uuid}</h3>
            <h3>batch: {batchUuid}</h3>
            <div className={"runner-dashboard"}>
                <div className={"runner-dashboard-graph-container"}>
                    <ResponsiveContainer width='60%' minWidth={"60rem"} aspect={3.0}>
                        <LineChart
                            data={rows}
                            margin={{
                                top: 5,
                                right: 30,
                                left: 20,
                                bottom: 5,
                            }}
                        >
                            <CartesianGrid strokeDasharray="3 3"/>
                            <XAxis dataKey="name"/>
                            <YAxis domain={[yMin, yMax]} scale={"linear"}/>
                            <Tooltip/>
                            <Legend/>
                            <Line type="monotone" dataKey="point" stroke="#8884d8" dot={false}/>
                        </LineChart>
                    </ResponsiveContainer>
                </div>
            </div>
        </div>
    );
}
