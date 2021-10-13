import * as jspb from 'google-protobuf'



export class SignupRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): SignupRequest;

  getPassword(): string;
  setPassword(value: string): SignupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SignupRequest): SignupRequest.AsObject;
  static serializeBinaryToWriter(message: SignupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignupRequest;
  static deserializeBinaryFromReader(message: SignupRequest, reader: jspb.BinaryReader): SignupRequest;
}

export namespace SignupRequest {
  export type AsObject = {
    email: string,
    password: string,
  }
}

export class SignupResponse extends jspb.Message {
  getResponseStatus(): ResponseStatus;
  setResponseStatus(value: ResponseStatus): SignupResponse;

  getSignupErrorsList(): Array<SignupError>;
  setSignupErrorsList(value: Array<SignupError>): SignupResponse;
  clearSignupErrorsList(): SignupResponse;
  addSignupErrors(value?: SignupError, index?: number): SignupError;

  getResponse(): string;
  setResponse(value: string): SignupResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SignupResponse): SignupResponse.AsObject;
  static serializeBinaryToWriter(message: SignupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignupResponse;
  static deserializeBinaryFromReader(message: SignupResponse, reader: jspb.BinaryReader): SignupResponse;
}

export namespace SignupResponse {
  export type AsObject = {
    responseStatus: ResponseStatus,
    signupErrorsList: Array<SignupError.AsObject>,
    response: string,
  }
}

export class SignupError extends jspb.Message {
  getErrorMessage(): string;
  setErrorMessage(value: string): SignupError;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignupError.AsObject;
  static toObject(includeInstance: boolean, msg: SignupError): SignupError.AsObject;
  static serializeBinaryToWriter(message: SignupError, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignupError;
  static deserializeBinaryFromReader(message: SignupError, reader: jspb.BinaryReader): SignupError;
}

export namespace SignupError {
  export type AsObject = {
    errorMessage: string,
  }
}

export enum ResponseStatus { 
  STATUS_FAIL = 0,
  STATUS_SUCCESS = 1,
}
