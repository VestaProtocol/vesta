const LIB = STD.require("lib", "0")

function deposit() {
    return LIB.functions.deposit()
}

CONTRACT.functions.deposit = deposit
