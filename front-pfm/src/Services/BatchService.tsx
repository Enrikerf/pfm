import {BatchServiceClient} from "../protobuf/generated/batch_grpc_web_pb";

const batchMessages = require('../protobuf/generated/batch_pb')
const regexExp = /^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$/gi;

export class BatchService {
    client: BatchServiceClient
    messages
    newUuid:string


    constructor() {
        this.client = new BatchServiceClient("http://localhost:8080", null, null)
        this.messages = batchMessages
        this.newUuid = ""
    }

    async create(taskUuid: string): Promise<string> {
        if (regexExp.test(taskUuid)) {
            let createBatchRequest = new this.messages.CreateBatchRequest()
            let createBatchParams = new this.messages.CreateBatchParams()
            createBatchParams.setTaskUuid(taskUuid)
            createBatchRequest.setBatchparams(createBatchParams)

            return new Promise((resolve,reject)=>{
                this.client.createBatch(createBatchRequest, {}, function (err, response) {
                    if (err) {
                        console.log(err);
                        reject(err)
                    } else {
                        resolve(response.getBatch().getUuid())
                    }
                })
            })
        } else {
            console.log("error")
            return Promise.resolve("")
        }
    }

}
