import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgExecute } from "./types/vesta/vm/tx";
import { MsgStore } from "./types/vesta/vm/tx";
import { MsgInstantiate } from "./types/vesta/vm/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/vesta.vm.MsgExecute", MsgExecute],
    ["/vesta.vm.MsgStore", MsgStore],
    ["/vesta.vm.MsgInstantiate", MsgInstantiate],
    
];

export { msgTypes }