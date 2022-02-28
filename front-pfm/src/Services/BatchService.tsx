import {BatchServiceClient} from "../protobuf/generated/batch_grpc_web_pb";

export class BatchService {
    batchServiceClient: BatchServiceClient

    constructor() {
        this.batchServiceClient = new BatchServiceClient("http://localhost:8080", null, null)
    }

    
}
