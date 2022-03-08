import * as React from "react";
import {useState} from "react";
import {useParams} from "react-router-dom";
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from 'recharts';
import "./RunnerTab.scss"
import Button from "@mui/material/Button";
import {TaskService} from "../../Services/TaskService";
import {BatchService} from "../../Services/BatchService";
import {ResultService} from "../../Services/ResultService";

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
    const taskService = new TaskService()
    const batchService = new BatchService()
    const resultService = new ResultService()

    async function newRun() {
        if (uuid !== undefined) {
            let newBatchUuid = await batchService.create(uuid)
            setBatchUuid(newBatchUuid)
            let stream = resultService.getStream(newBatchUuid)
            let newYmax: number = 0
            let newYmin: number = 0
            let newRows: GraphPoint[] = []
            let slice: GraphPoint[]
            stream.on('data', function (response: { getResultsList: () => any; }) {
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
                    slice = rows.slice(0, 1000 - (rows.length + newRows.length))
                } else {
                    slice = rows
                }
                let newTotalRows = [...slice, ...newRows]
                setRows(newTotalRows)
            });

            stream.on('end', function () {
                console.log("end");
                stream.cancel()
            });
        }
    }

    function clear() {
        setRows([])
        setYMax(0)
        setYMin(0)
    }

    async function stop() {
        if (uuid !== undefined) {
            taskService.stop(uuid)
        }
    }

    return (
        <div className="runner-container">
            <div className="runner-button-container">
                <Button
                    color={"success"}
                    size={"large"}
                    variant={"contained"}
                    className="runner-button-container-button"
                    onClick={newRun}
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
