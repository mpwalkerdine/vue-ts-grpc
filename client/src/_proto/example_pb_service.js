// package: 
// file: example.proto

var example_pb = require("./example_pb");
var grpc = require("grpc-web-client").grpc;

var ExampleService = (function () {
  function ExampleService() {}
  ExampleService.serviceName = "ExampleService";
  return ExampleService;
}());

ExampleService.CreateExample = {
  methodName: "CreateExample",
  service: ExampleService,
  requestStream: false,
  responseStream: false,
  requestType: example_pb.CreateExampleRequest,
  responseType: example_pb.Example
};

ExampleService.ListExamples = {
  methodName: "ListExamples",
  service: ExampleService,
  requestStream: false,
  responseStream: false,
  requestType: example_pb.ListExamplesRequest,
  responseType: example_pb.ListExamplesResponse
};

ExampleService.DeleteExample = {
  methodName: "DeleteExample",
  service: ExampleService,
  requestStream: false,
  responseStream: false,
  requestType: example_pb.DeleteExampleRequest,
  responseType: example_pb.Empty
};

exports.ExampleService = ExampleService;

function ExampleServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ExampleServiceClient.prototype.createExample = function createExample(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ExampleService.CreateExample, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ExampleServiceClient.prototype.listExamples = function listExamples(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ExampleService.ListExamples, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ExampleServiceClient.prototype.deleteExample = function deleteExample(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ExampleService.DeleteExample, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ExampleServiceClient = ExampleServiceClient;

