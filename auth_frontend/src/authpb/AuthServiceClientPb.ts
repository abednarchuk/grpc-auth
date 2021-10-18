/**
 * @fileoverview gRPC-Web generated client stub for auth
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as auth_pb from './auth_pb';


export class SignupServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSignUp = new grpcWeb.AbstractClientBase.MethodInfo(
    auth_pb.SignupResponse,
    (request: auth_pb.SignupRequest) => {
      return request.serializeBinary();
    },
    auth_pb.SignupResponse.deserializeBinary
  );

  signUp(
    request: auth_pb.SignupRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.SignupResponse>;

  signUp(
    request: auth_pb.SignupRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: auth_pb.SignupResponse) => void): grpcWeb.ClientReadableStream<auth_pb.SignupResponse>;

  signUp(
    request: auth_pb.SignupRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: auth_pb.SignupResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.SignupService/SignUp',
        request,
        metadata || {},
        this.methodInfoSignUp,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.SignupService/SignUp',
    request,
    metadata || {},
    this.methodInfoSignUp);
  }

  methodInfoIsUsernameAvailable = new grpcWeb.AbstractClientBase.MethodInfo(
    auth_pb.AvailabilityResponse,
    (request: auth_pb.UsernameAvailabilityRequest) => {
      return request.serializeBinary();
    },
    auth_pb.AvailabilityResponse.deserializeBinary
  );

  isUsernameAvailable(
    request: auth_pb.UsernameAvailabilityRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.AvailabilityResponse>;

  isUsernameAvailable(
    request: auth_pb.UsernameAvailabilityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: auth_pb.AvailabilityResponse) => void): grpcWeb.ClientReadableStream<auth_pb.AvailabilityResponse>;

  isUsernameAvailable(
    request: auth_pb.UsernameAvailabilityRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: auth_pb.AvailabilityResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.SignupService/IsUsernameAvailable',
        request,
        metadata || {},
        this.methodInfoIsUsernameAvailable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.SignupService/IsUsernameAvailable',
    request,
    metadata || {},
    this.methodInfoIsUsernameAvailable);
  }

  methodInfoIsEmailAvailable = new grpcWeb.AbstractClientBase.MethodInfo(
    auth_pb.AvailabilityResponse,
    (request: auth_pb.EmailAvailabilityRequest) => {
      return request.serializeBinary();
    },
    auth_pb.AvailabilityResponse.deserializeBinary
  );

  isEmailAvailable(
    request: auth_pb.EmailAvailabilityRequest,
    metadata: grpcWeb.Metadata | null): Promise<auth_pb.AvailabilityResponse>;

  isEmailAvailable(
    request: auth_pb.EmailAvailabilityRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: auth_pb.AvailabilityResponse) => void): grpcWeb.ClientReadableStream<auth_pb.AvailabilityResponse>;

  isEmailAvailable(
    request: auth_pb.EmailAvailabilityRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: auth_pb.AvailabilityResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/auth.SignupService/IsEmailAvailable',
        request,
        metadata || {},
        this.methodInfoIsEmailAvailable,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/auth.SignupService/IsEmailAvailable',
    request,
    metadata || {},
    this.methodInfoIsEmailAvailable);
  }

}

