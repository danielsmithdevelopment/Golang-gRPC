// package: users
// file: users_service.proto

import * as users_service_pb from "./users_service_pb";
import {grpc} from "grpc-web-client";

type UserServiceAddUser = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_service_pb.AddUserRequest;
  readonly responseType: typeof users_service_pb.User;
};

type UserServiceGetUsers = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_service_pb.GetUsersRequest;
  readonly responseType: typeof users_service_pb.GetUsersResponse;
};

type UserServiceDeleteUser = {
  readonly methodName: string;
  readonly service: typeof UserService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof users_service_pb.DeleteUserRequest;
  readonly responseType: typeof users_service_pb.User;
};

export class UserService {
  static readonly serviceName: string;
  static readonly AddUser: UserServiceAddUser;
  static readonly GetUsers: UserServiceGetUsers;
  static readonly DeleteUser: UserServiceDeleteUser;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class UserServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  addUser(
    requestMessage: users_service_pb.AddUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: users_service_pb.User|null) => void
  ): void;
  addUser(
    requestMessage: users_service_pb.AddUserRequest,
    callback: (error: ServiceError, responseMessage: users_service_pb.User|null) => void
  ): void;
  getUsers(
    requestMessage: users_service_pb.GetUsersRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: users_service_pb.GetUsersResponse|null) => void
  ): void;
  getUsers(
    requestMessage: users_service_pb.GetUsersRequest,
    callback: (error: ServiceError, responseMessage: users_service_pb.GetUsersResponse|null) => void
  ): void;
  deleteUser(
    requestMessage: users_service_pb.DeleteUserRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: users_service_pb.User|null) => void
  ): void;
  deleteUser(
    requestMessage: users_service_pb.DeleteUserRequest,
    callback: (error: ServiceError, responseMessage: users_service_pb.User|null) => void
  ): void;
}

