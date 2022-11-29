function send() {
    var sent = STD.bank.sendTokens("vst1nd3lnnxgymkdny572t84ctqj3jlueypshg4unz", "100token")
    return sent
}

function deposit() {
    var sent = STD.bank.sendTokens(CONTRACT.address, "100token")
    return sent
}

function withdraw() {
    var sent = STD.bank.withdrawTokens(CTX.sender, "100token")
    return sent
}

CONTRACT.functions.send = send
CONTRACT.functions.deposit = deposit
CONTRACT.functions.withdraw = withdraw