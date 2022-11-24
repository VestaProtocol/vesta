/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "vesta.vm";

export interface Program {
  name: string;
  creator: string;
  code: number;
}

function createBaseProgram(): Program {
  return { name: "", creator: "", code: 0 };
}

export const Program = {
  encode(message: Program, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.code !== 0) {
      writer.uint32(24).uint64(message.code);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Program {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProgram();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.code = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Program {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      code: isSet(object.code) ? Number(object.code) : 0,
    };
  },

  toJSON(message: Program): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.creator !== undefined && (obj.creator = message.creator);
    message.code !== undefined && (obj.code = Math.round(message.code));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Program>, I>>(object: I): Program {
    const message = createBaseProgram();
    message.name = object.name ?? "";
    message.creator = object.creator ?? "";
    message.code = object.code ?? 0;
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
