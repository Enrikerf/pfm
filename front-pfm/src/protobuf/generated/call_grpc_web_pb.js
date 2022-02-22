/**
 * @fileoverview gRPC-Web generated client stub for call
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.call = require('./call_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.call.CallServiceClient =
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
proto.call.CallServicePromiseClient =
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
 *   !proto.call.CallRequest,
 *   !proto.call.CallResponse>}
 */
const methodDescriptor_CallService_CallUnary = new grpc.web.MethodDescriptor(
  '/call.CallService/CallUnary',
  grpc.web.MethodType.UNARY,
  proto.call.CallRequest,
  proto.call.CallResponse,
  /**
   * @param {!proto.call.CallRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.call.CallResponse.deserializeBinary
);


/**
 * @param {!proto.call.CallRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.call.CallResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.call.CallResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.call.CallServiceClient.prototype.callUnary =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/call.CallService/CallUnary',
      request,
      metadata || {},
      methodDescriptor_CallService_CallUnary,
      callback);
};


/**
 * @param {!proto.call.CallRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.call.CallResponse>}
 *     Promise that resolves to the response
 */
proto.call.CallServicePromiseClient.prototype.callUnary =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/call.CallService/CallUnary',
      request,
      metadata || {},
      methodDescriptor_CallService_CallUnary);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.call.CallRequest,
 *   !proto.call.CallResponse>}
 */
const methodDescriptor_CallService_CallServerStream = new grpc.web.MethodDescriptor(
  '/call.CallService/CallServerStream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.call.CallRequest,
  proto.call.CallResponse,
  /**
   * @param {!proto.call.CallRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.call.CallResponse.deserializeBinary
);


/**
 * @param {!proto.call.CallRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.call.CallResponse>}
 *     The XHR Node Readable Stream
 */
proto.call.CallServiceClient.prototype.callServerStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/call.CallService/CallServerStream',
      request,
      metadata || {},
      methodDescriptor_CallService_CallServerStream);
};


/**
 * @param {!proto.call.CallRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.call.CallResponse>}
 *     The XHR Node Readable Stream
 */
proto.call.CallServicePromiseClient.prototype.callServerStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/call.CallService/CallServerStream',
      request,
      metadata || {},
      methodDescriptor_CallService_CallServerStream);
};


module.exports = proto.call;

