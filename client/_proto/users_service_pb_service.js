// package: users
// file: users_service.proto

var users_service_pb = require("./users_service_pb");
var grpc = require("grpc-web-client").grpc;

var UserService = (function () {
  function UserService() {}
  UserService.serviceName = "users.UserService";
  return UserService;
}());

UserService.AddUser = {
  methodName: "AddUser",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: users_service_pb.AddUserRequest,
  responseType: users_service_pb.User
};

UserService.GetUsers = {
  methodName: "GetUsers",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: users_service_pb.GetUsersRequest,
  responseType: users_service_pb.GetUsersResponse
};

UserService.DeleteUser = {
  methodName: "DeleteUser",
  service: UserService,
  requestStream: false,
  responseStream: false,
  requestType: users_service_pb.DeleteUserRequest,
  responseType: users_service_pb.User
};

exports.UserService = UserService;

function UserServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

UserServiceClient.prototype.addUser = function addUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(UserService.AddUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

UserServiceClient.prototype.getUsers = function getUsers(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(UserService.GetUsers, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

UserServiceClient.prototype.deleteUser = function deleteUser(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(UserService.DeleteUser, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.UserServiceClient = UserServiceClient;

