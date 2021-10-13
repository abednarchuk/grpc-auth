import * as jspb from 'google-protobuf'



export class User extends jspb.Message {
  getId(): string;
  setId(value: string): User;

  getUserName(): string;
  setUserName(value: string): User;

  getEmail(): string;
  setEmail(value: string): User;

  getPassword(): string;
  setPassword(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    userName: string,
    email: string,
    password: string,
  }
}

export class SignupRequest extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): SignupRequest;
  hasUser(): boolean;
  clearUser(): SignupRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SignupRequest): SignupRequest.AsObject;
  static serializeBinaryToWriter(message: SignupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignupRequest;
  static deserializeBinaryFromReader(message: SignupRequest, reader: jspb.BinaryReader): SignupRequest;
}

export namespace SignupRequest {
  export type AsObject = {
    user?: User.AsObject,
  }
}

export class SignupResponse extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): SignupResponse;
  hasUser(): boolean;
  clearUser(): SignupResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignupResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SignupResponse): SignupResponse.AsObject;
  static serializeBinaryToWriter(message: SignupResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignupResponse;
  static deserializeBinaryFromReader(message: SignupResponse, reader: jspb.BinaryReader): SignupResponse;
}

export namespace SignupResponse {
  export type AsObject = {
    user?: User.AsObject,
  }
}

