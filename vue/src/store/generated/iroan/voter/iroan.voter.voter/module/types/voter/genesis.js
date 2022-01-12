/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../voter/params";
import { Poll } from "../voter/poll";
export const protobufPackage = "iroan.voter.voter";
const baseGenesisState = { pollCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.pollList) {
            Poll.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.pollCount !== 0) {
            writer.uint32(24).uint64(message.pollCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.pollList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.pollList.push(Poll.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.pollCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.pollList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.pollList !== undefined && object.pollList !== null) {
            for (const e of object.pollList) {
                message.pollList.push(Poll.fromJSON(e));
            }
        }
        if (object.pollCount !== undefined && object.pollCount !== null) {
            message.pollCount = Number(object.pollCount);
        }
        else {
            message.pollCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.pollList) {
            obj.pollList = message.pollList.map((e) => e ? Poll.toJSON(e) : undefined);
        }
        else {
            obj.pollList = [];
        }
        message.pollCount !== undefined && (obj.pollCount = message.pollCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.pollList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.pollList !== undefined && object.pollList !== null) {
            for (const e of object.pollList) {
                message.pollList.push(Poll.fromPartial(e));
            }
        }
        if (object.pollCount !== undefined && object.pollCount !== null) {
            message.pollCount = object.pollCount;
        }
        else {
            message.pollCount = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
