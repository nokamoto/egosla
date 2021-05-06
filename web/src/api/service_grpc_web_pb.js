/**
 * @fileoverview gRPC-Web generated client stub for nokamoto.github.com.egosla.api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')
const proto = {};
proto.nokamoto = {};
proto.nokamoto.github = {};
proto.nokamoto.github.com = {};
proto.nokamoto.github.com.egosla = {};
proto.nokamoto.github.com.egosla.api = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.nokamoto.github.com.egosla.api.WatcherServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.nokamoto.github.com.egosla.api.WatcherServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.nokamoto.github.com.egosla.api.CreateWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.Watcher>}
 */
const methodDescriptor_WatcherService_CreateWatcher = new grpc.web.MethodDescriptor(
  '/nokamoto.github.com.egosla.api.WatcherService/CreateWatcher',
  grpc.web.MethodType.UNARY,
  proto.nokamoto.github.com.egosla.api.CreateWatcherRequest,
  proto.nokamoto.github.com.egosla.api.Watcher,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.CreateWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.Watcher.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.nokamoto.github.com.egosla.api.CreateWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.Watcher>}
 */
const methodInfo_WatcherService_CreateWatcher = new grpc.web.AbstractClientBase.MethodInfo(
  proto.nokamoto.github.com.egosla.api.Watcher,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.CreateWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.Watcher.deserializeBinary
);


/**
 * @param {!proto.nokamoto.github.com.egosla.api.CreateWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.nokamoto.github.com.egosla.api.Watcher)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.nokamoto.github.com.egosla.api.Watcher>|undefined}
 *     The XHR Node Readable Stream
 */
proto.nokamoto.github.com.egosla.api.WatcherServiceClient.prototype.createWatcher =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/CreateWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_CreateWatcher,
      callback);
};


/**
 * @param {!proto.nokamoto.github.com.egosla.api.CreateWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.nokamoto.github.com.egosla.api.Watcher>}
 *     Promise that resolves to the response
 */
proto.nokamoto.github.com.egosla.api.WatcherServicePromiseClient.prototype.createWatcher =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/CreateWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_CreateWatcher);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_WatcherService_DeleteWatcher = new grpc.web.MethodDescriptor(
  '/nokamoto.github.com.egosla.api.WatcherService/DeleteWatcher',
  grpc.web.MethodType.UNARY,
  proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_WatcherService_DeleteWatcher = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.nokamoto.github.com.egosla.api.WatcherServiceClient.prototype.deleteWatcher =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/DeleteWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_DeleteWatcher,
      callback);
};


/**
 * @param {!proto.nokamoto.github.com.egosla.api.DeleteWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.nokamoto.github.com.egosla.api.WatcherServicePromiseClient.prototype.deleteWatcher =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/DeleteWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_DeleteWatcher);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.nokamoto.github.com.egosla.api.GetWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.Watcher>}
 */
const methodDescriptor_WatcherService_GetWatcher = new grpc.web.MethodDescriptor(
  '/nokamoto.github.com.egosla.api.WatcherService/GetWatcher',
  grpc.web.MethodType.UNARY,
  proto.nokamoto.github.com.egosla.api.GetWatcherRequest,
  proto.nokamoto.github.com.egosla.api.Watcher,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.GetWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.Watcher.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.nokamoto.github.com.egosla.api.GetWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.Watcher>}
 */
const methodInfo_WatcherService_GetWatcher = new grpc.web.AbstractClientBase.MethodInfo(
  proto.nokamoto.github.com.egosla.api.Watcher,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.GetWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.Watcher.deserializeBinary
);


/**
 * @param {!proto.nokamoto.github.com.egosla.api.GetWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.nokamoto.github.com.egosla.api.Watcher)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.nokamoto.github.com.egosla.api.Watcher>|undefined}
 *     The XHR Node Readable Stream
 */
proto.nokamoto.github.com.egosla.api.WatcherServiceClient.prototype.getWatcher =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/GetWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_GetWatcher,
      callback);
};


/**
 * @param {!proto.nokamoto.github.com.egosla.api.GetWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.nokamoto.github.com.egosla.api.Watcher>}
 *     Promise that resolves to the response
 */
proto.nokamoto.github.com.egosla.api.WatcherServicePromiseClient.prototype.getWatcher =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/GetWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_GetWatcher);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.nokamoto.github.com.egosla.api.ListWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.ListWatcherResponse>}
 */
const methodDescriptor_WatcherService_ListWatcher = new grpc.web.MethodDescriptor(
  '/nokamoto.github.com.egosla.api.WatcherService/ListWatcher',
  grpc.web.MethodType.UNARY,
  proto.nokamoto.github.com.egosla.api.ListWatcherRequest,
  proto.nokamoto.github.com.egosla.api.ListWatcherResponse,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.ListWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.ListWatcherResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.nokamoto.github.com.egosla.api.ListWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.ListWatcherResponse>}
 */
const methodInfo_WatcherService_ListWatcher = new grpc.web.AbstractClientBase.MethodInfo(
  proto.nokamoto.github.com.egosla.api.ListWatcherResponse,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.ListWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.ListWatcherResponse.deserializeBinary
);


/**
 * @param {!proto.nokamoto.github.com.egosla.api.ListWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.nokamoto.github.com.egosla.api.ListWatcherResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.nokamoto.github.com.egosla.api.ListWatcherResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.nokamoto.github.com.egosla.api.WatcherServiceClient.prototype.listWatcher =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/ListWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_ListWatcher,
      callback);
};


/**
 * @param {!proto.nokamoto.github.com.egosla.api.ListWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.nokamoto.github.com.egosla.api.ListWatcherResponse>}
 *     Promise that resolves to the response
 */
proto.nokamoto.github.com.egosla.api.WatcherServicePromiseClient.prototype.listWatcher =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/ListWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_ListWatcher);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.Watcher>}
 */
const methodDescriptor_WatcherService_UpdateWatcher = new grpc.web.MethodDescriptor(
  '/nokamoto.github.com.egosla.api.WatcherService/UpdateWatcher',
  grpc.web.MethodType.UNARY,
  proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest,
  proto.nokamoto.github.com.egosla.api.Watcher,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.Watcher.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest,
 *   !proto.nokamoto.github.com.egosla.api.Watcher>}
 */
const methodInfo_WatcherService_UpdateWatcher = new grpc.web.AbstractClientBase.MethodInfo(
  proto.nokamoto.github.com.egosla.api.Watcher,
  /**
   * @param {!proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.nokamoto.github.com.egosla.api.Watcher.deserializeBinary
);


/**
 * @param {!proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.nokamoto.github.com.egosla.api.Watcher)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.nokamoto.github.com.egosla.api.Watcher>|undefined}
 *     The XHR Node Readable Stream
 */
proto.nokamoto.github.com.egosla.api.WatcherServiceClient.prototype.updateWatcher =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/UpdateWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_UpdateWatcher,
      callback);
};


/**
 * @param {!proto.nokamoto.github.com.egosla.api.UpdateWatcherRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.nokamoto.github.com.egosla.api.Watcher>}
 *     Promise that resolves to the response
 */
proto.nokamoto.github.com.egosla.api.WatcherServicePromiseClient.prototype.updateWatcher =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/UpdateWatcher',
      request,
      metadata || {},
      methodDescriptor_WatcherService_UpdateWatcher);
};


module.exports = proto.nokamoto.github.com.egosla.api;

