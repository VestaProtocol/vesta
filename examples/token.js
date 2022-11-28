
function mint() {
    let minted = STD.bank.mint("20")
    if (!minted) {
        return "failed mint"
    }
    minted = STD.bank.withdrawTokens(CTX.SENDER, "20" + STD.bank.token)
    if (!minted) {
        return "failed withdraw"
    }
}

function balance(address) {
    return STD.bank.balance(address, STD.bank.token)
}

CONTRACT.functions.mint = mint

CONTRACT.queries.balance = balance