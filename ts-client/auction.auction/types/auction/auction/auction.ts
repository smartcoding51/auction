/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "auction.auction";

export interface Auction {
  index: string;
  name: string;
  endAt: string;
  minBid: number;
  highestBid: number;
  highestBidder: string;
  seller: string;
  state: string;
}

function createBaseAuction(): Auction {
  return { index: "", name: "", endAt: "", minBid: 0, highestBid: 0, highestBidder: "", seller: "", state: "" };
}

export const Auction = {
  encode(message: Auction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.endAt !== "") {
      writer.uint32(26).string(message.endAt);
    }
    if (message.minBid !== 0) {
      writer.uint32(32).uint64(message.minBid);
    }
    if (message.highestBid !== 0) {
      writer.uint32(40).uint64(message.highestBid);
    }
    if (message.highestBidder !== "") {
      writer.uint32(50).string(message.highestBidder);
    }
    if (message.seller !== "") {
      writer.uint32(58).string(message.seller);
    }
    if (message.state !== "") {
      writer.uint32(66).string(message.state);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Auction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.endAt = reader.string();
          break;
        case 4:
          message.minBid = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.highestBid = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.highestBidder = reader.string();
          break;
        case 7:
          message.seller = reader.string();
          break;
        case 8:
          message.state = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Auction {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      name: isSet(object.name) ? String(object.name) : "",
      endAt: isSet(object.endAt) ? String(object.endAt) : "",
      minBid: isSet(object.minBid) ? Number(object.minBid) : 0,
      highestBid: isSet(object.highestBid) ? Number(object.highestBid) : 0,
      highestBidder: isSet(object.highestBidder) ? String(object.highestBidder) : "",
      seller: isSet(object.seller) ? String(object.seller) : "",
      state: isSet(object.state) ? String(object.state) : "",
    };
  },

  toJSON(message: Auction): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.endAt !== undefined && (obj.endAt = message.endAt);
    message.minBid !== undefined && (obj.minBid = Math.round(message.minBid));
    message.highestBid !== undefined && (obj.highestBid = Math.round(message.highestBid));
    message.highestBidder !== undefined && (obj.highestBidder = message.highestBidder);
    message.seller !== undefined && (obj.seller = message.seller);
    message.state !== undefined && (obj.state = message.state);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Auction>, I>>(object: I): Auction {
    const message = createBaseAuction();
    message.index = object.index ?? "";
    message.name = object.name ?? "";
    message.endAt = object.endAt ?? "";
    message.minBid = object.minBid ?? 0;
    message.highestBid = object.highestBid ?? 0;
    message.highestBidder = object.highestBidder ?? "";
    message.seller = object.seller ?? "";
    message.state = object.state ?? "";
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
