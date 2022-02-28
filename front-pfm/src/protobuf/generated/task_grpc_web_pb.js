/**
 * @fileoverview gRPC-Web generated client stub for task
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
proto.task = require('./task_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.task.TaskServiceClient =
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
proto.task.TaskServicePromiseClient =
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
 *   !proto.task.CreateTaskRequest,
 *   !proto.task.CreateTaskResponse>}
 */
const methodDescriptor_TaskService_CreateTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/CreateTask',
  grpc.web.MethodType.UNARY,
  proto.task.CreateTaskRequest,
  proto.task.CreateTaskResponse,
  /**
   * @param {!proto.task.CreateTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.CreateTaskResponse.deserializeBinary
);


/**
 * @param {!proto.task.CreateTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.task.CreateTaskResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.CreateTaskResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.createTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/CreateTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_CreateTask,
      callback);
};


/**
 * @param {!proto.task.CreateTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.CreateTaskResponse>}
 *     Promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.createTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/CreateTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_CreateTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.ReadTaskRequest,
 *   !proto.task.ReadTaskResponse>}
 */
const methodDescriptor_TaskService_ReadTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/ReadTask',
  grpc.web.MethodType.UNARY,
  proto.task.ReadTaskRequest,
  proto.task.ReadTaskResponse,
  /**
   * @param {!proto.task.ReadTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.ReadTaskResponse.deserializeBinary
);


/**
 * @param {!proto.task.ReadTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.task.ReadTaskResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.ReadTaskResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.readTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/ReadTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_ReadTask,
      callback);
};


/**
 * @param {!proto.task.ReadTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.ReadTaskResponse>}
 *     Promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.readTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/ReadTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_ReadTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.UpdateTaskRequest,
 *   !proto.task.UpdateTaskResponse>}
 */
const methodDescriptor_TaskService_UpdateTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/UpdateTask',
  grpc.web.MethodType.UNARY,
  proto.task.UpdateTaskRequest,
  proto.task.UpdateTaskResponse,
  /**
   * @param {!proto.task.UpdateTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.UpdateTaskResponse.deserializeBinary
);


/**
 * @param {!proto.task.UpdateTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.task.UpdateTaskResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.UpdateTaskResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.updateTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/UpdateTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_UpdateTask,
      callback);
};


/**
 * @param {!proto.task.UpdateTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.UpdateTaskResponse>}
 *     Promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.updateTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/UpdateTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_UpdateTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.DeleteTaskRequest,
 *   !proto.task.DeleteTaskResponse>}
 */
const methodDescriptor_TaskService_DeleteTask = new grpc.web.MethodDescriptor(
  '/task.TaskService/DeleteTask',
  grpc.web.MethodType.UNARY,
  proto.task.DeleteTaskRequest,
  proto.task.DeleteTaskResponse,
  /**
   * @param {!proto.task.DeleteTaskRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.DeleteTaskResponse.deserializeBinary
);


/**
 * @param {!proto.task.DeleteTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.task.DeleteTaskResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.DeleteTaskResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.deleteTask =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/DeleteTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_DeleteTask,
      callback);
};


/**
 * @param {!proto.task.DeleteTaskRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.DeleteTaskResponse>}
 *     Promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.deleteTask =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/DeleteTask',
      request,
      metadata || {},
      methodDescriptor_TaskService_DeleteTask);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.task.ListTasksRequest,
 *   !proto.task.ListTasksResponse>}
 */
const methodDescriptor_TaskService_ListTasks = new grpc.web.MethodDescriptor(
  '/task.TaskService/ListTasks',
  grpc.web.MethodType.UNARY,
  proto.task.ListTasksRequest,
  proto.task.ListTasksResponse,
  /**
   * @param {!proto.task.ListTasksRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.task.ListTasksResponse.deserializeBinary
);


/**
 * @param {!proto.task.ListTasksRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.task.ListTasksResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.task.ListTasksResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.task.TaskServiceClient.prototype.listTasks =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/task.TaskService/ListTasks',
      request,
      metadata || {},
      methodDescriptor_TaskService_ListTasks,
      callback);
};


/**
 * @param {!proto.task.ListTasksRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.task.ListTasksResponse>}
 *     Promise that resolves to the response
 */
proto.task.TaskServicePromiseClient.prototype.listTasks =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/task.TaskService/ListTasks',
      request,
      metadata || {},
      methodDescriptor_TaskService_ListTasks);
};


module.exports = proto.task;

