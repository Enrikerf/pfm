/**
 * @fileoverview gRPC-Web generated client stub for result
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
proto.result = require('./result_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.result.ResultServiceClient =
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
proto.result.ResultServicePromiseClient =
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
 *   !proto.result.CreateResultRequest,
 *   !proto.result.CreateResultResponse>}
 */
const methodDescriptor_ResultService_CreateResult = new grpc.web.MethodDescriptor(
  '/result.ResultService/CreateResult',
  grpc.web.MethodType.UNARY,
  proto.result.CreateResultRequest,
  proto.result.CreateResultResponse,
  /**
   * @param {!proto.result.CreateResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.result.CreateResultResponse.deserializeBinary
);


/**
 * @param {!proto.result.CreateResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.result.CreateResultResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.result.CreateResultResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.result.ResultServiceClient.prototype.createResult =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/result.ResultService/CreateResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_CreateResult,
      callback);
};


/**
 * @param {!proto.result.CreateResultRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.result.CreateResultResponse>}
 *     Promise that resolves to the response
 */
proto.result.ResultServicePromiseClient.prototype.createResult =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/result.ResultService/CreateResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_CreateResult);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.result.ReadResultRequest,
 *   !proto.result.ReadResultResponse>}
 */
const methodDescriptor_ResultService_ReadResult = new grpc.web.MethodDescriptor(
  '/result.ResultService/ReadResult',
  grpc.web.MethodType.UNARY,
  proto.result.ReadResultRequest,
  proto.result.ReadResultResponse,
  /**
   * @param {!proto.result.ReadResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.result.ReadResultResponse.deserializeBinary
);


/**
 * @param {!proto.result.ReadResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.result.ReadResultResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.result.ReadResultResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.result.ResultServiceClient.prototype.readResult =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/result.ResultService/ReadResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_ReadResult,
      callback);
};


/**
 * @param {!proto.result.ReadResultRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.result.ReadResultResponse>}
 *     Promise that resolves to the response
 */
proto.result.ResultServicePromiseClient.prototype.readResult =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/result.ResultService/ReadResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_ReadResult);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.result.UpdateResultRequest,
 *   !proto.result.UpdateResultResponse>}
 */
const methodDescriptor_ResultService_UpdateResult = new grpc.web.MethodDescriptor(
  '/result.ResultService/UpdateResult',
  grpc.web.MethodType.UNARY,
  proto.result.UpdateResultRequest,
  proto.result.UpdateResultResponse,
  /**
   * @param {!proto.result.UpdateResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.result.UpdateResultResponse.deserializeBinary
);


/**
 * @param {!proto.result.UpdateResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.result.UpdateResultResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.result.UpdateResultResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.result.ResultServiceClient.prototype.updateResult =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/result.ResultService/UpdateResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_UpdateResult,
      callback);
};


/**
 * @param {!proto.result.UpdateResultRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.result.UpdateResultResponse>}
 *     Promise that resolves to the response
 */
proto.result.ResultServicePromiseClient.prototype.updateResult =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/result.ResultService/UpdateResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_UpdateResult);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.result.DeleteResultRequest,
 *   !proto.result.DeleteResultResponse>}
 */
const methodDescriptor_ResultService_DeleteResult = new grpc.web.MethodDescriptor(
  '/result.ResultService/DeleteResult',
  grpc.web.MethodType.UNARY,
  proto.result.DeleteResultRequest,
  proto.result.DeleteResultResponse,
  /**
   * @param {!proto.result.DeleteResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.result.DeleteResultResponse.deserializeBinary
);


/**
 * @param {!proto.result.DeleteResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.result.DeleteResultResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.result.DeleteResultResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.result.ResultServiceClient.prototype.deleteResult =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/result.ResultService/DeleteResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_DeleteResult,
      callback);
};


/**
 * @param {!proto.result.DeleteResultRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.result.DeleteResultResponse>}
 *     Promise that resolves to the response
 */
proto.result.ResultServicePromiseClient.prototype.deleteResult =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/result.ResultService/DeleteResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_DeleteResult);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.result.ListResultRequest,
 *   !proto.result.ListResultResponse>}
 */
const methodDescriptor_ResultService_ListResult = new grpc.web.MethodDescriptor(
  '/result.ResultService/ListResult',
  grpc.web.MethodType.UNARY,
  proto.result.ListResultRequest,
  proto.result.ListResultResponse,
  /**
   * @param {!proto.result.ListResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.result.ListResultResponse.deserializeBinary
);


/**
 * @param {!proto.result.ListResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.result.ListResultResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.result.ListResultResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.result.ResultServiceClient.prototype.listResult =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/result.ResultService/ListResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_ListResult,
      callback);
};


/**
 * @param {!proto.result.ListResultRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.result.ListResultResponse>}
 *     Promise that resolves to the response
 */
proto.result.ResultServicePromiseClient.prototype.listResult =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/result.ResultService/ListResult',
      request,
      metadata || {},
      methodDescriptor_ResultService_ListResult);
};


module.exports = proto.result;

