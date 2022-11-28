
function mint() {
    let minted = STD.bank.mint("20")
    if (!minted) {
        return "failed mint"
    }
    minted = STD.bank.withdrawTokens(CTX.SENDER, "20vesta-" + CONTRACT.name)
    if (!minted) {
        return "failed withdraw"
    }
}

CONTRACT.functions.mint = mint

// CONTRACT.queries.balance = balance