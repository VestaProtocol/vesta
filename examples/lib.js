function send() {
    var sent = sendTokens("vst1nd3lnnxgymkdny572t84ctqj3jlueypshg4unz", "100token")
    return sent
}

function deposit() {
    var sent = sendTokens(CONTRACT, "100token")
    return sent
}

function withdraw() {
    var sent = withdrawTokens(SENDER, "100token")
    return sent
}

contractFunctions = {
    send: send,
    deposit: deposit
}