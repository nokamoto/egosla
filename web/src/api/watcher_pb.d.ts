import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';


export class Watcher extends jspb.Message {
  getName(): string;
  setName(value: string): Watcher;

  getKeywordsList(): Array<string>;
  setKeywordsList(value: Array<string>): Watcher;
  clearKeywordsList(): Watcher;
  addKeywords(value: string, index?: number): Watcher;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Watcher.AsObject;
  static toObject(includeInstance: boolean, msg: Watcher): Watcher.AsObject;
  static serializeBinaryToWriter(message: Watcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Watcher;
  static deserializeBinaryFromReader(message: Watcher, reader: jspb.BinaryReader): Watcher;
}

export namespace Watcher {
  export type AsObject = {
    name: string,
    keywordsList: Array<string>,
  }
}

export class CreateWatcherRequest extends jspb.Message {
  getWatcher(): Watcher | undefined;
  setWatcher(value?: Watcher): CreateWatcherRequest;
  hasWatcher(): boolean;
  clearWatcher(): CreateWatcherRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateWatcherRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateWatcherRequest): CreateWatcherRequest.AsObject;
  static serializeBinaryToWriter(message: CreateWatcherRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateWatcherRequest;
  static deserializeBinaryFromReader(message: CreateWatcherRequest, reader: jspb.BinaryReader): CreateWatcherRequest;
}

export namespace CreateWatcherRequest {
  export type AsObject = {
    watcher?: Watcher.AsObject,
  }
}

export class DeleteWatcherRequest extends jspb.Message {
  getName(): string;
  setName(value: string): DeleteWatcherRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteWatcherRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteWatcherRequest): DeleteWatcherRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteWatcherRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteWatcherRequest;
  static deserializeBinaryFromReader(message: DeleteWatcherRequest, reader: jspb.BinaryReader): DeleteWatcherRequest;
}

export namespace DeleteWatcherRequest {
  export type AsObject = {
    name: string,
  }
}

export class GetWatcherRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetWatcherRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetWatcherRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetWatcherRequest): GetWatcherRequest.AsObject;
  static serializeBinaryToWriter(message: GetWatcherRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetWatcherRequest;
  static deserializeBinaryFromReader(message: GetWatcherRequest, reader: jspb.BinaryReader): GetWatcherRequest;
}

export namespace GetWatcherRequest {
  export type AsObject = {
    name: string,
  }
}

export class ListWatcherRequest extends jspb.Message {
  getPageToken(): string;
  setPageToken(value: string): ListWatcherRequest;

  getPageSize(): number;
  setPageSize(value: number): ListWatcherRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListWatcherRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListWatcherRequest): ListWatcherRequest.AsObject;
  static serializeBinaryToWriter(message: ListWatcherRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListWatcherRequest;
  static deserializeBinaryFromReader(message: ListWatcherRequest, reader: jspb.BinaryReader): ListWatcherRequest;
}

export namespace ListWatcherRequest {
  export type AsObject = {
    pageToken: string,
    pageSize: number,
  }
}

export class ListWatcherResponse extends jspb.Message {
  getNextPageToken(): string;
  setNextPageToken(value: string): ListWatcherResponse;

  getWatchersList(): Array<Watcher>;
  setWatchersList(value: Array<Watcher>): ListWatcherResponse;
  clearWatchersList(): ListWatcherResponse;
  addWatchers(value?: Watcher, index?: number): Watcher;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListWatcherResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListWatcherResponse): ListWatcherResponse.AsObject;
  static serializeBinaryToWriter(message: ListWatcherResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListWatcherResponse;
  static deserializeBinaryFromReader(message: ListWatcherResponse, reader: jspb.BinaryReader): ListWatcherResponse;
}

export namespace ListWatcherResponse {
  export type AsObject = {
    nextPageToken: string,
    watchersList: Array<Watcher.AsObject>,
  }
}

export class UpdateWatcherRequest extends jspb.Message {
  getName(): string;
  setName(value: string): UpdateWatcherRequest;

  getWatcher(): Watcher | undefined;
  setWatcher(value?: Watcher): UpdateWatcherRequest;
  hasWatcher(): boolean;
  clearWatcher(): UpdateWatcherRequest;

  getUpdateMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setUpdateMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateWatcherRequest;
  hasUpdateMask(): boolean;
  clearUpdateMask(): UpdateWatcherRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateWatcherRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateWatcherRequest): UpdateWatcherRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateWatcherRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateWatcherRequest;
  static deserializeBinaryFromReader(message: UpdateWatcherRequest, reader: jspb.BinaryReader): UpdateWatcherRequest;
}

export namespace UpdateWatcherRequest {
  export type AsObject = {
    name: string,
    watcher?: Watcher.AsObject,
    updateMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

