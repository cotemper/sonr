// @generated by protobuf-ts 2.2.1 with parameter long_type_string,generate_dependencies
// @generated from protobuf file "bucket/tx.proto" (package "sonrio.sonr.bucket", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { Msg } from "./tx";
import type { MsgDeleteBucketResponse } from "./tx";
import type { MsgDeleteBucket } from "./tx";
import type { MsgUpdateBucketResponse } from "./tx";
import type { MsgUpdateBucket } from "./tx";
import type { MsgReadBucketResponse } from "./tx";
import type { MsgReadBucket } from "./tx";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { MsgCreateBucketResponse } from "./tx";
import type { MsgCreateBucket } from "./tx";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * Msg defines the Msg service.
 *
 * @generated from protobuf service sonrio.sonr.bucket.Msg
 */
export interface IMsgClient {
    /**
     * @generated from protobuf rpc: CreateBucket(sonrio.sonr.bucket.MsgCreateBucket) returns (sonrio.sonr.bucket.MsgCreateBucketResponse);
     */
    createBucket(input: MsgCreateBucket, options?: RpcOptions): UnaryCall<MsgCreateBucket, MsgCreateBucketResponse>;
    /**
     * @generated from protobuf rpc: ReadBucket(sonrio.sonr.bucket.MsgReadBucket) returns (sonrio.sonr.bucket.MsgReadBucketResponse);
     */
    readBucket(input: MsgReadBucket, options?: RpcOptions): UnaryCall<MsgReadBucket, MsgReadBucketResponse>;
    /**
     * @generated from protobuf rpc: UpdateBucket(sonrio.sonr.bucket.MsgUpdateBucket) returns (sonrio.sonr.bucket.MsgUpdateBucketResponse);
     */
    updateBucket(input: MsgUpdateBucket, options?: RpcOptions): UnaryCall<MsgUpdateBucket, MsgUpdateBucketResponse>;
    /**
     * this line is used by starport scaffolding # proto/tx/rpc
     *
     * @generated from protobuf rpc: DeleteBucket(sonrio.sonr.bucket.MsgDeleteBucket) returns (sonrio.sonr.bucket.MsgDeleteBucketResponse);
     */
    deleteBucket(input: MsgDeleteBucket, options?: RpcOptions): UnaryCall<MsgDeleteBucket, MsgDeleteBucketResponse>;
}
/**
 * Msg defines the Msg service.
 *
 * @generated from protobuf service sonrio.sonr.bucket.Msg
 */
export class MsgClient implements IMsgClient, ServiceInfo {
    typeName = Msg.typeName;
    methods = Msg.methods;
    options = Msg.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: CreateBucket(sonrio.sonr.bucket.MsgCreateBucket) returns (sonrio.sonr.bucket.MsgCreateBucketResponse);
     */
    createBucket(input: MsgCreateBucket, options?: RpcOptions): UnaryCall<MsgCreateBucket, MsgCreateBucketResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<MsgCreateBucket, MsgCreateBucketResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ReadBucket(sonrio.sonr.bucket.MsgReadBucket) returns (sonrio.sonr.bucket.MsgReadBucketResponse);
     */
    readBucket(input: MsgReadBucket, options?: RpcOptions): UnaryCall<MsgReadBucket, MsgReadBucketResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<MsgReadBucket, MsgReadBucketResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateBucket(sonrio.sonr.bucket.MsgUpdateBucket) returns (sonrio.sonr.bucket.MsgUpdateBucketResponse);
     */
    updateBucket(input: MsgUpdateBucket, options?: RpcOptions): UnaryCall<MsgUpdateBucket, MsgUpdateBucketResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<MsgUpdateBucket, MsgUpdateBucketResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * this line is used by starport scaffolding # proto/tx/rpc
     *
     * @generated from protobuf rpc: DeleteBucket(sonrio.sonr.bucket.MsgDeleteBucket) returns (sonrio.sonr.bucket.MsgDeleteBucketResponse);
     */
    deleteBucket(input: MsgDeleteBucket, options?: RpcOptions): UnaryCall<MsgDeleteBucket, MsgDeleteBucketResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<MsgDeleteBucket, MsgDeleteBucketResponse>("unary", this._transport, method, opt, input);
    }
}