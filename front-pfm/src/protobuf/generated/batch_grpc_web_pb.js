/**
 * @fileoverview gRPC-Web generated client stub for batch
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js')
const proto = {};
proto.batch = require('./batch_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.batch.BatchServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.batch.BatchServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.batch.CreateBatchRequest,
 *   !proto.batch.CreateBatchResponse>}
 */
const methodDescriptor_BatchService_CreateBatch = new grpc.web.MethodDescriptor(
  '/batch.BatchService/CreateBatch',
  grpc.web.MethodType.UNARY,
  proto.batch.CreateBatchRequest,
  proto.batch.CreateBatchResponse,
  /**
   * @param {!proto.batch.CreateBatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.batch.CreateBatchResponse.deserializeBinary
);


/**
 * @param {!proto.batch.CreateBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.batch.CreateBatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.batch.CreateBatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.batch.BatchServiceClient.prototype.createBatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/batch.BatchService/CreateBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_CreateBatch,
      callback);
};


/**
 * @param {!proto.batch.CreateBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.batch.CreateBatchResponse>}
 *     Promise that resolves to the response
 */
proto.batch.BatchServicePromiseClient.prototype.createBatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/batch.BatchService/CreateBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_CreateBatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.batch.ReadBatchRequest,
 *   !proto.batch.ReadBatchResponse>}
 */
const methodDescriptor_BatchService_ReadBatch = new grpc.web.MethodDescriptor(
  '/batch.BatchService/ReadBatch',
  grpc.web.MethodType.UNARY,
  proto.batch.ReadBatchRequest,
  proto.batch.ReadBatchResponse,
  /**
   * @param {!proto.batch.ReadBatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.batch.ReadBatchResponse.deserializeBinary
);


/**
 * @param {!proto.batch.ReadBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.batch.ReadBatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.batch.ReadBatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.batch.BatchServiceClient.prototype.readBatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/batch.BatchService/ReadBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_ReadBatch,
      callback);
};


/**
 * @param {!proto.batch.ReadBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.batch.ReadBatchResponse>}
 *     Promise that resolves to the response
 */
proto.batch.BatchServicePromiseClient.prototype.readBatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/batch.BatchService/ReadBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_ReadBatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.batch.UpdateBatchRequest,
 *   !proto.batch.UpdateBatchResponse>}
 */
const methodDescriptor_BatchService_UpdateBatch = new grpc.web.MethodDescriptor(
  '/batch.BatchService/UpdateBatch',
  grpc.web.MethodType.UNARY,
  proto.batch.UpdateBatchRequest,
  proto.batch.UpdateBatchResponse,
  /**
   * @param {!proto.batch.UpdateBatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.batch.UpdateBatchResponse.deserializeBinary
);


/**
 * @param {!proto.batch.UpdateBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.batch.UpdateBatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.batch.UpdateBatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.batch.BatchServiceClient.prototype.updateBatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/batch.BatchService/UpdateBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_UpdateBatch,
      callback);
};


/**
 * @param {!proto.batch.UpdateBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.batch.UpdateBatchResponse>}
 *     Promise that resolves to the response
 */
proto.batch.BatchServicePromiseClient.prototype.updateBatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/batch.BatchService/UpdateBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_UpdateBatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.batch.DeleteBatchRequest,
 *   !proto.batch.DeleteBatchResponse>}
 */
const methodDescriptor_BatchService_DeleteBatch = new grpc.web.MethodDescriptor(
  '/batch.BatchService/DeleteBatch',
  grpc.web.MethodType.UNARY,
  proto.batch.DeleteBatchRequest,
  proto.batch.DeleteBatchResponse,
  /**
   * @param {!proto.batch.DeleteBatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.batch.DeleteBatchResponse.deserializeBinary
);


/**
 * @param {!proto.batch.DeleteBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.batch.DeleteBatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.batch.DeleteBatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.batch.BatchServiceClient.prototype.deleteBatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/batch.BatchService/DeleteBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_DeleteBatch,
      callback);
};


/**
 * @param {!proto.batch.DeleteBatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.batch.DeleteBatchResponse>}
 *     Promise that resolves to the response
 */
proto.batch.BatchServicePromiseClient.prototype.deleteBatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/batch.BatchService/DeleteBatch',
      request,
      metadata || {},
      methodDescriptor_BatchService_DeleteBatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.batch.ListBatchesRequest,
 *   !proto.batch.ListBatchesResponse>}
 */
const methodDescriptor_BatchService_ListBatches = new grpc.web.MethodDescriptor(
  '/batch.BatchService/ListBatches',
  grpc.web.MethodType.UNARY,
  proto.batch.ListBatchesRequest,
  proto.batch.ListBatchesResponse,
  /**
   * @param {!proto.batch.ListBatchesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.batch.ListBatchesResponse.deserializeBinary
);


/**
 * @param {!proto.batch.ListBatchesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.batch.ListBatchesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.batch.ListBatchesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.batch.BatchServiceClient.prototype.listBatches =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/batch.BatchService/ListBatches',
      request,
      metadata || {},
      methodDescriptor_BatchService_ListBatches,
      callback);
};


/**
 * @param {!proto.batch.ListBatchesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.batch.ListBatchesResponse>}
 *     Promise that resolves to the response
 */
proto.batch.BatchServicePromiseClient.prototype.listBatches =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/batch.BatchService/ListBatches',
      request,
      metadata || {},
      methodDescriptor_BatchService_ListBatches);
};


module.exports = proto.batch;

