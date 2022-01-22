// @generated by protobuf-ts 2.2.1 with parameter long_type_string,generate_dependencies
// @generated from protobuf file "bucket/tx.proto" (package "sonrio.sonr.bucket", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgCreateBucket
 */
export interface MsgCreateBucket {
    /**
     * @generated from protobuf field: string creator = 1;
     */
    creator: string;
    /**
     * @generated from protobuf field: string label = 2;
     */
    label: string;
    /**
     * @generated from protobuf field: string description = 3;
     */
    description: string;
    /**
     * @generated from protobuf field: string kind = 4;
     */
    kind: string;
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgCreateBucketResponse
 */
export interface MsgCreateBucketResponse {
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgReadBucket
 */
export interface MsgReadBucket {
    /**
     * @generated from protobuf field: string creator = 1;
     */
    creator: string;
    /**
     * @generated from protobuf field: string did = 2;
     */
    did: string;
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgReadBucketResponse
 */
export interface MsgReadBucketResponse {
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgUpdateBucket
 */
export interface MsgUpdateBucket {
    /**
     * @generated from protobuf field: string creator = 1;
     */
    creator: string;
    /**
     * @generated from protobuf field: string did = 2;
     */
    did: string;
    /**
     * @generated from protobuf field: string label = 3;
     */
    label: string;
    /**
     * @generated from protobuf field: string description = 4;
     */
    description: string;
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgUpdateBucketResponse
 */
export interface MsgUpdateBucketResponse {
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgDeleteBucket
 */
export interface MsgDeleteBucket {
    /**
     * @generated from protobuf field: string creator = 1;
     */
    creator: string;
    /**
     * @generated from protobuf field: string did = 2;
     */
    did: string;
    /**
     * @generated from protobuf field: string publicKey = 3;
     */
    publicKey: string;
}
/**
 * @generated from protobuf message sonrio.sonr.bucket.MsgDeleteBucketResponse
 */
export interface MsgDeleteBucketResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class MsgCreateBucket$Type extends MessageType<MsgCreateBucket> {
    constructor() {
        super("sonrio.sonr.bucket.MsgCreateBucket", [
            { no: 1, name: "creator", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "label", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "kind", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<MsgCreateBucket>): MsgCreateBucket {
        const message = { creator: "", label: "", description: "", kind: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgCreateBucket>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgCreateBucket): MsgCreateBucket {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string creator */ 1:
                    message.creator = reader.string();
                    break;
                case /* string label */ 2:
                    message.label = reader.string();
                    break;
                case /* string description */ 3:
                    message.description = reader.string();
                    break;
                case /* string kind */ 4:
                    message.kind = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MsgCreateBucket, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string creator = 1; */
        if (message.creator !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.creator);
        /* string label = 2; */
        if (message.label !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.label);
        /* string description = 3; */
        if (message.description !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.description);
        /* string kind = 4; */
        if (message.kind !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.kind);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgCreateBucket
 */
export const MsgCreateBucket = new MsgCreateBucket$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgCreateBucketResponse$Type extends MessageType<MsgCreateBucketResponse> {
    constructor() {
        super("sonrio.sonr.bucket.MsgCreateBucketResponse", []);
    }
    create(value?: PartialMessage<MsgCreateBucketResponse>): MsgCreateBucketResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgCreateBucketResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgCreateBucketResponse): MsgCreateBucketResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: MsgCreateBucketResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgCreateBucketResponse
 */
export const MsgCreateBucketResponse = new MsgCreateBucketResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgReadBucket$Type extends MessageType<MsgReadBucket> {
    constructor() {
        super("sonrio.sonr.bucket.MsgReadBucket", [
            { no: 1, name: "creator", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "did", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<MsgReadBucket>): MsgReadBucket {
        const message = { creator: "", did: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgReadBucket>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgReadBucket): MsgReadBucket {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string creator */ 1:
                    message.creator = reader.string();
                    break;
                case /* string did */ 2:
                    message.did = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MsgReadBucket, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string creator = 1; */
        if (message.creator !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.creator);
        /* string did = 2; */
        if (message.did !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.did);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgReadBucket
 */
export const MsgReadBucket = new MsgReadBucket$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgReadBucketResponse$Type extends MessageType<MsgReadBucketResponse> {
    constructor() {
        super("sonrio.sonr.bucket.MsgReadBucketResponse", []);
    }
    create(value?: PartialMessage<MsgReadBucketResponse>): MsgReadBucketResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgReadBucketResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgReadBucketResponse): MsgReadBucketResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: MsgReadBucketResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgReadBucketResponse
 */
export const MsgReadBucketResponse = new MsgReadBucketResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgUpdateBucket$Type extends MessageType<MsgUpdateBucket> {
    constructor() {
        super("sonrio.sonr.bucket.MsgUpdateBucket", [
            { no: 1, name: "creator", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "did", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "label", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<MsgUpdateBucket>): MsgUpdateBucket {
        const message = { creator: "", did: "", label: "", description: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgUpdateBucket>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgUpdateBucket): MsgUpdateBucket {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string creator */ 1:
                    message.creator = reader.string();
                    break;
                case /* string did */ 2:
                    message.did = reader.string();
                    break;
                case /* string label */ 3:
                    message.label = reader.string();
                    break;
                case /* string description */ 4:
                    message.description = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MsgUpdateBucket, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string creator = 1; */
        if (message.creator !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.creator);
        /* string did = 2; */
        if (message.did !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.did);
        /* string label = 3; */
        if (message.label !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.label);
        /* string description = 4; */
        if (message.description !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.description);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgUpdateBucket
 */
export const MsgUpdateBucket = new MsgUpdateBucket$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgUpdateBucketResponse$Type extends MessageType<MsgUpdateBucketResponse> {
    constructor() {
        super("sonrio.sonr.bucket.MsgUpdateBucketResponse", []);
    }
    create(value?: PartialMessage<MsgUpdateBucketResponse>): MsgUpdateBucketResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgUpdateBucketResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgUpdateBucketResponse): MsgUpdateBucketResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: MsgUpdateBucketResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgUpdateBucketResponse
 */
export const MsgUpdateBucketResponse = new MsgUpdateBucketResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgDeleteBucket$Type extends MessageType<MsgDeleteBucket> {
    constructor() {
        super("sonrio.sonr.bucket.MsgDeleteBucket", [
            { no: 1, name: "creator", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "did", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "publicKey", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<MsgDeleteBucket>): MsgDeleteBucket {
        const message = { creator: "", did: "", publicKey: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgDeleteBucket>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgDeleteBucket): MsgDeleteBucket {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string creator */ 1:
                    message.creator = reader.string();
                    break;
                case /* string did */ 2:
                    message.did = reader.string();
                    break;
                case /* string publicKey */ 3:
                    message.publicKey = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MsgDeleteBucket, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string creator = 1; */
        if (message.creator !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.creator);
        /* string did = 2; */
        if (message.did !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.did);
        /* string publicKey = 3; */
        if (message.publicKey !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.publicKey);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgDeleteBucket
 */
export const MsgDeleteBucket = new MsgDeleteBucket$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MsgDeleteBucketResponse$Type extends MessageType<MsgDeleteBucketResponse> {
    constructor() {
        super("sonrio.sonr.bucket.MsgDeleteBucketResponse", []);
    }
    create(value?: PartialMessage<MsgDeleteBucketResponse>): MsgDeleteBucketResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MsgDeleteBucketResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MsgDeleteBucketResponse): MsgDeleteBucketResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: MsgDeleteBucketResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message sonrio.sonr.bucket.MsgDeleteBucketResponse
 */
export const MsgDeleteBucketResponse = new MsgDeleteBucketResponse$Type();
/**
 * @generated ServiceType for protobuf service sonrio.sonr.bucket.Msg
 */
export const Msg = new ServiceType("sonrio.sonr.bucket.Msg", [
    { name: "CreateBucket", options: {}, I: MsgCreateBucket, O: MsgCreateBucketResponse },
    { name: "ReadBucket", options: {}, I: MsgReadBucket, O: MsgReadBucketResponse },
    { name: "UpdateBucket", options: {}, I: MsgUpdateBucket, O: MsgUpdateBucketResponse },
    { name: "DeleteBucket", options: {}, I: MsgDeleteBucket, O: MsgDeleteBucketResponse }
]);