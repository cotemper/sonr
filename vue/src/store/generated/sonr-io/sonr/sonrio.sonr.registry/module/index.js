// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRegisterName } from "./types/registry/tx";
import { MsgCreateAccount } from "./types/registry/tx";
import { MsgRegisterService } from "./types/registry/tx";
const types = [
    ["/sonrio.sonr.registry.MsgRegisterName", MsgRegisterName],
    ["/sonrio.sonr.registry.MsgCreateAccount", MsgCreateAccount],
    ["/sonrio.sonr.registry.MsgRegisterService", MsgRegisterService],
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
        msgRegisterName: (data) => ({ typeUrl: "/sonrio.sonr.registry.MsgRegisterName", value: MsgRegisterName.fromPartial(data) }),
        msgCreateAccount: (data) => ({ typeUrl: "/sonrio.sonr.registry.MsgCreateAccount", value: MsgCreateAccount.fromPartial(data) }),
        msgRegisterService: (data) => ({ typeUrl: "/sonrio.sonr.registry.MsgRegisterService", value: MsgRegisterService.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
