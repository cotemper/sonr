/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Did } from "../registry/did";

export const protobufPackage = "sonrio.sonr.registry";

export interface MsgRegisterName {
  creator: string;
  name: string;
  fingerprint: string;
}

export interface MsgRegisterNameResponse {
  code: number;
  did: Did | undefined;
}

export interface MsgRegisterService {
  creator: string;
  serviceName: string;
  publicKey: string;
}

export interface MsgRegisterServiceResponse {
  code: number;
  did: Did | undefined;
}

export interface MsgCreateAccount {
  creator: string;
  home: string;
  deviceId: string;
  fingerprint: string;
  os: string;
  model: string;
  arch: string;
  publicKey: string;
  metadata: string;
}

export interface MsgCreateAccountResponse {
  code: number;
  did: Did | undefined;
}

const baseMsgRegisterName: object = { creator: "", name: "", fingerprint: "" };

export const MsgRegisterName = {
  encode(message: MsgRegisterName, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.fingerprint !== "") {
      writer.uint32(26).string(message.fingerprint);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterName {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterName } as MsgRegisterName;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.fingerprint = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterName {
    const message = { ...baseMsgRegisterName } as MsgRegisterName;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.fingerprint !== undefined && object.fingerprint !== null) {
      message.fingerprint = String(object.fingerprint);
    } else {
      message.fingerprint = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterName): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.fingerprint !== undefined &&
      (obj.fingerprint = message.fingerprint);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRegisterName>): MsgRegisterName {
    const message = { ...baseMsgRegisterName } as MsgRegisterName;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.fingerprint !== undefined && object.fingerprint !== null) {
      message.fingerprint = object.fingerprint;
    } else {
      message.fingerprint = "";
    }
    return message;
  },
};

const baseMsgRegisterNameResponse: object = { code: 0 };

export const MsgRegisterNameResponse = {
  encode(
    message: MsgRegisterNameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterNameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterNameResponse,
    } as MsgRegisterNameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterNameResponse {
    const message = {
      ...baseMsgRegisterNameResponse,
    } as MsgRegisterNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromJSON(object.did);
    } else {
      message.did = undefined;
    }
    return message;
  },

  toJSON(message: MsgRegisterNameResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterNameResponse>
  ): MsgRegisterNameResponse {
    const message = {
      ...baseMsgRegisterNameResponse,
    } as MsgRegisterNameResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromPartial(object.did);
    } else {
      message.did = undefined;
    }
    return message;
  },
};

const baseMsgRegisterService: object = {
  creator: "",
  serviceName: "",
  publicKey: "",
};

export const MsgRegisterService = {
  encode(
    message: MsgRegisterService,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.serviceName !== "") {
      writer.uint32(18).string(message.serviceName);
    }
    if (message.publicKey !== "") {
      writer.uint32(26).string(message.publicKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRegisterService {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRegisterService } as MsgRegisterService;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.serviceName = reader.string();
          break;
        case 3:
          message.publicKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterService {
    const message = { ...baseMsgRegisterService } as MsgRegisterService;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.serviceName !== undefined && object.serviceName !== null) {
      message.serviceName = String(object.serviceName);
    } else {
      message.serviceName = "";
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = String(object.publicKey);
    } else {
      message.publicKey = "";
    }
    return message;
  },

  toJSON(message: MsgRegisterService): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.serviceName !== undefined &&
      (obj.serviceName = message.serviceName);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRegisterService>): MsgRegisterService {
    const message = { ...baseMsgRegisterService } as MsgRegisterService;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.serviceName !== undefined && object.serviceName !== null) {
      message.serviceName = object.serviceName;
    } else {
      message.serviceName = "";
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = object.publicKey;
    } else {
      message.publicKey = "";
    }
    return message;
  },
};

const baseMsgRegisterServiceResponse: object = { code: 0 };

export const MsgRegisterServiceResponse = {
  encode(
    message: MsgRegisterServiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRegisterServiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRegisterServiceResponse,
    } as MsgRegisterServiceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterServiceResponse {
    const message = {
      ...baseMsgRegisterServiceResponse,
    } as MsgRegisterServiceResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromJSON(object.did);
    } else {
      message.did = undefined;
    }
    return message;
  },

  toJSON(message: MsgRegisterServiceResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRegisterServiceResponse>
  ): MsgRegisterServiceResponse {
    const message = {
      ...baseMsgRegisterServiceResponse,
    } as MsgRegisterServiceResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromPartial(object.did);
    } else {
      message.did = undefined;
    }
    return message;
  },
};

const baseMsgCreateAccount: object = {
  creator: "",
  home: "",
  deviceId: "",
  fingerprint: "",
  os: "",
  model: "",
  arch: "",
  publicKey: "",
  metadata: "",
};

export const MsgCreateAccount = {
  encode(message: MsgCreateAccount, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.home !== "") {
      writer.uint32(18).string(message.home);
    }
    if (message.deviceId !== "") {
      writer.uint32(26).string(message.deviceId);
    }
    if (message.fingerprint !== "") {
      writer.uint32(34).string(message.fingerprint);
    }
    if (message.os !== "") {
      writer.uint32(42).string(message.os);
    }
    if (message.model !== "") {
      writer.uint32(50).string(message.model);
    }
    if (message.arch !== "") {
      writer.uint32(58).string(message.arch);
    }
    if (message.publicKey !== "") {
      writer.uint32(66).string(message.publicKey);
    }
    if (message.metadata !== "") {
      writer.uint32(74).string(message.metadata);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateAccount } as MsgCreateAccount;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.home = reader.string();
          break;
        case 3:
          message.deviceId = reader.string();
          break;
        case 4:
          message.fingerprint = reader.string();
          break;
        case 5:
          message.os = reader.string();
          break;
        case 6:
          message.model = reader.string();
          break;
        case 7:
          message.arch = reader.string();
          break;
        case 8:
          message.publicKey = reader.string();
          break;
        case 9:
          message.metadata = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAccount {
    const message = { ...baseMsgCreateAccount } as MsgCreateAccount;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.home !== undefined && object.home !== null) {
      message.home = String(object.home);
    } else {
      message.home = "";
    }
    if (object.deviceId !== undefined && object.deviceId !== null) {
      message.deviceId = String(object.deviceId);
    } else {
      message.deviceId = "";
    }
    if (object.fingerprint !== undefined && object.fingerprint !== null) {
      message.fingerprint = String(object.fingerprint);
    } else {
      message.fingerprint = "";
    }
    if (object.os !== undefined && object.os !== null) {
      message.os = String(object.os);
    } else {
      message.os = "";
    }
    if (object.model !== undefined && object.model !== null) {
      message.model = String(object.model);
    } else {
      message.model = "";
    }
    if (object.arch !== undefined && object.arch !== null) {
      message.arch = String(object.arch);
    } else {
      message.arch = "";
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = String(object.publicKey);
    } else {
      message.publicKey = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = String(object.metadata);
    } else {
      message.metadata = "";
    }
    return message;
  },

  toJSON(message: MsgCreateAccount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.home !== undefined && (obj.home = message.home);
    message.deviceId !== undefined && (obj.deviceId = message.deviceId);
    message.fingerprint !== undefined &&
      (obj.fingerprint = message.fingerprint);
    message.os !== undefined && (obj.os = message.os);
    message.model !== undefined && (obj.model = message.model);
    message.arch !== undefined && (obj.arch = message.arch);
    message.publicKey !== undefined && (obj.publicKey = message.publicKey);
    message.metadata !== undefined && (obj.metadata = message.metadata);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateAccount>): MsgCreateAccount {
    const message = { ...baseMsgCreateAccount } as MsgCreateAccount;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.home !== undefined && object.home !== null) {
      message.home = object.home;
    } else {
      message.home = "";
    }
    if (object.deviceId !== undefined && object.deviceId !== null) {
      message.deviceId = object.deviceId;
    } else {
      message.deviceId = "";
    }
    if (object.fingerprint !== undefined && object.fingerprint !== null) {
      message.fingerprint = object.fingerprint;
    } else {
      message.fingerprint = "";
    }
    if (object.os !== undefined && object.os !== null) {
      message.os = object.os;
    } else {
      message.os = "";
    }
    if (object.model !== undefined && object.model !== null) {
      message.model = object.model;
    } else {
      message.model = "";
    }
    if (object.arch !== undefined && object.arch !== null) {
      message.arch = object.arch;
    } else {
      message.arch = "";
    }
    if (object.publicKey !== undefined && object.publicKey !== null) {
      message.publicKey = object.publicKey;
    } else {
      message.publicKey = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = object.metadata;
    } else {
      message.metadata = "";
    }
    return message;
  },
};

const baseMsgCreateAccountResponse: object = { code: 0 };

export const MsgCreateAccountResponse = {
  encode(
    message: MsgCreateAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.did !== undefined) {
      Did.encode(message.did, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateAccountResponse,
    } as MsgCreateAccountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.int32();
          break;
        case 2:
          message.did = Did.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAccountResponse {
    const message = {
      ...baseMsgCreateAccountResponse,
    } as MsgCreateAccountResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = Number(object.code);
    } else {
      message.code = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromJSON(object.did);
    } else {
      message.did = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateAccountResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    message.did !== undefined &&
      (obj.did = message.did ? Did.toJSON(message.did) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateAccountResponse>
  ): MsgCreateAccountResponse {
    const message = {
      ...baseMsgCreateAccountResponse,
    } as MsgCreateAccountResponse;
    if (object.code !== undefined && object.code !== null) {
      message.code = object.code;
    } else {
      message.code = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = Did.fromPartial(object.did);
    } else {
      message.did = undefined;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse>;
  RegisterService(
    request: MsgRegisterService
  ): Promise<MsgRegisterServiceResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateAccount(request: MsgCreateAccount): Promise<MsgCreateAccountResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse> {
    const data = MsgRegisterName.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "RegisterName",
      data
    );
    return promise.then((data) =>
      MsgRegisterNameResponse.decode(new Reader(data))
    );
  }

  RegisterService(
    request: MsgRegisterService
  ): Promise<MsgRegisterServiceResponse> {
    const data = MsgRegisterService.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "RegisterService",
      data
    );
    return promise.then((data) =>
      MsgRegisterServiceResponse.decode(new Reader(data))
    );
  }

  CreateAccount(request: MsgCreateAccount): Promise<MsgCreateAccountResponse> {
    const data = MsgCreateAccount.encode(request).finish();
    const promise = this.rpc.request(
      "sonrio.sonr.registry.Msg",
      "CreateAccount",
      data
    );
    return promise.then((data) =>
      MsgCreateAccountResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
