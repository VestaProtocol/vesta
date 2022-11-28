var LIB = require("lib")

function deposit() {
    return LIB.deposit()
}

CONTRACT.functions.deposit = deposit
