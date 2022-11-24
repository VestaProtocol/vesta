/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Contracts } from "./contracts";
import { Params } from "./params";
import { Program } from "./program";

export const protobufPackage = "vesta.vm";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetContractsRequest {
  id: number;
}

export interface QueryGetContractsResponse {
  Contracts: Contracts | undefined;
}

export interface QueryAllContractsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllContractsResponse {
  Contracts: Contracts[];
  pagination: PageResponse | undefined;
}

export interface QueryGetProgramRequest {
  name: string;
}

export interface QueryGetProgramResponse {
  program: Program | undefined;
}

export interface QueryAllProgramRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProgramResponse {
  program: Program[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetContractsRequest(): QueryGetContractsRequest {
  return { id: 0 };
}

export const QueryGetContractsRequest = {
  encode(message: QueryGetContractsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetContractsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetContractsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractsRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetContractsRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetContractsRequest>, I>>(object: I): QueryGetContractsRequest {
    const message = createBaseQueryGetContractsRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetContractsResponse(): QueryGetContractsResponse {
  return { Contracts: undefined };
}

export const QueryGetContractsResponse = {
  encode(message: QueryGetContractsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Contracts !== undefined) {
      Contracts.encode(message.Contracts, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetContractsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetContractsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Contracts = Contracts.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetContractsResponse {
    return { Contracts: isSet(object.Contracts) ? Contracts.fromJSON(object.Contracts) : undefined };
  },

  toJSON(message: QueryGetContractsResponse): unknown {
    const obj: any = {};
    message.Contracts !== undefined
      && (obj.Contracts = message.Contracts ? Contracts.toJSON(message.Contracts) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetContractsResponse>, I>>(object: I): QueryGetContractsResponse {
    const message = createBaseQueryGetContractsResponse();
    message.Contracts = (object.Contracts !== undefined && object.Contracts !== null)
      ? Contracts.fromPartial(object.Contracts)
      : undefined;
    return message;
  },
};

function createBaseQueryAllContractsRequest(): QueryAllContractsRequest {
  return { pagination: undefined };
}

export const QueryAllContractsRequest = {
  encode(message: QueryAllContractsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllContractsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllContractsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllContractsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllContractsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllContractsRequest>, I>>(object: I): QueryAllContractsRequest {
    const message = createBaseQueryAllContractsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllContractsResponse(): QueryAllContractsResponse {
  return { Contracts: [], pagination: undefined };
}

export const QueryAllContractsResponse = {
  encode(message: QueryAllContractsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Contracts) {
      Contracts.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllContractsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllContractsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Contracts.push(Contracts.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllContractsResponse {
    return {
      Contracts: Array.isArray(object?.Contracts) ? object.Contracts.map((e: any) => Contracts.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllContractsResponse): unknown {
    const obj: any = {};
    if (message.Contracts) {
      obj.Contracts = message.Contracts.map((e) => e ? Contracts.toJSON(e) : undefined);
    } else {
      obj.Contracts = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllContractsResponse>, I>>(object: I): QueryAllContractsResponse {
    const message = createBaseQueryAllContractsResponse();
    message.Contracts = object.Contracts?.map((e) => Contracts.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetProgramRequest(): QueryGetProgramRequest {
  return { name: "" };
}

export const QueryGetProgramRequest = {
  encode(message: QueryGetProgramRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProgramRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProgramRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProgramRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: QueryGetProgramRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProgramRequest>, I>>(object: I): QueryGetProgramRequest {
    const message = createBaseQueryGetProgramRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseQueryGetProgramResponse(): QueryGetProgramResponse {
  return { program: undefined };
}

export const QueryGetProgramResponse = {
  encode(message: QueryGetProgramResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.program !== undefined) {
      Program.encode(message.program, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProgramResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProgramResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.program = Program.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProgramResponse {
    return { program: isSet(object.program) ? Program.fromJSON(object.program) : undefined };
  },

  toJSON(message: QueryGetProgramResponse): unknown {
    const obj: any = {};
    message.program !== undefined && (obj.program = message.program ? Program.toJSON(message.program) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProgramResponse>, I>>(object: I): QueryGetProgramResponse {
    const message = createBaseQueryGetProgramResponse();
    message.program = (object.program !== undefined && object.program !== null)
      ? Program.fromPartial(object.program)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProgramRequest(): QueryAllProgramRequest {
  return { pagination: undefined };
}

export const QueryAllProgramRequest = {
  encode(message: QueryAllProgramRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProgramRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProgramRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProgramRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllProgramRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProgramRequest>, I>>(object: I): QueryAllProgramRequest {
    const message = createBaseQueryAllProgramRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProgramResponse(): QueryAllProgramResponse {
  return { program: [], pagination: undefined };
}

export const QueryAllProgramResponse = {
  encode(message: QueryAllProgramResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.program) {
      Program.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProgramResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProgramResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.program.push(Program.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProgramResponse {
    return {
      program: Array.isArray(object?.program) ? object.program.map((e: any) => Program.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllProgramResponse): unknown {
    const obj: any = {};
    if (message.program) {
      obj.program = message.program.map((e) => e ? Program.toJSON(e) : undefined);
    } else {
      obj.program = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProgramResponse>, I>>(object: I): QueryAllProgramResponse {
    const message = createBaseQueryAllProgramResponse();
    message.program = object.program?.map((e) => Program.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Contracts by id. */
  Contracts(request: QueryGetContractsRequest): Promise<QueryGetContractsResponse>;
  /** Queries a list of Contracts items. */
  ContractsAll(request: QueryAllContractsRequest): Promise<QueryAllContractsResponse>;
  /** Queries a Program by index. */
  Program(request: QueryGetProgramRequest): Promise<QueryGetProgramResponse>;
  /** Queries a list of Program items. */
  ProgramAll(request: QueryAllProgramRequest): Promise<QueryAllProgramResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Contracts = this.Contracts.bind(this);
    this.ContractsAll = this.ContractsAll.bind(this);
    this.Program = this.Program.bind(this);
    this.ProgramAll = this.ProgramAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Contracts(request: QueryGetContractsRequest): Promise<QueryGetContractsResponse> {
    const data = QueryGetContractsRequest.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Query", "Contracts", data);
    return promise.then((data) => QueryGetContractsResponse.decode(new _m0.Reader(data)));
  }

  ContractsAll(request: QueryAllContractsRequest): Promise<QueryAllContractsResponse> {
    const data = QueryAllContractsRequest.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Query", "ContractsAll", data);
    return promise.then((data) => QueryAllContractsResponse.decode(new _m0.Reader(data)));
  }

  Program(request: QueryGetProgramRequest): Promise<QueryGetProgramResponse> {
    const data = QueryGetProgramRequest.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Query", "Program", data);
    return promise.then((data) => QueryGetProgramResponse.decode(new _m0.Reader(data)));
  }

  ProgramAll(request: QueryAllProgramRequest): Promise<QueryAllProgramResponse> {
    const data = QueryAllProgramRequest.encode(request).finish();
    const promise = this.rpc.request("vesta.vm.Query", "ProgramAll", data);
    return promise.then((data) => QueryAllProgramResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
