/**
 * @fileoverview gRPC-Web generated client stub for nokamoto.github.com.egosla.api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as api_service_pb from '../api/service_pb';


export class WatcherServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCreateWatcher = new grpcWeb.AbstractClientBase.MethodInfo(
    api_service_pb.Watcher,
    (request: api_service_pb.CreateWatcherRequest) => {
      return request.serializeBinary();
    },
    api_service_pb.Watcher.deserializeBinary
  );

  createWatcher(
    request: api_service_pb.CreateWatcherRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_service_pb.Watcher>;

  createWatcher(
    request: api_service_pb.CreateWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_service_pb.Watcher) => void): grpcWeb.ClientReadableStream<api_service_pb.Watcher>;

  createWatcher(
    request: api_service_pb.CreateWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_service_pb.Watcher) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.WatcherService/CreateWatcher',
        request,
        metadata || {},
        this.methodInfoCreateWatcher,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/CreateWatcher',
    request,
    metadata || {},
    this.methodInfoCreateWatcher);
  }

  methodInfoDeleteWatcher = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: api_service_pb.DeleteWatcherRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteWatcher(
    request: api_service_pb.DeleteWatcherRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteWatcher(
    request: api_service_pb.DeleteWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteWatcher(
    request: api_service_pb.DeleteWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.WatcherService/DeleteWatcher',
        request,
        metadata || {},
        this.methodInfoDeleteWatcher,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/DeleteWatcher',
    request,
    metadata || {},
    this.methodInfoDeleteWatcher);
  }

  methodInfoGetWatcher = new grpcWeb.AbstractClientBase.MethodInfo(
    api_service_pb.Watcher,
    (request: api_service_pb.GetWatcherRequest) => {
      return request.serializeBinary();
    },
    api_service_pb.Watcher.deserializeBinary
  );

  getWatcher(
    request: api_service_pb.GetWatcherRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_service_pb.Watcher>;

  getWatcher(
    request: api_service_pb.GetWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_service_pb.Watcher) => void): grpcWeb.ClientReadableStream<api_service_pb.Watcher>;

  getWatcher(
    request: api_service_pb.GetWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_service_pb.Watcher) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.WatcherService/GetWatcher',
        request,
        metadata || {},
        this.methodInfoGetWatcher,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/GetWatcher',
    request,
    metadata || {},
    this.methodInfoGetWatcher);
  }

  methodInfoListWatcher = new grpcWeb.AbstractClientBase.MethodInfo(
    api_service_pb.ListWatcherResponse,
    (request: api_service_pb.ListWatcherRequest) => {
      return request.serializeBinary();
    },
    api_service_pb.ListWatcherResponse.deserializeBinary
  );

  listWatcher(
    request: api_service_pb.ListWatcherRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_service_pb.ListWatcherResponse>;

  listWatcher(
    request: api_service_pb.ListWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_service_pb.ListWatcherResponse) => void): grpcWeb.ClientReadableStream<api_service_pb.ListWatcherResponse>;

  listWatcher(
    request: api_service_pb.ListWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_service_pb.ListWatcherResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.WatcherService/ListWatcher',
        request,
        metadata || {},
        this.methodInfoListWatcher,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/ListWatcher',
    request,
    metadata || {},
    this.methodInfoListWatcher);
  }

  methodInfoUpdateWatcher = new grpcWeb.AbstractClientBase.MethodInfo(
    api_service_pb.Watcher,
    (request: api_service_pb.UpdateWatcherRequest) => {
      return request.serializeBinary();
    },
    api_service_pb.Watcher.deserializeBinary
  );

  updateWatcher(
    request: api_service_pb.UpdateWatcherRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_service_pb.Watcher>;

  updateWatcher(
    request: api_service_pb.UpdateWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_service_pb.Watcher) => void): grpcWeb.ClientReadableStream<api_service_pb.Watcher>;

  updateWatcher(
    request: api_service_pb.UpdateWatcherRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_service_pb.Watcher) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.WatcherService/UpdateWatcher',
        request,
        metadata || {},
        this.methodInfoUpdateWatcher,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.WatcherService/UpdateWatcher',
    request,
    metadata || {},
    this.methodInfoUpdateWatcher);
  }

}

