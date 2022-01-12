// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeletePoll } from "./types/voter/tx";
import { MsgUpdatePoll } from "./types/voter/tx";
import { MsgCreatePoll } from "./types/voter/tx";
const types = [
    ["/iroan.voter.voter.MsgDeletePoll", MsgDeletePoll],
    ["/iroan.voter.voter.MsgUpdatePoll", MsgUpdatePoll],
    ["/iroan.voter.voter.MsgCreatePoll", MsgCreatePoll],
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
        msgDeletePoll: (data) => ({ typeUrl: "/iroan.voter.voter.MsgDeletePoll", value: MsgDeletePoll.fromPartial(data) }),
        msgUpdatePoll: (data) => ({ typeUrl: "/iroan.voter.voter.MsgUpdatePoll", value: MsgUpdatePoll.fromPartial(data) }),
        msgCreatePoll: (data) => ({ typeUrl: "/iroan.voter.voter.MsgCreatePoll", value: MsgCreatePoll.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
