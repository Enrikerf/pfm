/**
 * @fileoverview gRPC-Web generated client stub for step
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
proto.step = require('./step_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.step.StepServiceClient =
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
proto.step.StepServicePromiseClient =
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
 *   !proto.step.CreateStepRequest,
 *   !proto.step.CreateStepResponse>}
 */
const methodDescriptor_StepService_CreateStep = new grpc.web.MethodDescriptor(
  '/step.StepService/CreateStep',
  grpc.web.MethodType.UNARY,
  proto.step.CreateStepRequest,
  proto.step.CreateStepResponse,
  /**
   * @param {!proto.step.CreateStepRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.step.CreateStepResponse.deserializeBinary
);


/**
 * @param {!proto.step.CreateStepRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.step.CreateStepResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.step.CreateStepResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.step.StepServiceClient.prototype.createStep =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/step.StepService/CreateStep',
      request,
      metadata || {},
      methodDescriptor_StepService_CreateStep,
      callback);
};


/**
 * @param {!proto.step.CreateStepRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.step.CreateStepResponse>}
 *     Promise that resolves to the response
 */
proto.step.StepServicePromiseClient.prototype.createStep =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/step.StepService/CreateStep',
      request,
      metadata || {},
      methodDescriptor_StepService_CreateStep);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.step.ReadStepRequest,
 *   !proto.step.ReadStepResponse>}
 */
const methodDescriptor_StepService_ReadStep = new grpc.web.MethodDescriptor(
  '/step.StepService/ReadStep',
  grpc.web.MethodType.UNARY,
  proto.step.ReadStepRequest,
  proto.step.ReadStepResponse,
  /**
   * @param {!proto.step.ReadStepRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.step.ReadStepResponse.deserializeBinary
);


/**
 * @param {!proto.step.ReadStepRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.step.ReadStepResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.step.ReadStepResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.step.StepServiceClient.prototype.readStep =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/step.StepService/ReadStep',
      request,
      metadata || {},
      methodDescriptor_StepService_ReadStep,
      callback);
};


/**
 * @param {!proto.step.ReadStepRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.step.ReadStepResponse>}
 *     Promise that resolves to the response
 */
proto.step.StepServicePromiseClient.prototype.readStep =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/step.StepService/ReadStep',
      request,
      metadata || {},
      methodDescriptor_StepService_ReadStep);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.step.UpdateStepRequest,
 *   !proto.step.UpdateStepResponse>}
 */
const methodDescriptor_StepService_UpdateStep = new grpc.web.MethodDescriptor(
  '/step.StepService/UpdateStep',
  grpc.web.MethodType.UNARY,
  proto.step.UpdateStepRequest,
  proto.step.UpdateStepResponse,
  /**
   * @param {!proto.step.UpdateStepRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.step.UpdateStepResponse.deserializeBinary
);


/**
 * @param {!proto.step.UpdateStepRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.step.UpdateStepResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.step.UpdateStepResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.step.StepServiceClient.prototype.updateStep =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/step.StepService/UpdateStep',
      request,
      metadata || {},
      methodDescriptor_StepService_UpdateStep,
      callback);
};


/**
 * @param {!proto.step.UpdateStepRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.step.UpdateStepResponse>}
 *     Promise that resolves to the response
 */
proto.step.StepServicePromiseClient.prototype.updateStep =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/step.StepService/UpdateStep',
      request,
      metadata || {},
      methodDescriptor_StepService_UpdateStep);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.step.DeleteStepRequest,
 *   !proto.step.DeleteStepResponse>}
 */
const methodDescriptor_StepService_DeleteStep = new grpc.web.MethodDescriptor(
  '/step.StepService/DeleteStep',
  grpc.web.MethodType.UNARY,
  proto.step.DeleteStepRequest,
  proto.step.DeleteStepResponse,
  /**
   * @param {!proto.step.DeleteStepRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.step.DeleteStepResponse.deserializeBinary
);


/**
 * @param {!proto.step.DeleteStepRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.step.DeleteStepResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.step.DeleteStepResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.step.StepServiceClient.prototype.deleteStep =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/step.StepService/DeleteStep',
      request,
      metadata || {},
      methodDescriptor_StepService_DeleteStep,
      callback);
};


/**
 * @param {!proto.step.DeleteStepRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.step.DeleteStepResponse>}
 *     Promise that resolves to the response
 */
proto.step.StepServicePromiseClient.prototype.deleteStep =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/step.StepService/DeleteStep',
      request,
      metadata || {},
      methodDescriptor_StepService_DeleteStep);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.step.ListStepsRequest,
 *   !proto.step.ListStepsResponse>}
 */
const methodDescriptor_StepService_ListSteps = new grpc.web.MethodDescriptor(
  '/step.StepService/ListSteps',
  grpc.web.MethodType.UNARY,
  proto.step.ListStepsRequest,
  proto.step.ListStepsResponse,
  /**
   * @param {!proto.step.ListStepsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.step.ListStepsResponse.deserializeBinary
);


/**
 * @param {!proto.step.ListStepsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.step.ListStepsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.step.ListStepsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.step.StepServiceClient.prototype.listSteps =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/step.StepService/ListSteps',
      request,
      metadata || {},
      methodDescriptor_StepService_ListSteps,
      callback);
};


/**
 * @param {!proto.step.ListStepsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.step.ListStepsResponse>}
 *     Promise that resolves to the response
 */
proto.step.StepServicePromiseClient.prototype.listSteps =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/step.StepService/ListSteps',
      request,
      metadata || {},
      methodDescriptor_StepService_ListSteps);
};


module.exports = proto.step;

