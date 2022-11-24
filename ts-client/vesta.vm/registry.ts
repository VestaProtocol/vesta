import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgStore } from "./types/vesta/vm/tx";
import { MsgExecute } from "./types/vesta/vm/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/vesta.vm.MsgStore", MsgStore],
    ["/vesta.vm.MsgExecute", MsgExecute],
    
];

export { msgTypes }