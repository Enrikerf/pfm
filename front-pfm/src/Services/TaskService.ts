import {TaskServiceClient} from "../protobuf/generated/task_grpc_web_pb";

const taskMessages = require('../protobuf/generated/task_pb')
const regexExp = /^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$/gi;

export class TaskService {
    client: TaskServiceClient
    messages


    constructor() {
        this.client = new TaskServiceClient("http://2ab0-77-225-241-204.ngrok.io", null, null)
        this.messages = taskMessages
    }


    stop(uuid: string) {
        if (regexExp.test(uuid)) {
            let updateTaskRequest = new taskMessages.UpdateTaskRequest()
            let updateTaskParams = new taskMessages.EditableTaskParams()
            updateTaskRequest.setTaskUuid(uuid)
            updateTaskParams.setStatus(3)
            updateTaskRequest.setParams(updateTaskParams)
            this.client.updateTask(updateTaskRequest, {}, function (err, response) {
                if (err) {
                    console.log(err);
                } else {
                    console.log("stopped: " + uuid)
                }
            })
        } else {
            console.log("invalid uuid to stop")
        }

    }

}
