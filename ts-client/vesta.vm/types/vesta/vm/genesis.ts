/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Contracts } from "./contracts";
import { Cronjobs } from "./cronjobs";
import { Params } from "./params";
import { Program } from "./program";
import { Romdata } from "./romdata";

export const protobufPackage = "vesta.vm";

/** GenesisState defines the vm module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  contractsList: Contracts[];
  contractsCount: number;
  programList: Program[];
  romdataList: Romdata[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  cronjobsList: Cronjobs[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    contractsList: [],
    contractsCount: 0,
    programList: [],
    romdataList: [],
    cronjobsList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.contractsList) {
      Contracts.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.contractsCount !== 0) {
      writer.uint32(24).uint64(message.contractsCount);
    }
    for (const v of message.programList) {
      Program.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.romdataList) {
      Romdata.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.cronjobsList) {
      Cronjobs.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.contractsList.push(Contracts.decode(reader, reader.uint32()));
          break;
        case 3:
          message.contractsCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.programList.push(Program.decode(reader, reader.uint32()));
          break;
        case 5:
          message.romdataList.push(Romdata.decode(reader, reader.uint32()));
          break;
        case 6:
          message.cronjobsList.push(Cronjobs.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      contractsList: Array.isArray(object?.contractsList)
        ? object.contractsList.map((e: any) => Contracts.fromJSON(e))
        : [],
      contractsCount: isSet(object.contractsCount) ? Number(object.contractsCount) : 0,
      programList: Array.isArray(object?.programList) ? object.programList.map((e: any) => Program.fromJSON(e)) : [],
      romdataList: Array.isArray(object?.romdataList) ? object.romdataList.map((e: any) => Romdata.fromJSON(e)) : [],
      cronjobsList: Array.isArray(object?.cronjobsList)
        ? object.cronjobsList.map((e: any) => Cronjobs.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.contractsList) {
      obj.contractsList = message.contractsList.map((e) => e ? Contracts.toJSON(e) : undefined);
    } else {
      obj.contractsList = [];
    }
    message.contractsCount !== undefined && (obj.contractsCount = Math.round(message.contractsCount));
    if (message.programList) {
      obj.programList = message.programList.map((e) => e ? Program.toJSON(e) : undefined);
    } else {
      obj.programList = [];
    }
    if (message.romdataList) {
      obj.romdataList = message.romdataList.map((e) => e ? Romdata.toJSON(e) : undefined);
    } else {
      obj.romdataList = [];
    }
    if (message.cronjobsList) {
      obj.cronjobsList = message.cronjobsList.map((e) => e ? Cronjobs.toJSON(e) : undefined);
    } else {
      obj.cronjobsList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.contractsList = object.contractsList?.map((e) => Contracts.fromPartial(e)) || [];
    message.contractsCount = object.contractsCount ?? 0;
    message.programList = object.programList?.map((e) => Program.fromPartial(e)) || [];
    message.romdataList = object.romdataList?.map((e) => Romdata.fromPartial(e)) || [];
    message.cronjobsList = object.cronjobsList?.map((e) => Cronjobs.fromPartial(e)) || [];
    return message;
  },
};

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
