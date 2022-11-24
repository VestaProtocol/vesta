// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgExecute } from "./types/vesta/vm/tx";
import { MsgInstantiate } from "./types/vesta/vm/tx";
import { MsgStore } from "./types/vesta/vm/tx";


export { MsgExecute, MsgInstantiate, MsgStore };

type sendMsgExecuteParams = {
  value: MsgExecute,
  fee?: StdFee,
  memo?: string
};

type sendMsgInstantiateParams = {
  value: MsgInstantiate,
  fee?: StdFee,
  memo?: string
};

type sendMsgStoreParams = {
  value: MsgStore,
  fee?: StdFee,
  memo?: string
};


type msgExecuteParams = {
  value: MsgExecute,
};

type msgInstantiateParams = {
  value: MsgInstantiate,
};

type msgStoreParams = {
  value: MsgStore,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgExecute({ value, fee, memo }: sendMsgExecuteParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgExecute: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgExecute({ value: MsgExecute.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgExecute: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgInstantiate({ value, fee, memo }: sendMsgInstantiateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgInstantiate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgInstantiate({ value: MsgInstantiate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgInstantiate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgStore({ value, fee, memo }: sendMsgStoreParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgStore: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgStore({ value: MsgStore.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgStore: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgExecute({ value }: msgExecuteParams): EncodeObject {
			try {
				return { typeUrl: "/vesta.vm.MsgExecute", value: MsgExecute.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgExecute: Could not create message: ' + e.message)
			}
		},
		
		msgInstantiate({ value }: msgInstantiateParams): EncodeObject {
			try {
				return { typeUrl: "/vesta.vm.MsgInstantiate", value: MsgInstantiate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgInstantiate: Could not create message: ' + e.message)
			}
		},
		
		msgStore({ value }: msgStoreParams): EncodeObject {
			try {
				return { typeUrl: "/vesta.vm.MsgStore", value: MsgStore.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgStore: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			VestaVm: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;