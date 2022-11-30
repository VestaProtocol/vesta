
CONTRACT.functions.mint = function () {
    let minted = STD.bank.mint("1000000")
    if (!minted) {
        return "failed mint"
    }
}

CONTRACT.functions.fund = function (runs) {
    var sent = STD.bank.sendTokens(CONTRACT.address, (500000 * Number(runs)) + "token")
    return sent
}

CONTRACT.queries.balance = function (address) {
    return STD.bank.balance(CONTRACT.address, STD.bank.token)
}
