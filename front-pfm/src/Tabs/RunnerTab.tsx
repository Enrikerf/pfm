import * as React from "react";
import EnhancedTable from "../Components/Table";
import taskMessages from "../protobuf/generated/task_pb";
import resultMessages from "../protobuf/generated/result_pb";
import {ResultServiceClient} from "../protobuf/generated/result_grpc_web_pb";
import {TaskServiceClient} from "../protobuf/generated/task_grpc_web_pb";
import {useParams} from "react-router-dom";
import {useState} from "react";

export default function RunnerTab() {

    let {taskUuid} = useParams()
    const [batchUuid,setBatchUuid] = useState("")

    function getResults(): void {
        const resultMessages = require('../protobuf/generated/result_pb');
        const taskMessages = require('../protobuf/generated/task_pb');

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
    function run(){

    }
    return (
        <div>
            <button onClick={getResults}>Run</button>
            <p>1 create batch on use effect</p>
            <p>2 wait to press run</p>
            <p>3 on run execute getResultsStream</p>
        </div>
    );
}
