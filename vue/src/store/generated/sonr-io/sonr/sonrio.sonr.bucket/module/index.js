// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteBucket } from "./types/bucket/tx";
import { MsgCreateBucket } from "./types/bucket/tx";
import { MsgReadBucket } from "./types/bucket/tx";
import { MsgUpdateBucket } from "./types/bucket/tx";
const types = [
    ["/sonrio.sonr.bucket.MsgDeleteBucket", MsgDeleteBucket],
    ["/sonrio.sonr.bucket.MsgCreateBucket", MsgCreateBucket],
    ["/sonrio.sonr.bucket.MsgReadBucket", MsgReadBucket],
    ["/sonrio.sonr.bucket.MsgUpdateBucket", MsgUpdateBucket],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgDeleteBucket: (data) => ({ typeUrl: "/sonrio.sonr.bucket.MsgDeleteBucket", value: MsgDeleteBucket.fromPartial(data) }),
        msgCreateBucket: (data) => ({ typeUrl: "/sonrio.sonr.bucket.MsgCreateBucket", value: MsgCreateBucket.fromPartial(data) }),
        msgReadBucket: (data) => ({ typeUrl: "/sonrio.sonr.bucket.MsgReadBucket", value: MsgReadBucket.fromPartial(data) }),
        msgUpdateBucket: (data) => ({ typeUrl: "/sonrio.sonr.bucket.MsgUpdateBucket", value: MsgUpdateBucket.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
