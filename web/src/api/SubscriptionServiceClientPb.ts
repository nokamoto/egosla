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
import * as api_subscription_pb from '../api/subscription_pb';


export class SubscriptionServiceClient {
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

  methodInfoCreateSubscription = new grpcWeb.AbstractClientBase.MethodInfo(
    api_subscription_pb.Subscription,
    (request: api_subscription_pb.CreateSubscriptionRequest) => {
      return request.serializeBinary();
    },
    api_subscription_pb.Subscription.deserializeBinary
  );

  createSubscription(
    request: api_subscription_pb.CreateSubscriptionRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_subscription_pb.Subscription>;

  createSubscription(
    request: api_subscription_pb.CreateSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_subscription_pb.Subscription) => void): grpcWeb.ClientReadableStream<api_subscription_pb.Subscription>;

  createSubscription(
    request: api_subscription_pb.CreateSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_subscription_pb.Subscription) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.SubscriptionService/CreateSubscription',
        request,
        metadata || {},
        this.methodInfoCreateSubscription,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.SubscriptionService/CreateSubscription',
    request,
    metadata || {},
    this.methodInfoCreateSubscription);
  }

  methodInfoDeleteSubscription = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: api_subscription_pb.DeleteSubscriptionRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteSubscription(
    request: api_subscription_pb.DeleteSubscriptionRequest,
    metadata: grpcWeb.Metadata | null): Promise<google_protobuf_empty_pb.Empty>;

  deleteSubscription(
    request: api_subscription_pb.DeleteSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  deleteSubscription(
    request: api_subscription_pb.DeleteSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.SubscriptionService/DeleteSubscription',
        request,
        metadata || {},
        this.methodInfoDeleteSubscription,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.SubscriptionService/DeleteSubscription',
    request,
    metadata || {},
    this.methodInfoDeleteSubscription);
  }

  methodInfoGetSubscription = new grpcWeb.AbstractClientBase.MethodInfo(
    api_subscription_pb.Subscription,
    (request: api_subscription_pb.GetSubscriptionRequest) => {
      return request.serializeBinary();
    },
    api_subscription_pb.Subscription.deserializeBinary
  );

  getSubscription(
    request: api_subscription_pb.GetSubscriptionRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_subscription_pb.Subscription>;

  getSubscription(
    request: api_subscription_pb.GetSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_subscription_pb.Subscription) => void): grpcWeb.ClientReadableStream<api_subscription_pb.Subscription>;

  getSubscription(
    request: api_subscription_pb.GetSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_subscription_pb.Subscription) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.SubscriptionService/GetSubscription',
        request,
        metadata || {},
        this.methodInfoGetSubscription,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.SubscriptionService/GetSubscription',
    request,
    metadata || {},
    this.methodInfoGetSubscription);
  }

  methodInfoListSubscription = new grpcWeb.AbstractClientBase.MethodInfo(
    api_subscription_pb.ListSubscriptionResponse,
    (request: api_subscription_pb.ListSubscriptionRequest) => {
      return request.serializeBinary();
    },
    api_subscription_pb.ListSubscriptionResponse.deserializeBinary
  );

  listSubscription(
    request: api_subscription_pb.ListSubscriptionRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_subscription_pb.ListSubscriptionResponse>;

  listSubscription(
    request: api_subscription_pb.ListSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_subscription_pb.ListSubscriptionResponse) => void): grpcWeb.ClientReadableStream<api_subscription_pb.ListSubscriptionResponse>;

  listSubscription(
    request: api_subscription_pb.ListSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_subscription_pb.ListSubscriptionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.SubscriptionService/ListSubscription',
        request,
        metadata || {},
        this.methodInfoListSubscription,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.SubscriptionService/ListSubscription',
    request,
    metadata || {},
    this.methodInfoListSubscription);
  }

  methodInfoUpdateSubscription = new grpcWeb.AbstractClientBase.MethodInfo(
    api_subscription_pb.Subscription,
    (request: api_subscription_pb.UpdateSubscriptionRequest) => {
      return request.serializeBinary();
    },
    api_subscription_pb.Subscription.deserializeBinary
  );

  updateSubscription(
    request: api_subscription_pb.UpdateSubscriptionRequest,
    metadata: grpcWeb.Metadata | null): Promise<api_subscription_pb.Subscription>;

  updateSubscription(
    request: api_subscription_pb.UpdateSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: api_subscription_pb.Subscription) => void): grpcWeb.ClientReadableStream<api_subscription_pb.Subscription>;

  updateSubscription(
    request: api_subscription_pb.UpdateSubscriptionRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: api_subscription_pb.Subscription) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/nokamoto.github.com.egosla.api.SubscriptionService/UpdateSubscription',
        request,
        metadata || {},
        this.methodInfoUpdateSubscription,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/nokamoto.github.com.egosla.api.SubscriptionService/UpdateSubscription',
    request,
    metadata || {},
    this.methodInfoUpdateSubscription);
  }

}

