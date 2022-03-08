import {ResultServiceClient} from "../protobuf/generated/result_grpc_web_pb";

const resultMessages = require('../protobuf/generated/result_pb')
const regexExp = /^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$/gi;

export class ResultService {
    client: ResultServiceClient
    messages
    newUuid: string

    constructor() {
        this.client = new ResultServiceClient("http://localhost:8080", null, null)
        this.messages = resultMessages
        this.newUuid = ""
    }

    getStream(batchUuid: string) {
        let streamResultsRequest = new resultMessages.StreamResultsRequest()
        streamResultsRequest.setBatchUuid(batchUuid)
        return this.client.streamResults(streamResultsRequest, {})
    }
}
