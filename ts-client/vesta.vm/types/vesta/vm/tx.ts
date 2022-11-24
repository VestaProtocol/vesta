/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "vesta.vm";

export interface MsgStore {
  creator: string;
  source: string;
}

export interface MsgStoreResponse {
  code: string;
}

export interface MsgExecute {
  creator: string;
  contract: string;
  function: string;
  args: string;
}

export interface MsgExecuteResponse {
  console: string;
}

export interface MsgInstantiate {
  creator: string;
  name: string;
  code: string;
}

export interface MsgInstantiateResponse {
}

function createBaseMsgStore(): MsgStore {
  return { creator: "", source: "" };
}

export const MsgStore = {
  encode(message: MsgStore, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.source !== "") {
      writer.uint32(18).string(message.source);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStore {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStore();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.source = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStore {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      source: isSet(object.source) ? String(object.source) : "",
    };
  },

  toJSON(message: MsgStore): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.source !== undefined && (obj.source = message.source);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStore>, I>>(object: I): MsgStore {
    const message = createBaseMsgStore();
    message.creator = object.creator ?? "";
    message.source = object.source ?? "";
    return message;
  },
};

function createBaseMsgStoreResponse(): MsgStoreResponse {
  return { code: "" };
}

export const MsgStoreResponse = {
  encode(message: MsgStoreResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStoreResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStoreResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStoreResponse {
    return { code: isSet(object.code) ? String(object.code) : "" };
  },

  toJSON(message: MsgStoreResponse): unknown {
    const obj: any = {};
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStoreResponse>, I>>(object: I): MsgStoreResponse {
    const message = createBaseMsgStoreResponse();
    message.code = object.code ?? "";
    return message;
  },
};

function createBaseMsgExecute(): MsgExecute {
  return { creator: "", contract: "", function: "", args: "" };
}

export const MsgExecute = {
  encode(message: MsgExecute, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.contract !== "") {
      writer.uint32(18).string(message.contract);
    }
    if (message.function !== "") {
      writer.uint32(26).string(message.function);
    }
    if (message.args !== "") {
      writer.uint32(34).string(message.args);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgExecute {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgExecute();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.contract = reader.string();
          break;
        case 3:
          message.function = reader.string();
          break;
        case 4:
          message.args = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExecute {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      contract: isSet(object.contract) ? String(object.contract) : "",
      function: isSet(object.function) ? String(object.function) : "",
      args: isSet(object.args) ? String(object.args) : "",
    };
  },

  toJSON(message: MsgExecute): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.contract !== undefined && (obj.contract = message.contract);
    message.function !== undefined && (obj.function = message.function);
    message.args !== undefined && (obj.args = message.args);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgExecute>, I>>(object: I): MsgExecute {
    const message = createBaseMsgExecute();
    message.creator = object.creator ?? "";
    message.contract = object.contract ?? "";
    message.function = object.function ?? "";
    message.args = object.args ?? "";
    return message;
  },
};

function createBaseMsgExecuteResponse(): MsgExecuteResponse {
  return { console: "" };
}

export const MsgExecuteResponse = {
  encode(message: MsgExecuteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.console !== "") {
      writer.uint32(10).string(message.console);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgExecuteResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgExecuteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.console = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExecuteResponse {
    return { console: isSet(object.console) ? String(object.console) : "" };
  },

  toJSON(message: MsgExecuteResponse): unknown {
    const obj: any = {};
    message.console !== undefined && (obj.console = message.console);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgExecuteResponse>, I>>(object: I): MsgExecuteResponse {
    const message = createBaseMsgExecuteResponse();
    message.console = object.console ?? "";
    return message;
  },
};

function createBaseMsgInstantiate(): MsgInstantiate {
  return { creator: "", name: "", code: "" };
}

export const MsgInstantiate = {
  encode(message: MsgInstantiate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.code !== "") {
      writer.uint32(26).string(message.code);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgInstantiate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgInstantiate();
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
          message.code = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInstantiate {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      code: isSet(object.code) ? String(object.code) : "",
    };
  },

  toJSON(message: MsgInstantiate): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.code !== undefined && (obj.code = message.code);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgInstantiate>, I>>(object: I): MsgInstantiate {
    const message = createBaseMsgInstantiate();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.code = object.code ?? "";
    return message;
  },
};

function createBaseMsgInstantiateResponse(): MsgInstantiateResponse {
  return {};
}

export const MsgInstantiateResponse = {
  encode(_: MsgInstantiateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgInstantiateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgInstantiateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgInstantiateResponse {
    return {};
  },

  toJSON(_: MsgInstantiateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgInstantiateResponse>, I>>(_: I): MsgInstantiateResponse {
    const message = createBaseMsgInstantiateResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Store(request: MsgStore): Promise<MsgStoreResponse>;
  Execute(request: MsgExecute): Promise<MsgExecuteResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Instantiate(request: MsgInstantiate): Promise<MsgInstantiateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Store = this.Store.bind(this);
    this.Execute = this.Execute.bind(this);
    this.Instantiate = this.Instantiate.bind(this);
  }
  Store(request: MsgStore): Promise<MsgStoreResponse> {
    const data = MsgStore.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Msg", "Store", data);
    return promise.then((data) => MsgStoreResponse.decode(new _m0.Reader(data)));
  }

  Execute(request: MsgExecute): Promise<MsgExecuteResponse> {
    const data = MsgExecute.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Msg", "Execute", data);
    return promise.then((data) => MsgExecuteResponse.decode(new _m0.Reader(data)));
  }

  Instantiate(request: MsgInstantiate): Promise<MsgInstantiateResponse> {
    const data = MsgInstantiate.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Msg", "Instantiate", data);
    return promise.then((data) => MsgInstantiateResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
