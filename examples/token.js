
CONTRACT.functions.mint = function() {
    let minted = STD.bank.mint("1000000")
    if (!minted) {
        return "failed mint"
    }
    minted = STD.bank.withdrawTokens(CTX.sender, "1000000" + STD.bank.token)
    if (!minted) {
        return "failed withdraw"
    }
}

CONTRACT.queries.balance = function(address) {
    return STD.bank.balance(address, STD.bank.token)
}

