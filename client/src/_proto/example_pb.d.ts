// package: 
// file: example.proto

import * as jspb from "google-protobuf";

export class CreateExampleRequest extends jspb.Message {
  getParent(): string;
  setParent(value: string): void;

  hasExample(): boolean;
  clearExample(): void;
  getExample(): Example | undefined;
  setExample(value?: Example): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateExampleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateExampleRequest): CreateExampleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateExampleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateExampleRequest;
  static deserializeBinaryFromReader(message: CreateExampleRequest, reader: jspb.BinaryReader): CreateExampleRequest;
}

export namespace CreateExampleRequest {
  export type AsObject = {
    parent: string,
    example?: Example.AsObject,
  }
}

export class Example extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getText(): string;
  setText(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Example.AsObject;
  static toObject(includeInstance: boolean, msg: Example): Example.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Example, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Example;
  static deserializeBinaryFromReader(message: Example, reader: jspb.BinaryReader): Example;
}

export namespace Example {
  export type AsObject = {
    name: string,
    text: string,
  }
}

export class ListExamplesRequest extends jspb.Message {
  getParent(): string;
  setParent(value: string): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  getPageToken(): string;
  setPageToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListExamplesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListExamplesRequest): ListExamplesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListExamplesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListExamplesRequest;
  static deserializeBinaryFromReader(message: ListExamplesRequest, reader: jspb.BinaryReader): ListExamplesRequest;
}

export namespace ListExamplesRequest {
  export type AsObject = {
    parent: string,
    pageSize: number,
    pageToken: string,
  }
}

export class ListExamplesResponse extends jspb.Message {
  clearExamplesList(): void;
  getExamplesList(): Array<Example>;
  setExamplesList(value: Array<Example>): void;
  addExamples(value?: Example, index?: number): Example;

  getNextPageToken(): string;
  setNextPageToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListExamplesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListExamplesResponse): ListExamplesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListExamplesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListExamplesResponse;
  static deserializeBinaryFromReader(message: ListExamplesResponse, reader: jspb.BinaryReader): ListExamplesResponse;
}

export namespace ListExamplesResponse {
  export type AsObject = {
    examplesList: Array<Example.AsObject>,
    nextPageToken: string,
  }
}

export class DeleteExampleRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteExampleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteExampleRequest): DeleteExampleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteExampleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteExampleRequest;
  static deserializeBinaryFromReader(message: DeleteExampleRequest, reader: jspb.BinaryReader): DeleteExampleRequest;
}

export namespace DeleteExampleRequest {
  export type AsObject = {
    name: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

