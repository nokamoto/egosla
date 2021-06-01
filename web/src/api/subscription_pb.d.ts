import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';


export class Subscription extends jspb.Message {
  getName(): string;
  setName(value: string): Subscription;

  getWatcher(): string;
  setWatcher(value: string): Subscription;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Subscription.AsObject;
  static toObject(includeInstance: boolean, msg: Subscription): Subscription.AsObject;
  static serializeBinaryToWriter(message: Subscription, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Subscription;
  static deserializeBinaryFromReader(message: Subscription, reader: jspb.BinaryReader): Subscription;
}

export namespace Subscription {
  export type AsObject = {
    name: string,
    watcher: string,
  }
}

export class CreateSubscriptionRequest extends jspb.Message {
  getSubscription(): Subscription | undefined;
  setSubscription(value?: Subscription): CreateSubscriptionRequest;
  hasSubscription(): boolean;
  clearSubscription(): CreateSubscriptionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSubscriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSubscriptionRequest): CreateSubscriptionRequest.AsObject;
  static serializeBinaryToWriter(message: CreateSubscriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSubscriptionRequest;
  static deserializeBinaryFromReader(message: CreateSubscriptionRequest, reader: jspb.BinaryReader): CreateSubscriptionRequest;
}

export namespace CreateSubscriptionRequest {
  export type AsObject = {
    subscription?: Subscription.AsObject,
  }
}

export class DeleteSubscriptionRequest extends jspb.Message {
  getName(): string;
  setName(value: string): DeleteSubscriptionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteSubscriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteSubscriptionRequest): DeleteSubscriptionRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteSubscriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteSubscriptionRequest;
  static deserializeBinaryFromReader(message: DeleteSubscriptionRequest, reader: jspb.BinaryReader): DeleteSubscriptionRequest;
}

export namespace DeleteSubscriptionRequest {
  export type AsObject = {
    name: string,
  }
}

export class GetSubscriptionRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetSubscriptionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSubscriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSubscriptionRequest): GetSubscriptionRequest.AsObject;
  static serializeBinaryToWriter(message: GetSubscriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSubscriptionRequest;
  static deserializeBinaryFromReader(message: GetSubscriptionRequest, reader: jspb.BinaryReader): GetSubscriptionRequest;
}

export namespace GetSubscriptionRequest {
  export type AsObject = {
    name: string,
  }
}

export class ListSubscriptionRequest extends jspb.Message {
  getPageToken(): string;
  setPageToken(value: string): ListSubscriptionRequest;

  getPageSize(): number;
  setPageSize(value: number): ListSubscriptionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSubscriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListSubscriptionRequest): ListSubscriptionRequest.AsObject;
  static serializeBinaryToWriter(message: ListSubscriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSubscriptionRequest;
  static deserializeBinaryFromReader(message: ListSubscriptionRequest, reader: jspb.BinaryReader): ListSubscriptionRequest;
}

export namespace ListSubscriptionRequest {
  export type AsObject = {
    pageToken: string,
    pageSize: number,
  }
}

export class ListSubscriptionResponse extends jspb.Message {
  getNextPageToken(): string;
  setNextPageToken(value: string): ListSubscriptionResponse;

  getSubscriptionsList(): Array<Subscription>;
  setSubscriptionsList(value: Array<Subscription>): ListSubscriptionResponse;
  clearSubscriptionsList(): ListSubscriptionResponse;
  addSubscriptions(value?: Subscription, index?: number): Subscription;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSubscriptionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListSubscriptionResponse): ListSubscriptionResponse.AsObject;
  static serializeBinaryToWriter(message: ListSubscriptionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSubscriptionResponse;
  static deserializeBinaryFromReader(message: ListSubscriptionResponse, reader: jspb.BinaryReader): ListSubscriptionResponse;
}

export namespace ListSubscriptionResponse {
  export type AsObject = {
    nextPageToken: string,
    subscriptionsList: Array<Subscription.AsObject>,
  }
}

export class UpdateSubscriptionRequest extends jspb.Message {
  getName(): string;
  setName(value: string): UpdateSubscriptionRequest;

  getSubscription(): Subscription | undefined;
  setSubscription(value?: Subscription): UpdateSubscriptionRequest;
  hasSubscription(): boolean;
  clearSubscription(): UpdateSubscriptionRequest;

  getUpdateMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setUpdateMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateSubscriptionRequest;
  hasUpdateMask(): boolean;
  clearUpdateMask(): UpdateSubscriptionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateSubscriptionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateSubscriptionRequest): UpdateSubscriptionRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateSubscriptionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateSubscriptionRequest;
  static deserializeBinaryFromReader(message: UpdateSubscriptionRequest, reader: jspb.BinaryReader): UpdateSubscriptionRequest;
}

export namespace UpdateSubscriptionRequest {
  export type AsObject = {
    name: string,
    subscription?: Subscription.AsObject,
    updateMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

