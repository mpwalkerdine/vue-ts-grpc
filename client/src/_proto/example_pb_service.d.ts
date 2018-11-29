// package: 
// file: example.proto

import * as example_pb from "./example_pb";
import {grpc} from "grpc-web-client";

type ExampleServiceCreateExample = {
  readonly methodName: string;
  readonly service: typeof ExampleService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof example_pb.CreateExampleRequest;
  readonly responseType: typeof example_pb.Example;
};

type ExampleServiceListExamples = {
  readonly methodName: string;
  readonly service: typeof ExampleService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof example_pb.ListExamplesRequest;
  readonly responseType: typeof example_pb.ListExamplesResponse;
};

type ExampleServiceDeleteExample = {
  readonly methodName: string;
  readonly service: typeof ExampleService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof example_pb.DeleteExampleRequest;
  readonly responseType: typeof example_pb.Empty;
};

export class ExampleService {
  static readonly serviceName: string;
  static readonly CreateExample: ExampleServiceCreateExample;
  static readonly ListExamples: ExampleServiceListExamples;
  static readonly DeleteExample: ExampleServiceDeleteExample;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: () => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: () => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ExampleServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createExample(
    requestMessage: example_pb.CreateExampleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: example_pb.Example|null) => void
  ): UnaryResponse;
  createExample(
    requestMessage: example_pb.CreateExampleRequest,
    callback: (error: ServiceError|null, responseMessage: example_pb.Example|null) => void
  ): UnaryResponse;
  listExamples(
    requestMessage: example_pb.ListExamplesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: example_pb.ListExamplesResponse|null) => void
  ): UnaryResponse;
  listExamples(
    requestMessage: example_pb.ListExamplesRequest,
    callback: (error: ServiceError|null, responseMessage: example_pb.ListExamplesResponse|null) => void
  ): UnaryResponse;
  deleteExample(
    requestMessage: example_pb.DeleteExampleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: example_pb.Empty|null) => void
  ): UnaryResponse;
  deleteExample(
    requestMessage: example_pb.DeleteExampleRequest,
    callback: (error: ServiceError|null, responseMessage: example_pb.Empty|null) => void
  ): UnaryResponse;
}

