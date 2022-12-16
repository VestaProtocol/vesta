import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCron } from "./types/vesta/vm/tx";
import { MsgInstantiate } from "./types/vesta/vm/tx";
import { MsgExecute } from "./types/vesta/vm/tx";
import { MsgUpgrade } from "./types/vesta/vm/tx";
import { MsgStore } from "./types/vesta/vm/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/vesta.vm.MsgCron", MsgCron],
    ["/vesta.vm.MsgInstantiate", MsgInstantiate],
    ["/vesta.vm.MsgExecute", MsgExecute],
    ["/vesta.vm.MsgUpgrade", MsgUpgrade],
    ["/vesta.vm.MsgStore", MsgStore],
    
];

export { msgTypes }